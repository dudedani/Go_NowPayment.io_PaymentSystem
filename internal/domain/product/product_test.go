package product

import (
	"testing"
	"time"

	"github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/order"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Test helper functions

// createTestInventory creates a valid inventory for testing
func createTestInventory() Inventory {
	inventory, _ := NewInventory(100, 0, 5) // Start with no reserved stock
	return inventory
}

// createTestCategory creates a valid category for testing
func createTestCategory() Category {
	category, _ := NewCategory("Electronics", "Electronic products", nil)
	return category
}

// Tests for ProductStatus (already implemented)

func TestProductStatus(t *testing.T) {
	t.Run("valid statuses", func(t *testing.T) {
		validStatuses := []ProductStatus{
			StatusActive, StatusInactive, StatusOutOfStock, StatusDiscontinued,
		}
		
		for _, status := range validStatuses {
			assert.True(t, status.IsValid())
		}
	})
	
	t.Run("invalid status", func(t *testing.T) {
		invalidStatus := ProductStatus("INVALID")
		assert.False(t, invalidStatus.IsValid())
	})
	
	t.Run("can be ordered", func(t *testing.T) {
		assert.True(t, StatusActive.CanBeOrdered())
		assert.False(t, StatusInactive.CanBeOrdered())
		assert.False(t, StatusOutOfStock.CanBeOrdered())
		assert.False(t, StatusDiscontinued.CanBeOrdered())
	})
}

// Tests for Inventory

func TestInventory(t *testing.T) {
	t.Run("create valid inventory", func(t *testing.T) {
		inventory, err := NewInventory(100, 10, 5)
		
		assert.NoError(t, err)
		assert.Equal(t, 100, inventory.Quantity)
		assert.Equal(t, 10, inventory.ReservedQuantity)
		assert.Equal(t, 5, inventory.MinimumStock)
		assert.Equal(t, 90, inventory.AvailableQuantity())
	})
	
	t.Run("cannot create inventory with negative quantity", func(t *testing.T) {
		_, err := NewInventory(-10, 0, 5)
		
		assert.Error(t, err)
		assert.Equal(t, ErrNegativeStock, err)
	})
	
	t.Run("cannot create inventory with negative reserved", func(t *testing.T) {
		_, err := NewInventory(100, -5, 5)
		
		assert.Error(t, err)
		assert.Equal(t, ErrNegativeReserved, err)
	})
	
	t.Run("cannot create inventory with reserved exceeding total", func(t *testing.T) {
		_, err := NewInventory(100, 150, 5)
		
		assert.Error(t, err)
		assert.Equal(t, ErrReservedExceedsTotal, err)
	})
	
	t.Run("reserve stock", func(t *testing.T) {
		inventory, _ := NewInventory(100, 10, 5)
		
		err := inventory.ReserveStock(20)
		
		assert.NoError(t, err)
		assert.Equal(t, 30, inventory.ReservedQuantity)
		assert.Equal(t, 70, inventory.AvailableQuantity())
	})
	
	t.Run("cannot reserve more than available", func(t *testing.T) {
		inventory, _ := NewInventory(100, 10, 5)
		
		err := inventory.ReserveStock(95) // Only 90 available
		
		assert.Error(t, err)
		assert.Equal(t, ErrInsufficientStock, err)
	})
	
	t.Run("release stock", func(t *testing.T) {
		inventory, _ := NewInventory(100, 30, 5)
		
		err := inventory.ReleaseStock(10)
		
		assert.NoError(t, err)
		assert.Equal(t, 20, inventory.ReservedQuantity)
		assert.Equal(t, 80, inventory.AvailableQuantity())
	})
	
	t.Run("fulfill stock", func(t *testing.T) {
		inventory, _ := NewInventory(100, 30, 5)
		
		err := inventory.FulfillStock(10)
		
		assert.NoError(t, err)
		assert.Equal(t, 90, inventory.Quantity)
		assert.Equal(t, 20, inventory.ReservedQuantity)
		assert.Equal(t, 70, inventory.AvailableQuantity())
	})
	
	t.Run("check low stock", func(t *testing.T) {
		inventory, _ := NewInventory(5, 0, 10)
		
		assert.True(t, inventory.IsLowStock())
	})
	
	t.Run("check out of stock", func(t *testing.T) {
		inventory, _ := NewInventory(10, 10, 5)
		
		assert.True(t, inventory.IsOutOfStock())
	})
}

// Tests for Category

func TestCategory(t *testing.T) {
	t.Run("create valid category", func(t *testing.T) {
		category, err := NewCategory("Electronics", "Electronic products", nil)
		
		assert.NoError(t, err)
		assert.Equal(t, "Electronics", category.Name)
		assert.Equal(t, "Electronic products", category.Description)
		assert.True(t, category.IsRoot())
		assert.False(t, category.HasParent())
	})
	
	t.Run("cannot create category with empty name", func(t *testing.T) {
		_, err := NewCategory("", "Description", nil)
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidCategoryName, err)
	})
	
	t.Run("create category with parent", func(t *testing.T) {
		parentID := uuid.New()
		category, err := NewCategory("Smartphones", "Mobile phones", &parentID)
		
		assert.NoError(t, err)
		assert.False(t, category.IsRoot())
		assert.True(t, category.HasParent())
		assert.Equal(t, parentID, *category.GetParentID())
	})
	
	t.Run("update category name", func(t *testing.T) {
		category, _ := NewCategory("Electronics", "Electronic products", nil)
		
		err := category.UpdateName("Updated Electronics")
		
		assert.NoError(t, err)
		assert.Equal(t, "Updated Electronics", category.Name)
	})
	
	t.Run("cannot update to empty name", func(t *testing.T) {
		category, _ := NewCategory("Electronics", "Electronic products", nil)
		
		err := category.UpdateName("")
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidCategoryName, err)
	})
	
	t.Run("set and remove parent", func(t *testing.T) {
		category, _ := NewCategory("Smartphones", "Mobile phones", nil)
		parentID := uuid.New()
		
		err := category.SetParent(parentID)
		assert.NoError(t, err)
		assert.True(t, category.HasParent())
		
		category.RemoveParent()
		assert.True(t, category.IsRoot())
	})
	
	t.Run("cannot set self as parent", func(t *testing.T) {
		category, _ := NewCategory("Electronics", "Electronic products", nil)
		
		err := category.SetParent(category.ID)
		
		assert.Error(t, err)
		assert.Equal(t, ErrCircularReference, err)
	})
}

// Tests for CategoryPath

func TestCategoryPath(t *testing.T) {
	t.Run("create breadcrumb", func(t *testing.T) {
		cat1, _ := NewCategory("Electronics", "Root category", nil)
		cat2, _ := NewCategory("Smartphones", "Phone category", &cat1.ID)
		cat3, _ := NewCategory("iPhone", "Apple phones", &cat2.ID)
		
		path := NewCategoryPath([]Category{cat1, cat2, cat3})
		breadcrumb := path.GetBreadcrumb(" > ")
		
		assert.Equal(t, "Electronics > Smartphones > iPhone", breadcrumb)
		assert.Equal(t, 3, path.GetDepth())
		assert.False(t, path.IsEmpty())
	})
	
	t.Run("empty path", func(t *testing.T) {
		path := NewCategoryPath([]Category{})
		
		assert.True(t, path.IsEmpty())
		assert.Equal(t, 0, path.GetDepth())
		assert.Equal(t, "", path.GetBreadcrumb(" > "))
	})
}

// Tests for Product Entity

func TestProduct(t *testing.T) {
	// Helper function to create valid test money
	createTestMoney := func(amount float64, currency string) order.Money {
		money, _ := order.NewMoney(amount, currency)
		return money
	}
	
	// Helper function to create valid test product
	createTestProduct := func() *Product {
		inventory := createTestInventory()
		category := createTestCategory()
		price := createTestMoney(99.99, "USD")
		
		product, _ := NewProduct(
			"iPhone 14",
			"Latest Apple smartphone",
			"IPHONE-14-001",
			price,
			category,
			inventory,
		)
		
		return product
	}
	
	t.Run("create valid product", func(t *testing.T) {
		inventory := createTestInventory()
		category := createTestCategory()
		price := createTestMoney(99.99, "USD")
		
		product, err := NewProduct(
			"iPhone 14",
			"Latest Apple smartphone", 
			"IPHONE-14-001",
			price,
			category,
			inventory,
		)
		
		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, "iPhone 14", product.Name)
		assert.Equal(t, "IPHONE-14-001", product.SKU)
		assert.Equal(t, StatusInactive, product.Status)
		assert.False(t, product.CreatedAt.IsZero())
		assert.False(t, product.UpdatedAt.IsZero())
	})
	
	t.Run("cannot create product with empty name", func(t *testing.T) {
		inventory := createTestInventory()
		category := createTestCategory()
		price := createTestMoney(99.99, "USD")
		
		product, err := NewProduct("", "Description", "SKU-001", price, category, inventory)
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyName, err)
		assert.Nil(t, product)
	})
	
	t.Run("cannot create product with empty SKU", func(t *testing.T) {
		inventory := createTestInventory()
		category := createTestCategory()
		price := createTestMoney(99.99, "USD")
		
		product, err := NewProduct("iPhone", "Description", "", price, category, inventory)
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptySKU, err)
		assert.Nil(t, product)
	})
	
	t.Run("cannot create product with invalid price", func(t *testing.T) {
		inventory := createTestInventory()
		category := createTestCategory()
		price := createTestMoney(-10.0, "USD")
		
		product, err := NewProduct("iPhone", "Description", "SKU-001", price, category, inventory)
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidPrice, err)
		assert.Nil(t, product)
	})
	
	t.Run("update product price", func(t *testing.T) {
		product := createTestProduct()
		newPrice := createTestMoney(149.99, "USD")
		oldUpdateTime := product.UpdatedAt
		time.Sleep(1 * time.Millisecond)
		
		err := product.UpdatePrice(newPrice)
		
		assert.NoError(t, err)
		assert.Equal(t, 149.99, product.Price.Amount)
		assert.True(t, product.UpdatedAt.After(oldUpdateTime))
	})
	
	t.Run("cannot update price of discontinued product", func(t *testing.T) {
		product := createTestProduct()
		product.Status = StatusDiscontinued
		newPrice := createTestMoney(149.99, "USD")
		
		err := product.UpdatePrice(newPrice)
		
		assert.Error(t, err)
		assert.Equal(t, ErrCannotUpdateDiscontinued, err)
	})
	
	t.Run("activate product", func(t *testing.T) {
		product := createTestProduct()
		
		err := product.Activate()
		
		assert.NoError(t, err)
		assert.Equal(t, StatusActive, product.Status)
	})
	
	t.Run("cannot activate product without stock", func(t *testing.T) {
		inventory, _ := NewInventory(0, 0, 5) // No stock
		category := createTestCategory()
		price := createTestMoney(99.99, "USD")
		
		product, _ := NewProduct("iPhone", "Description", "SKU-001", price, category, inventory)
		
		err := product.Activate()
		
		assert.Error(t, err)
		assert.Equal(t, ErrCannotActivateWithoutStock, err)
	})
	
	t.Run("cannot activate discontinued product", func(t *testing.T) {
		product := createTestProduct()
		product.Status = StatusDiscontinued
		
		err := product.Activate()
		
		assert.Error(t, err)
		assert.Equal(t, ErrCannotActivateDiscontinued, err)
	})
	
	t.Run("deactivate product", func(t *testing.T) {
		product := createTestProduct()
		_ = product.Activate() // First activate it
		
		err := product.Deactivate()
		
		assert.NoError(t, err)
		assert.Equal(t, StatusInactive, product.Status)
	})
	
	t.Run("cannot deactivate inactive product", func(t *testing.T) {
		product := createTestProduct() // Starts as inactive
		
		err := product.Deactivate()
		
		assert.Error(t, err)
		assert.Equal(t, ErrCannotDeactivateInactive, err)
	})
	
	t.Run("discontinue product", func(t *testing.T) {
		product := createTestProduct()
		
		err := product.Discontinue()
		
		assert.NoError(t, err)
		assert.Equal(t, StatusDiscontinued, product.Status)
	})
	
	t.Run("cannot discontinue product with reserved stock", func(t *testing.T) {
		product := createTestProduct()
		_ = product.Activate()
		_ = product.ReserveStock(5) // Reserve some stock
		
		err := product.Discontinue()
		
		assert.Error(t, err)
		assert.Equal(t, ErrCannotDiscontinueWithReserved, err)
	})
	
	t.Run("add stock to product", func(t *testing.T) {
		product := createTestProduct()
		originalQuantity := product.GetTotalQuantity()
		
		err := product.AddStock(50)
		
		assert.NoError(t, err)
		assert.Equal(t, originalQuantity+50, product.GetTotalQuantity())
	})
	
	t.Run("reserve stock", func(t *testing.T) {
		product := createTestProduct()
		_ = product.Activate()
		originalAvailable := product.GetAvailableQuantity()
		
		err := product.ReserveStock(10)
		
		assert.NoError(t, err)
		assert.Equal(t, originalAvailable-10, product.GetAvailableQuantity())
		assert.Equal(t, 10, product.GetReservedQuantity())
	})
	
	t.Run("cannot reserve stock from inactive product", func(t *testing.T) {
		product := createTestProduct() // Starts as inactive
		
		err := product.ReserveStock(10)
		
		assert.Error(t, err)
		assert.Equal(t, ErrProductNotActive, err)
	})
	
	t.Run("product goes out of stock after reservation", func(t *testing.T) {
		inventory, _ := NewInventory(10, 0, 5)
		category := createTestCategory()
		price := createTestMoney(99.99, "USD")
		
		product, _ := NewProduct("iPhone", "Description", "SKU-001", price, category, inventory)
		_ = product.Activate()
		
		err := product.ReserveStock(10) // Reserve all available stock
		
		assert.NoError(t, err)
		assert.Equal(t, StatusOutOfStock, product.Status)
	})
	
	t.Run("release stock", func(t *testing.T) {
		product := createTestProduct()
		_ = product.Activate()
		_ = product.ReserveStock(10)
		originalReserved := product.GetReservedQuantity()
		
		err := product.ReleaseStock(5)
		
		assert.NoError(t, err)
		assert.Equal(t, originalReserved-5, product.GetReservedQuantity())
	})
	
	t.Run("fulfill stock", func(t *testing.T) {
		product := createTestProduct()
		_ = product.Activate()
		_ = product.ReserveStock(10)
		originalTotal := product.GetTotalQuantity()
		originalReserved := product.GetReservedQuantity()
		
		err := product.FulfillStock(5)
		
		assert.NoError(t, err)
		assert.Equal(t, originalTotal-5, product.GetTotalQuantity())
		assert.Equal(t, originalReserved-5, product.GetReservedQuantity())
	})
	
	t.Run("check if product is available for order", func(t *testing.T) {
		product := createTestProduct()
		
		// Inactive product should not be available
		assert.False(t, product.IsAvailableForOrder(5))
		
		// Active product with sufficient stock should be available
		_ = product.Activate()
		assert.True(t, product.IsAvailableForOrder(5))
		
		// Active product with insufficient stock should not be available
		assert.False(t, product.IsAvailableForOrder(200))
	})
	
	t.Run("check if product can be deleted", func(t *testing.T) {
		product := createTestProduct()
		
		// Inactive product with no reserved stock can be deleted
		assert.True(t, product.CanBeDeleted())
		
		// Active product cannot be deleted
		_ = product.Activate()
		assert.False(t, product.CanBeDeleted())
		
		// Product with reserved stock cannot be deleted
		_ = product.Deactivate()
		_ = product.Activate()
		_ = product.ReserveStock(5)
		_ = product.Deactivate()
		assert.False(t, product.CanBeDeleted())
	})
	
	t.Run("check low stock", func(t *testing.T) {
		inventory, _ := NewInventory(5, 0, 10) // Quantity (5) <= Minimum (10)
		category := createTestCategory()
		price := createTestMoney(99.99, "USD")
		
		product, _ := NewProduct("iPhone", "Description", "SKU-001", price, category, inventory)
		
		assert.True(t, product.IsLowStock())
	})
	
	t.Run("update description", func(t *testing.T) {
		product := createTestProduct()
		oldUpdateTime := product.UpdatedAt
		time.Sleep(1 * time.Millisecond)
		
		err := product.UpdateDescription("Updated description")
		
		assert.NoError(t, err)
		assert.Equal(t, "Updated description", product.Description)
		assert.True(t, product.UpdatedAt.After(oldUpdateTime))
	})
	
	t.Run("update category", func(t *testing.T) {
		product := createTestProduct()
		newCategory, _ := NewCategory("Smartphones", "Mobile phones", nil)
		oldUpdateTime := product.UpdatedAt
		time.Sleep(1 * time.Millisecond)
		
		err := product.UpdateCategory(newCategory)
		
		assert.NoError(t, err)
		assert.Equal(t, "Smartphones", product.Category.Name)
		assert.True(t, product.UpdatedAt.After(oldUpdateTime))
	})
}