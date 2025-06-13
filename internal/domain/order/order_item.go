package order

import (
    "github.com/google/uuid"
)

// OrderItem represents an item within an order
type OrderItem struct {
    ProductID  uuid.UUID
    Quantity   int
    UnitPrice  Money
    Subtotal   Money
}

// NewOrderItem creates a new order item with validation
func NewOrderItem(productID uuid.UUID, quantity int, unitPrice Money) (OrderItem, error) {
    if quantity <= 0 {
        return OrderItem{}, ErrInvalidQuantity
    }

    // Calculate subtotal
    subtotalAmount := unitPrice.Amount * float64(quantity)
    subtotal := Money{
        Amount:   subtotalAmount,
        Currency: unitPrice.Currency,
    }

    return OrderItem{
        ProductID: productID,
        Quantity:  quantity,
        UnitPrice: unitPrice,
        Subtotal:  subtotal,
    }, nil
}