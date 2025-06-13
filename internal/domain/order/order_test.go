package order

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Test helper functions for order.go testing only

// createTestMoney creates a Money value for testing
func createTestMoney(amount float64) Money {
    return Money{Amount: amount, Currency: "USD"}
}

// createTestItem creates an OrderItem for testing (simplified for order.go testing)
func createTestItem() OrderItem {
    return OrderItem{
        ProductID:  uuid.New(),
        Quantity:   2,
        UnitPrice:  createTestMoney(10.0),
        Subtotal:   createTestMoney(20.0),
    }
}

// createTestItemWithID creates an OrderItem with specific product ID for testing
func createTestItemWithID(productID uuid.UUID) OrderItem {
    return OrderItem{
        ProductID:  productID,
        Quantity:   1,
        UnitPrice:  createTestMoney(15.0),
        Subtotal:   createTestMoney(15.0),
    }
}

// createTestOrder creates an Order with one item for testing
func createTestOrder() (*Order, error) {
    return NewOrder("customer123", []OrderItem{createTestItem()})
}

// Tests for Order entity functions only (order.go)

func TestNewOrder(t *testing.T) {
    t.Run("success with valid inputs", func(t *testing.T) {
        // Arrange
        customerID := "customer123"
        item := createTestItem()
        
        // Act
        order, err := NewOrder(customerID, []OrderItem{item})
        
        // Assert
        assert.NoError(t, err)
        assert.NotNil(t, order)
        assert.Equal(t, customerID, order.CustomerID)
        assert.Equal(t, StatusCreated, order.Status)
        assert.Equal(t, createTestMoney(20.0), order.TotalAmount)
        assert.Len(t, order.Items, 1)
        assert.False(t, order.CreatedAt.IsZero())
        assert.False(t, order.UpdatedAt.IsZero())
        assert.Nil(t, order.PaymentID)
        assert.Nil(t, order.CompletedAt)
    })
    
    t.Run("error with empty customer ID", func(t *testing.T) {
        // Act
        order, err := NewOrder("", []OrderItem{createTestItem()})
        
        // Assert
        assert.Error(t, err)
        assert.Equal(t, ErrEmptyCustomerID, err)
        assert.Nil(t, order)
    })
    
    t.Run("error with no items", func(t *testing.T) {
        // Act
        order, err := NewOrder("customer123", []OrderItem{})
        
        // Assert
        assert.Error(t, err)
        assert.Equal(t, ErrOrderWithoutItems, err)
        assert.Nil(t, order)
    })
}

func TestOrderStatusTransitions(t *testing.T) {
    t.Run("mark as paid", func(t *testing.T) {
        // Arrange
        order, _ := createTestOrder()
        oldUpdateTime := order.UpdatedAt
        time.Sleep(1 * time.Millisecond) // Ensure time difference
        
        // Act
        err := order.MarkAsPaid("payment123")
        
        // Assert
        assert.NoError(t, err)
        assert.Equal(t, StatusPaid, order.Status)
        assert.Equal(t, "payment123", *order.PaymentID)
        assert.True(t, order.UpdatedAt.After(oldUpdateTime))
    })
	
    
    t.Run("cannot mark as paid twice", func(t *testing.T) {
        // Arrange
        order, _ := createTestOrder()
        _ = order.MarkAsPaid("payment123")
        
        // Act
        err := order.MarkAsPaid("payment456")
        
        // Assert
        assert.Error(t, err)
        assert.Equal(t, ErrInvalidStatusTransition, err)
        assert.Equal(t, "payment123", *order.PaymentID) // Should remain unchanged
    })
    
    t.Run("mark as fulfilled", func(t *testing.T) {
        // Arrange
        order, _ := createTestOrder()
        _ = order.MarkAsPaid("payment123") // First mark as paid
        oldUpdateTime := order.UpdatedAt
        time.Sleep(1 * time.Millisecond) // Ensure time difference
        
        // Act
        err := order.MarkAsFulfilled()
        
        // Assert
        assert.NoError(t, err)
        assert.Equal(t, StatusFulfilled, order.Status)
        assert.NotNil(t, order.CompletedAt)
        assert.True(t, order.UpdatedAt.After(oldUpdateTime))
    })
    
    t.Run("cannot fulfill unpaid order", func(t *testing.T) {
        // Arrange
        order, _ := createTestOrder() // Order is just created, not paid
        
        // Act
        err := order.MarkAsFulfilled()
        
        // Assert
        assert.Error(t, err)
        assert.Equal(t, ErrInvalidStatusTransition, err)
        assert.Nil(t, order.CompletedAt)
    })
}

func TestOrderCancel(t *testing.T) {
    t.Run("cancel created order", func(t *testing.T) {
        // Arrange
        order, _ := createTestOrder()
        oldUpdateTime := order.UpdatedAt
        time.Sleep(1 * time.Millisecond)
        
        // Act
        err := order.Cancel()
        
        // Assert
        assert.NoError(t, err)
        assert.Equal(t, StatusCancelled, order.Status)
        assert.True(t, order.UpdatedAt.After(oldUpdateTime))
    })
    
    t.Run("cancel paid order", func(t *testing.T) {
        // Arrange
        order, _ := createTestOrder()
        _ = order.MarkAsPaid("payment123")
        
        // Act
        err := order.Cancel()
        
        // Assert
        assert.NoError(t, err)
        assert.Equal(t, StatusCancelled, order.Status)
    })
    
    t.Run("cannot cancel fulfilled order", func(t *testing.T) {
        // Arrange
        order, _ := createTestOrder()
        _ = order.MarkAsPaid("payment123")
        _ = order.MarkAsFulfilled()
        
        // Act
        err := order.Cancel()
        
        // Assert
        assert.Error(t, err)
        assert.Equal(t, ErrCannotCancelFulfilledOrder, err)
        assert.Equal(t, StatusFulfilled, order.Status) // Status should remain unchanged
    })
}

func TestOrderItemManagement(t *testing.T) {
    t.Run("add item to order", func(t *testing.T) {
        // Arrange
        order, _ := createTestOrder() // Starts with 1 item
        newItem := createTestItem()   // Create another item
        newItem.ProductID = uuid.New() // Make it a different product
        oldUpdateTime := order.UpdatedAt
        time.Sleep(1 * time.Millisecond)
        
        // Act
        err := order.AddItem(newItem)
        
        // Assert
        assert.NoError(t, err)
        assert.Len(t, order.Items, 2)
        assert.Equal(t, createTestMoney(40.0), order.TotalAmount) // 2 items at $20 each
        assert.True(t, order.UpdatedAt.After(oldUpdateTime))
    })
    
    t.Run("cannot add item after payment", func(t *testing.T) {
        // Arrange
        order, _ := createTestOrder()
        _ = order.MarkAsPaid("payment123")
        originalItemCount := len(order.Items)
        
        // Act
        err := order.AddItem(createTestItem())
        
        // Assert
        assert.Error(t, err)
        assert.Equal(t, ErrCannotModifyOrder, err)
        assert.Len(t, order.Items, originalItemCount) // Items should remain unchanged
    })
    
    t.Run("remove item from order", func(t *testing.T) {
        // Arrange
        item1 := createTestItemWithID(uuid.New())
        item2 := createTestItemWithID(uuid.New())
        order, _ := NewOrder("customer123", []OrderItem{item1, item2})
        oldUpdateTime := order.UpdatedAt
        time.Sleep(1 * time.Millisecond)
        
        // Act
        err := order.RemoveItem(item1.ProductID)
        
        // Assert
        assert.NoError(t, err)
        assert.Len(t, order.Items, 1)
        assert.Equal(t, item2.ProductID, order.Items[0].ProductID)
        assert.True(t, order.UpdatedAt.After(oldUpdateTime))
    })
    
    t.Run("cannot remove non-existent item", func(t *testing.T) {
        // Arrange
        order, _ := createTestOrder()
        nonExistentProductID := uuid.New()
        
        // Act
        err := order.RemoveItem(nonExistentProductID)
        
        // Assert
        assert.Error(t, err)
        assert.Equal(t, ErrItemNotFound, err)
    })
    
    t.Run("removing last item cancels order", func(t *testing.T) {
        // Arrange
        item := createTestItem()
        order, _ := NewOrder("customer123", []OrderItem{item})
        
        // Act
        err := order.RemoveItem(item.ProductID)
        
        // Assert
        assert.NoError(t, err)
        assert.Equal(t, StatusCancelled, order.Status)
        assert.Len(t, order.Items, 0) // Items should be empty after cancellation
    })
    
    t.Run("cannot modify order after fulfillment", func(t *testing.T) {
        // Arrange
        order, _ := createTestOrder()
        _ = order.MarkAsPaid("payment123")
        _ = order.MarkAsFulfilled()
        
        // Act
        addErr := order.AddItem(createTestItem())
        removeErr := order.RemoveItem(order.Items[0].ProductID)
        
        // Assert
        assert.Error(t, addErr)
        assert.Equal(t, ErrCannotModifyOrder, addErr)
        assert.Error(t, removeErr)
        assert.Equal(t, ErrCannotModifyOrder, removeErr)
    })
}

func TestCalculateTotalAmount(t *testing.T) {
    t.Run("calculate total with multiple items", func(t *testing.T) {
        // Arrange
        item1 := createTestItemWithID(uuid.New())
        item2 := createTestItemWithID(uuid.New())
        items := []OrderItem{item1, item2}
        
        // Act
        total, err := calculateTotalAmount(items)
        
        // Assert
        assert.NoError(t, err)
        assert.Equal(t, Money{Amount: 30.0, Currency: "USD"}, total) // 15 + 15
    })
    
    t.Run("order must have at least one item", func(t *testing.T) {
        // Act
        _, err := calculateTotalAmount([]OrderItem{})
        
        // Assert
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "order must have at least one item")
    })
    
    t.Run("error with inconsistent currencies", func(t *testing.T) {
        // Arrange
        item1 := OrderItem{
            ProductID: uuid.New(),
            Quantity:  1,
            UnitPrice: Money{Amount: 10.0, Currency: "USD"},
            Subtotal:  Money{Amount: 10.0, Currency: "USD"},
        }
        item2 := OrderItem{
            ProductID: uuid.New(),
            Quantity:  1,
            UnitPrice: Money{Amount: 10.0, Currency: "EUR"},
            Subtotal:  Money{Amount: 10.0, Currency: "EUR"},
        }
        
        // Act
        _, err := calculateTotalAmount([]OrderItem{item1, item2})
        
        // Assert
        assert.Error(t, err)
        assert.Equal(t, ErrInconsistentCurrency, err)
    })
}