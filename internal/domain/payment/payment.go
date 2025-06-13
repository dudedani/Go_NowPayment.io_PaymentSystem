package payment

import (
	"time"

	"github.com/google/uuid"
)

// Payment represents a payment transaction (Aggregate Root)
type Payment struct {
	// Identity
	ID      string
	OrderID string

	// Payment Details
	Amount         float64
	Currency       string      // Fiat currency (USD, EUR, etc.)
	CryptoAmount   float64     // Amount in cryptocurrency
	CryptoCurrency CryptoCurrency
	
	// Status and Lifecycle
	Status    PaymentStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
	
	// Transaction Details
	PaymentMethod    PaymentMethod
	TransactionHash  string    // Blockchain transaction hash
	Confirmations    int       // Number of blockchain confirmations
	RequiredConfirmations int  // Required confirmations for completion

	// External Service Integration
	NowPaymentsID    string    // NowPayments payment ID
	CallbackURL      string    // Webhook callback URL
	
	// Refund Information
	RefundedAmount   float64
	RefundTransactionHash string
	RefundedAt       *time.Time
}

// NewPayment creates a new payment with validation
func NewPayment(orderID string, amount float64, currency string, cryptoSymbol string, walletAddress string, expirationMinutes int) (*Payment, error) {
	// Validate inputs
	if orderID == "" {
		return nil, ErrEmptyOrderID
	}
	
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	
	if currency == "" {
		return nil, ErrInvalidCurrency
	}
	
	// Create payment method
	paymentMethod, err := NewPaymentMethod(cryptoSymbol, walletAddress, expirationMinutes)
	if err != nil {
		return nil, err
	}
	
	// Validate crypto amount (will be set later via UpdateCryptoAmount)
	crypto, err := GetCryptoCurrencyBySymbol(cryptoSymbol)
	if err != nil {
		return nil, err
	}
	
	now := time.Now()
	expiresAt := now.Add(time.Duration(expirationMinutes) * time.Minute)
	
	payment := &Payment{
		ID:      uuid.New().String(),
		OrderID: orderID,
		
		Amount:         amount,
		Currency:       currency,
		CryptoAmount:   0, // Will be set when crypto rate is calculated
		CryptoCurrency: crypto,
		
		Status:    StatusPending,
		CreatedAt: now,
		UpdatedAt: now,
		ExpiresAt: expiresAt,
		
		PaymentMethod:         paymentMethod,
		TransactionHash:       "",
		Confirmations:         0,
		RequiredConfirmations: getRequiredConfirmations(crypto),
		
		RefundedAmount: 0,
	}
	
	return payment, nil
}

// UpdateCryptoAmount sets the cryptocurrency amount based on current exchange rates
func (p *Payment) UpdateCryptoAmount(cryptoAmount float64) error {
	if p.Status.IsFinal() {
		return ErrCannotUpdateFinalPayment
	}
	
	if cryptoAmount <= 0 {
		return ErrInvalidAmount
	}
	
	// Validate amount meets minimum requirements
	if err := p.CryptoCurrency.ValidateAmount(cryptoAmount); err != nil {
		return err
	}
	
	p.CryptoAmount = cryptoAmount
	p.UpdatedAt = time.Now()
	
	return nil
}

// SetNowPaymentsID sets the external payment service ID
func (p *Payment) SetNowPaymentsID(nowPaymentsID string) error {
	if p.Status.IsFinal() {
		return ErrCannotUpdateFinalPayment
	}
	
	if nowPaymentsID == "" {
		return ErrEmptyPaymentID
	}
	
	p.NowPaymentsID = nowPaymentsID
	p.UpdatedAt = time.Now()
	
	return nil
}

// SetCallbackURL sets the webhook callback URL
func (p *Payment) SetCallbackURL(callbackURL string) error {
	if p.Status.IsFinal() {
		return ErrCannotUpdateFinalPayment
	}
	
	p.CallbackURL = callbackURL
	p.UpdatedAt = time.Now()
	
	return nil
}

// MarkAsConfirming transitions payment to confirming status when transaction is detected
func (p *Payment) MarkAsConfirming(transactionHash string) error {
	if p.Status != StatusPending {
		return ErrInvalidStatusTransition
	}
	
	if p.IsExpired() {
		return ErrPaymentExpired
	}
	
	if transactionHash == "" {
		return ErrInvalidTransactionHash
	}
	
	p.Status = StatusConfirming
	p.TransactionHash = transactionHash
	p.Confirmations = 0
	p.UpdatedAt = time.Now()
	
	return nil
}

// UpdateConfirmations updates the number of blockchain confirmations
func (p *Payment) UpdateConfirmations(confirmations int) error {
	if p.Status != StatusConfirming {
		return ErrInvalidStatusTransition
	}
	
	if confirmations < 0 {
		return ErrInsufficientConfirmations
	}
	
	p.Confirmations = confirmations
	p.UpdatedAt = time.Now()
	
	// Auto-confirm if we have enough confirmations
	if confirmations >= p.RequiredConfirmations {
		return p.markAsConfirmed()
	}
	
	return nil
}

// markAsConfirmed transitions payment to confirmed status (internal method)
func (p *Payment) markAsConfirmed() error {
	// Allow confirming from pending, confirming, or failed status
	if p.Status != StatusPending && p.Status != StatusConfirming && p.Status != StatusFailed {
		return ErrInvalidStatusTransition
	}
	
	p.Status = StatusConfirmed
	p.UpdatedAt = time.Now()
	
	return nil
}

// MarkAsConfirmed manually confirms the payment (for admin actions)
func (p *Payment) MarkAsConfirmed() error {
	if p.Status == StatusConfirmed {
		return ErrPaymentAlreadyConfirmed
	}
	
	// Allow confirming from pending, confirming, or failed status
	if p.Status != StatusPending && p.Status != StatusConfirming && p.Status != StatusFailed {
		return ErrCannotConfirmPayment
	}
	
	return p.markAsConfirmed()
}

// MarkAsFailed transitions payment to failed status
func (p *Payment) MarkAsFailed() error {
	if p.Status.IsFinal() {
		return ErrInvalidStatusTransition
	}
	
	p.Status = StatusFailed
	p.UpdatedAt = time.Now()
	
	return nil
}

// MarkAsExpired transitions payment to expired status
func (p *Payment) MarkAsExpired() error {
	if p.Status.IsFinal() {
		return ErrInvalidStatusTransition
	}
	
	p.Status = StatusExpired
	p.UpdatedAt = time.Now()
	
	return nil
}

// Cancel cancels the payment if it's still cancellable
func (p *Payment) Cancel() error {
	if !p.Status.CanBeCancelled() {
		return ErrCannotCancelPayment
	}
	
	p.Status = StatusCancelled
	p.UpdatedAt = time.Now()
	
	return nil
}

// Refund processes a full refund of the payment
func (p *Payment) Refund() error {
	return p.PartialRefund(p.CryptoAmount)
}

// PartialRefund processes a partial refund of the payment
func (p *Payment) PartialRefund(refundAmount float64) error {
	if !p.Status.CanBeRefunded() {
		return ErrCannotRefundPayment
	}
	
	if refundAmount <= 0 {
		return ErrInvalidAmount
	}
	
	if p.RefundedAmount+refundAmount > p.CryptoAmount {
		return ErrRefundAmountExceedsPayment
	}
	
	if p.RefundedAmount > 0 {
		return ErrRefundAlreadyProcessed
	}
	
	p.RefundedAmount += refundAmount
	now := time.Now()
	p.RefundedAt = &now
	
	// If fully refunded, mark as refunded
	if p.RefundedAmount >= p.CryptoAmount {
		p.Status = StatusRefunded
	}
	
	p.UpdatedAt = time.Now()
	
	return nil
}

// SetRefundTransactionHash sets the blockchain transaction hash for the refund
func (p *Payment) SetRefundTransactionHash(transactionHash string) error {
	if p.RefundedAmount == 0 {
		return ErrRefundAlreadyProcessed
	}
	
	if transactionHash == "" {
		return ErrInvalidTransactionHash
	}
	
	p.RefundTransactionHash = transactionHash
	p.UpdatedAt = time.Now()
	
	return nil
}

// ValidateAmount checks if the provided amount matches the expected payment amount
func (p *Payment) ValidateAmount(receivedAmount float64) error {
	tolerance := 0.0001 // Small tolerance for crypto amounts (0.01%)
	
	if receivedAmount < (p.CryptoAmount - tolerance) {
		return ErrInsufficientAmount
	}
	
	if receivedAmount > (p.CryptoAmount + tolerance) {
		return ErrExcessiveAmount
	}
	
	return nil
}

// Query methods

// IsExpired checks if the payment has expired
func (p *Payment) IsExpired() bool {
	return time.Now().After(p.ExpiresAt)
}

// IsCompleted checks if the payment is completed
func (p *Payment) IsCompleted() bool {
	return p.Status.IsCompleted()
}

// IsPending checks if the payment is still pending
func (p *Payment) IsPending() bool {
	return p.Status == StatusPending
}

// IsConfirming checks if the payment is confirming
func (p *Payment) IsConfirming() bool {
	return p.Status == StatusConfirming
}

// CanBeCancelled checks if the payment can be cancelled
func (p *Payment) CanBeCancelled() bool {
	return p.Status.CanBeCancelled()
}

// CanBeRefunded checks if the payment can be refunded
func (p *Payment) CanBeRefunded() bool {
	return p.Status.CanBeRefunded()
}

// GetTimeUntilExpiry returns time until payment expires
func (p *Payment) GetTimeUntilExpiry() time.Duration {
	if p.IsExpired() {
		return 0
	}
	return time.Until(p.ExpiresAt)
}

// GetRemainingRefundableAmount returns the amount that can still be refunded
func (p *Payment) GetRemainingRefundableAmount() float64 {
	if !p.CanBeRefunded() {
		return 0
	}
	return p.CryptoAmount - p.RefundedAmount
}

// IsFullyRefunded checks if the payment has been fully refunded
func (p *Payment) IsFullyRefunded() bool {
	return p.RefundedAmount >= p.CryptoAmount && p.RefundedAmount > 0
}

// GetCryptoSymbol returns the cryptocurrency symbol
func (p *Payment) GetCryptoSymbol() string {
	return p.CryptoCurrency.Symbol
}

// GetWalletAddress returns the destination wallet address
func (p *Payment) GetWalletAddress() string {
	return p.PaymentMethod.WalletAddress
}

// Helper functions

// getRequiredConfirmations returns the required number of confirmations for a cryptocurrency
func getRequiredConfirmations(crypto CryptoCurrency) int {
	switch crypto.Symbol {
	case "BTC":
		return 2  // Bitcoin requires 2 confirmations
	case "ETH":
		return 12 // Ethereum requires 12 confirmations
	case "LTC":
		return 6  // Litecoin requires 6 confirmations
	case "BCH":
		return 6  // Bitcoin Cash requires 6 confirmations
	case "XRP":
		return 1  // Ripple requires 1 confirmation
	case "DOGE":
		return 6  // Dogecoin requires 6 confirmations
	default:
		return 6  // Default to 6 confirmations
	}
}