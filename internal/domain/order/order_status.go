package order

// OrderStatus represents the current state of an order
type OrderStatus string

const (
	StatusCreated   OrderStatus = "CREATED"
	StatusPaid      OrderStatus = "PAID"
	StatusFulfilled OrderStatus = "FULFILLED"  
	StatusCancelled OrderStatus = "CANCELLED"
)
