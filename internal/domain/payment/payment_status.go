package payment

// PaymentStatus represents the current state of a payment
type PaymentStatus string

const (
	StatusPending    PaymentStatus = "PENDING"     // Payment initiated, waiting for cryptocurrency
	StatusConfirming PaymentStatus = "CONFIRMING" // Transaction detected, waiting for confirmations
	StatusConfirmed  PaymentStatus = "CONFIRMED"  // Payment confirmed and completed
	StatusFailed     PaymentStatus = "FAILED"     // Payment failed or rejected
	StatusExpired    PaymentStatus = "EXPIRED"    // Payment expired (timeout)
	StatusRefunded   PaymentStatus = "REFUNDED"   // Payment refunded
	StatusCancelled  PaymentStatus = "CANCELLED"  // Payment cancelled by user
)

// IsValid checks if the payment status is valid
func (ps PaymentStatus) IsValid() bool {
	switch ps {
	case StatusPending, StatusConfirming, StatusConfirmed, StatusFailed, StatusExpired, StatusRefunded, StatusCancelled:
		return true
	default:
		return false
	}
}

// IsCompleted checks if the payment is in a completed state
func (ps PaymentStatus) IsCompleted() bool {
	return ps == StatusConfirmed
}

// IsFinal checks if the payment is in a final state (cannot be changed)
func (ps PaymentStatus) IsFinal() bool {
	switch ps {
	case StatusConfirmed, StatusFailed, StatusExpired, StatusRefunded, StatusCancelled:
		return true
	default:
		return false
	}
}

// CanBeRefunded checks if the payment can be refunded
func (ps PaymentStatus) CanBeRefunded() bool {
	return ps == StatusConfirmed
}

// CanBeCancelled checks if the payment can be cancelled
func (ps PaymentStatus) CanBeCancelled() bool {
	return ps == StatusPending || ps == StatusConfirming
}