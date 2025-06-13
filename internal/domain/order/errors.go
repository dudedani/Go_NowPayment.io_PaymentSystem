package order

import "errors"

// Error definitions
var (
	ErrEmptyCustomerID         = errors.New("customer ID cannot be empty")
	ErrOrderWithoutItems       = errors.New("order must have at least one item")
	ErrInvalidQuantity         = errors.New("quantity must be greater than zero")
	ErrInvalidStatusTransition = errors.New("invalid order status transition")
	ErrCannotModifyOrder       = errors.New("cannot modify order after payment")
	ErrItemNotFound            = errors.New("item not found in order")
	ErrOrderMustHaveItems      = errors.New("order must have at least one item")
	ErrInconsistentCurrency    = errors.New("all items must have the same currency")
	ErrCannotCancelFulfilledOrder = errors.New("cannot cancel a fulfilled order")
)
