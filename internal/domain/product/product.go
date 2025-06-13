package product

import (
	"time"

	"github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/order"
	"github.com/google/uuid"
)

// Product represents a product that can be sold
type Product struct {
	ID          uuid.UUID     // Unique identifier
	Name        string        // Product name
	Description string        // Product description
	SKU         string        // Stock Keeping Unit (unique)
	Price       order.Money   // Product price (reusing Money from order domain)
	Category    Category      // Product category
	Inventory   Inventory     // Stock information
	Status      ProductStatus // Current product status
	CreatedAt   time.Time     // When product was created
	UpdatedAt   time.Time     // When product was last updated
}

// NewProduct creates a new product with validation
func NewProduct(name, description, sku string, price order.Money, category Category, inventory Inventory) (*Product, error) {
	// Validate required fields
	if name == "" {
		return nil, ErrEmptyName
	}
	
	if sku == "" {
		return nil, ErrEmptySKU
	}
	
	if price.Amount <= 0 {
		return nil, ErrInvalidPrice
	}
	
	if price.Currency == "" {
		return nil, ErrInvalidCurrency
	}
	
	// Create product
	now := time.Now()
	product := &Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		SKU:         sku,
		Price:       price,
		Category:    category,
		Inventory:   inventory,
		Status:      StatusInactive, // New products start as inactive
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	
	return product, nil
}

// UpdatePrice updates the product price
func (p *Product) UpdatePrice(newPrice order.Money) error {
	if newPrice.Amount <= 0 {
		return ErrInvalidPrice
	}
	
	if newPrice.Currency == "" {
		return ErrInvalidCurrency
	}
	
	// Only allow price updates for active/inactive products
	if p.Status == StatusDiscontinued {
		return ErrCannotUpdateDiscontinued
	}
	
	p.Price = newPrice
	p.UpdatedAt = time.Now()
	
	return nil
}

// UpdateDescription updates the product description
func (p *Product) UpdateDescription(description string) error {
	// Only allow updates for active/inactive products
	if p.Status == StatusDiscontinued {
		return ErrCannotUpdateDiscontinued
	}
	
	p.Description = description
	p.UpdatedAt = time.Now()
	
	return nil
}

// UpdateCategory updates the product category
func (p *Product) UpdateCategory(category Category) error {
	// Only allow updates for active/inactive products
	if p.Status == StatusDiscontinued {
		return ErrCannotUpdateDiscontinued
	}
	
	p.Category = category
	p.UpdatedAt = time.Now()
	
	return nil
}

// Activate makes the product available for purchase
func (p *Product) Activate() error {
	// Cannot activate if no stock available
	if p.Inventory.AvailableQuantity() <= 0 {
		return ErrCannotActivateWithoutStock
	}
	
	// Only allow activation from inactive status
	if p.Status == StatusDiscontinued {
		return ErrCannotActivateDiscontinued
	}
	
	p.Status = StatusActive
	p.UpdatedAt = time.Now()
	
	return nil
}

// Deactivate makes the product unavailable for purchase
func (p *Product) Deactivate() error {
	// Only active products can be deactivated
	if p.Status != StatusActive {
		return ErrCannotDeactivateInactive
	}
	
	p.Status = StatusInactive
	p.UpdatedAt = time.Now()
	
	return nil
}

// MarkOutOfStock marks product as out of stock
func (p *Product) MarkOutOfStock() {
	p.Status = StatusOutOfStock
	p.UpdatedAt = time.Now()
}

// Discontinue permanently discontinues the product
func (p *Product) Discontinue() error {
	// Cannot discontinue if there are reserved quantities
	if p.Inventory.ReservedQuantity > 0 {
		return ErrCannotDiscontinueWithReserved
	}
	
	p.Status = StatusDiscontinued
	p.UpdatedAt = time.Now()
	
	return nil
}

// AddStock increases product inventory
func (p *Product) AddStock(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	
	err := p.Inventory.AddStock(quantity)
	if err != nil {
		return err
	}
	
	// If product was out of stock and now has stock, reactivate it
	if p.Status == StatusOutOfStock && p.Inventory.AvailableQuantity() > 0 {
		p.Status = StatusInactive // Set to inactive, requires manual activation
	}
	
	p.UpdatedAt = time.Now()
	return nil
}

// ReserveStock reserves inventory for an order
func (p *Product) ReserveStock(quantity int) error {
	// Only active products can have stock reserved
	if !p.Status.CanBeOrdered() {
		return ErrProductNotActive
	}
	
	err := p.Inventory.ReserveStock(quantity)
	if err != nil {
		return err
	}
	
	// Check if product should be marked as out of stock
	if p.Inventory.AvailableQuantity() == 0 {
		p.MarkOutOfStock()
	}
	
	p.UpdatedAt = time.Now()
	return nil
}

// ReleaseStock releases reserved inventory (e.g., when order is cancelled)
func (p *Product) ReleaseStock(quantity int) error {
	err := p.Inventory.ReleaseStock(quantity)
	if err != nil {
		return err
	}
	
	// If product was out of stock and now has available stock, make it active
	if p.Status == StatusOutOfStock && p.Inventory.AvailableQuantity() > 0 {
		p.Status = StatusActive
	}
	
	p.UpdatedAt = time.Now()
	return nil
}

// FulfillStock removes inventory when order is fulfilled
func (p *Product) FulfillStock(quantity int) error {
	err := p.Inventory.FulfillStock(quantity)
	if err != nil {
		return err
	}
	
	p.UpdatedAt = time.Now()
	return nil
}

// UpdateInventoryMinimum updates the minimum stock threshold
func (p *Product) UpdateInventoryMinimum(minimum int) error {
	err := p.Inventory.UpdateMinimumStock(minimum)
	if err != nil {
		return err
	}
	
	p.UpdatedAt = time.Now()
	return nil
}

// IsAvailableForOrder checks if product can be ordered
func (p *Product) IsAvailableForOrder(quantity int) bool {
	return p.Status.CanBeOrdered() && p.Inventory.IsAvailable(quantity)
}

// IsLowStock checks if product inventory is low
func (p *Product) IsLowStock() bool {
	return p.Inventory.IsLowStock()
}

// GetAvailableQuantity returns available stock quantity
func (p *Product) GetAvailableQuantity() int {
	return p.Inventory.AvailableQuantity()
}

// GetTotalQuantity returns total stock quantity
func (p *Product) GetTotalQuantity() int {
	return p.Inventory.Quantity
}

// GetReservedQuantity returns reserved stock quantity
func (p *Product) GetReservedQuantity() int {
	return p.Inventory.ReservedQuantity
}

// IsOutOfStock checks if product is out of stock
func (p *Product) IsOutOfStock() bool {
	return p.Inventory.IsOutOfStock()
}

// CanBeDeleted checks if product can be safely deleted
func (p *Product) CanBeDeleted() bool {
	// Cannot delete active products or products with reserved stock
	return p.Status != StatusActive && p.Inventory.ReservedQuantity == 0
}

// GetCategoryPath returns the category breadcrumb path
func (p *Product) GetCategoryPath(separator string) string {
	// This would typically be resolved by a domain service that knows about category hierarchies
	// For now, return the category name
	return p.Category.Name
}

