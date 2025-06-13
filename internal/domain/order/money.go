package order

import "errors"

// Money represents a monetary value with currency
type Money struct {
	Amount   float64
	Currency string
}

// NewMoney creates a new Money value object with validation
func NewMoney(amount float64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, errors.New("amount cannot be negative")
	}
	
	if currency == "" {
		return Money{}, errors.New("currency cannot be empty")
	}
	
	return Money{Amount: amount, Currency: currency}, nil
}
