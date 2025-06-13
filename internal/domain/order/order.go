package order

import (
	"errors"
	"time"

	"github.com/google/uuid"
	//"golang.org/x/text/currency"
)

// Order represents an order in the system
type Order struct {
    ID            uuid.UUID
    CustomerID    string
    Items         []OrderItem
    Status        OrderStatus
    TotalAmount   Money
    CreatedAt     time.Time
    UpdatedAt     time.Time
    PaymentID     *string // Optional, set when payment is created
    CompletedAt   *time.Time
}

// - NewOrder creates a new order with the given customer ID and items
func NewOrder(customerID string, items []OrderItem)(*Order,error){
    if customerID==""{
        return nil,ErrEmptyCustomerID
    }

    if len(items)==0{
        return nil,ErrOrderWithoutItems
    }

    totalAmount,err:= calculateTotalAmount(items)
    if err != nil {
        return nil, err
    }

    now:=time.Now()
    return &Order{
        ID:          uuid.New(),
        CustomerID:  customerID,
        Items:       items,
        Status:      StatusCreated,
        TotalAmount: totalAmount,
        CreatedAt:   now,
        UpdatedAt:   now,
    },nil
}


//calculateTotalAmount calculates the total amount of the order
func calculateTotalAmount(items []OrderItem) (Money, error) {
    if len(items) == 0 {
        return Money{}, errors.New("order must have at least one item")
    }

    currency:= items[0].UnitPrice.Currency
    var totalAmount float64

    //Sum up the subtotal of each item and verify currency consistency
    for _, item := range items {
        if item.UnitPrice.Currency != currency {
            return Money{}, ErrInconsistentCurrency 
        }
        totalAmount += item.Subtotal.Amount
    }
    return Money{
        Amount:   totalAmount,
        Currency: currency,
    }, nil
}


// - MarkAsPaid
func (o *Order) MarkAsPaid(paymentID string) error {
 if paymentID == "" {
        return errors.New("payment ID cannot be empty")
    }
    
    // Only allow transition from Created to Paid
    if o.Status != StatusCreated {
        return ErrInvalidStatusTransition
    }
    
    // Update order state
    o.Status = StatusPaid
    o.PaymentID = &paymentID
    o.UpdatedAt = time.Now()
    
    return nil
}
// - MarkAsFulfilled
func (o *Order) MarkAsFulfilled() error {

    //Only allow transition from Paid to Fulfilled
    if o.Status != StatusPaid {
        return ErrInvalidStatusTransition
    }
    // Update order state
    o.Status = StatusFulfilled
    now := time.Now()
    o.CompletedAt = &now
    o.UpdatedAt = now
    return nil
}
// - Cancel
func (o*Order)Cancel()error{
    // Cannot cancel fulfilled orders (already shipped/completed)
    if o.Status == StatusFulfilled {
        return ErrCannotCancelFulfilledOrder
    }
    
    // Only allow cancellation from Created or Paid status
    if o.Status != StatusCreated && o.Status != StatusPaid {
        return ErrInvalidStatusTransition
    }
    
    // Update order state
    o.Status = StatusCancelled
    o.UpdatedAt = time.Now()
    
    return nil

}

// AddItem adds an item to the order if the order can be modified
func (o *Order) AddItem(item OrderItem) error {
    // Cannot modify orders that are already paid and fulfilled
    if o.Status != StatusCreated {
        return ErrCannotModifyOrder
    }
    
    // Check if item with same ProductID already exists
    for i, existingItem := range o.Items {
        if existingItem.ProductID == item.ProductID {
            // Update existing item quantity and subtotal
            o.Items[i].Quantity += item.Quantity
            o.Items[i].Subtotal = Money{
                Amount:   o.Items[i].UnitPrice.Amount * float64(o.Items[i].Quantity),
                Currency: o.Items[i].UnitPrice.Currency,
            }
            
            // Recalculate total amount
            totalAmount, err := calculateTotalAmount(o.Items)
            if err != nil {
                return err
            }
            
            o.TotalAmount = totalAmount
            o.UpdatedAt = time.Now()
            return nil
        }
    }
    
    // Add new item if ProductID doesn't exist
    o.Items = append(o.Items, item)
    
    // Recalculate total amount
    totalAmount, err := calculateTotalAmount(o.Items)
    if err != nil {
        return err
    }
    
    o.TotalAmount = totalAmount
    o.UpdatedAt = time.Now()
    
    return nil
}
// - RemoveItem
func (o*Order) RemoveItem(productID uuid.UUID) error{
    //cannot modify orders that are already paid and fulfilled
    if o.Status != StatusCreated{
        return ErrCannotModifyOrder
    }

    //Find the item to remove
    for i, existingItem:= range o.Items{
        if existingItem.ProductID == productID{
            //Item found -> Remove the Item from the slice
            o.Items = append(o.Items[:i], o.Items[i+1:]...)

            // Check if order is now empty - if so, cancel the order
            if len(o.Items) == 0 {
                return o.Cancel()
            }
            
            // Recalculate total amount
            totalAmount, err := calculateTotalAmount(o.Items)
            if err != nil {
                return err
            }
            
            o.TotalAmount = totalAmount
            o.UpdatedAt = time.Now()
            
            return nil
        }
    }
    return ErrItemNotFound
}