package payment

import "errors"

// Payment domain errors organized by category

// === Validation Errors ===
var (
	ErrInvalidAmount           = errors.New("payment amount must be positive")
	ErrInvalidCurrency         = errors.New("payment currency is invalid")
	ErrEmptyOrderID            = errors.New("order ID cannot be empty")
	ErrEmptyPaymentID          = errors.New("payment ID cannot be empty")
	ErrInvalidCryptoCurrency   = errors.New("cryptocurrency is not supported")
	ErrInvalidWalletAddress    = errors.New("wallet address is invalid")
	ErrInvalidTransactionHash  = errors.New("transaction hash is invalid")
)

// === Business Rule Errors ===
var (
	ErrPaymentAlreadyConfirmed = errors.New("payment is already confirmed")
	ErrPaymentNotFound         = errors.New("payment not found")
	ErrInsufficientAmount      = errors.New("payment amount is insufficient")
	ErrExcessiveAmount         = errors.New("payment amount exceeds expected amount")
	ErrPaymentExpired          = errors.New("payment has expired")
	ErrOrderAlreadyPaid        = errors.New("order is already paid")
)

// === State Transition Errors ===
var (
	ErrCannotConfirmPayment    = errors.New("payment cannot be confirmed")
	ErrCannotCancelPayment     = errors.New("payment cannot be cancelled")
	ErrCannotRefundPayment     = errors.New("payment cannot be refunded")
	ErrCannotUpdateFinalPayment = errors.New("cannot update finalized payment")
	ErrInvalidStatusTransition = errors.New("invalid payment status transition")
)

// === External Service Errors ===
var (
	ErrNowPaymentsAPIError     = errors.New("NowPayments API error")
	ErrPaymentServiceTimeout   = errors.New("payment service timeout")
	ErrWebhookValidationFailed = errors.New("webhook signature validation failed")
	ErrInvalidWebhookPayload   = errors.New("invalid webhook payload")
)

// === Cryptocurrency Errors ===
var (
	ErrUnsupportedCrypto       = errors.New("cryptocurrency not supported")
	ErrInvalidCryptoAmount     = errors.New("cryptocurrency amount is invalid")
	ErrNetworkCongestion       = errors.New("cryptocurrency network is congested")
	ErrInsufficientConfirmations = errors.New("insufficient blockchain confirmations")
)

// === Refund Errors ===
var (
	ErrRefundAlreadyProcessed  = errors.New("refund already processed")
	ErrPartialRefundNotAllowed = errors.New("partial refund not allowed")
	ErrRefundAmountExceedsPayment = errors.New("refund amount exceeds original payment")
	ErrRefundDeadlineExpired   = errors.New("refund deadline has expired")
)