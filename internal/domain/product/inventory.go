package product

// Inventory represents product stock information
type Inventory struct {
	Quantity         int // Total quantity in stock
	ReservedQuantity int // Quantity reserved for pending orders
	MinimumStock     int // Minimum stock level for reordering
}

// NewInventory creates a new inventory with validation
func NewInventory(quantity, reserved, minimum int) (Inventory, error) {
	if quantity < 0 {
		return Inventory{}, ErrNegativeStock
	}
	
	if reserved < 0 {
		return Inventory{}, ErrNegativeReserved
	}
	
	if minimum < 0 {
		return Inventory{}, ErrNegativeMinimum
	}
	
	if reserved > quantity {
		return Inventory{}, ErrReservedExceedsTotal
	}
	
	return Inventory{
		Quantity:         quantity,
		ReservedQuantity: reserved,
		MinimumStock:     minimum,
	}, nil
}

// AvailableQuantity returns the quantity available for new orders
func (inv Inventory) AvailableQuantity() int {
	return inv.Quantity - inv.ReservedQuantity
}

// IsAvailable checks if the requested quantity is available
func (inv Inventory) IsAvailable(requestedQuantity int) bool {
	if requestedQuantity <= 0 {
		return false
	}
	return inv.AvailableQuantity() >= requestedQuantity
}

// IsLowStock checks if inventory is below minimum stock level
func (inv Inventory) IsLowStock() bool {
	return inv.Quantity <= inv.MinimumStock
}

// IsOutOfStock checks if no inventory is available
func (inv Inventory) IsOutOfStock() bool {
	return inv.AvailableQuantity() <= 0
}

// AddStock increases the total quantity
func (inv *Inventory) AddStock(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	
	inv.Quantity += quantity
	return nil
}

// ReserveStock reserves quantity for an order
func (inv *Inventory) ReserveStock(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	
	if !inv.IsAvailable(quantity) {
		return ErrInsufficientStock
	}
	
	inv.ReservedQuantity += quantity
	return nil
}

// ReleaseStock releases reserved quantity (e.g., when order is cancelled)
func (inv *Inventory) ReleaseStock(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	
	if quantity > inv.ReservedQuantity {
		return ErrCannotReleaseMoreThanReserved
	}
	
	inv.ReservedQuantity -= quantity
	return nil
}

// FulfillStock removes stock when order is fulfilled
func (inv *Inventory) FulfillStock(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	
	if quantity > inv.ReservedQuantity {
		return ErrCannotFulfillMoreThanReserved
	}
	
	inv.Quantity -= quantity
	inv.ReservedQuantity -= quantity
	return nil
}

// UpdateMinimumStock updates the minimum stock threshold
func (inv *Inventory) UpdateMinimumStock(minimum int) error {
	if minimum < 0 {
		return ErrNegativeMinimum
	}
	
	inv.MinimumStock = minimum
	return nil
}