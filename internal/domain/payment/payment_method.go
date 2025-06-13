package payment

import (
	"time"
)

// PaymentMethod represents the method used for payment
type PaymentMethod struct {
	Type           PaymentType     // Type of payment (crypto)
	CryptoCurrency CryptoCurrency  // Cryptocurrency used
	WalletAddress  string          // Destination wallet address
	ExpiresAt      time.Time       // When this payment method expires
}

// PaymentType represents the type of payment
type PaymentType string

const (
	TypeCryptocurrency PaymentType = "CRYPTOCURRENCY"
	// Future payment types can be added here
	// TypeCreditCard     PaymentType = "CREDIT_CARD"
	// TypeBankTransfer   PaymentType = "BANK_TRANSFER"
)

// NewPaymentMethod creates a new payment method
func NewPaymentMethod(cryptoSymbol string, walletAddress string, expirationMinutes int) (PaymentMethod, error) {
	if walletAddress == "" {
		return PaymentMethod{}, ErrInvalidWalletAddress
	}
	
	crypto, err := GetCryptoCurrencyBySymbol(cryptoSymbol)
	if err != nil {
		return PaymentMethod{}, err
	}
	
	if !crypto.IsActive {
		return PaymentMethod{}, ErrUnsupportedCrypto
	}
	
	expiresAt := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)
	
	return PaymentMethod{
		Type:           TypeCryptocurrency,
		CryptoCurrency: crypto,
		WalletAddress:  walletAddress,
		ExpiresAt:      expiresAt,
	}, nil
}

// IsExpired checks if the payment method has expired
func (pm PaymentMethod) IsExpired() bool {
	return time.Now().After(pm.ExpiresAt)
}

// TimeUntilExpiry returns the duration until the payment method expires
func (pm PaymentMethod) TimeUntilExpiry() time.Duration {
	if pm.IsExpired() {
		return 0
	}
	return time.Until(pm.ExpiresAt)
}

// ValidateAmount validates if the amount is acceptable for this payment method
func (pm PaymentMethod) ValidateAmount(amount float64) error {
	return pm.CryptoCurrency.ValidateAmount(amount)
}

// GetCryptoSymbol returns the cryptocurrency symbol
func (pm PaymentMethod) GetCryptoSymbol() string {
	return pm.CryptoCurrency.Symbol
}

// GetWalletAddress returns the wallet address
func (pm PaymentMethod) GetWalletAddress() string {
	return pm.WalletAddress
}

// IsActive checks if the payment method is still active (not expired and crypto is supported)
func (pm PaymentMethod) IsActive() bool {
	return !pm.IsExpired() && pm.CryptoCurrency.IsActive
}