package customer

import (
	"time"

	"github.com/google/uuid"
)

// CustomerStatus represents the status of a customer account
type CustomerStatus string

const (
	StatusActive   CustomerStatus = "ACTIVE"
	StatusInactive CustomerStatus = "INACTIVE"
	StatusSuspended CustomerStatus = "SUSPENDED"
)

// IsValid checks if the customer status is valid
func (cs CustomerStatus) IsValid() bool {
	switch cs {
	case StatusActive, StatusInactive, StatusSuspended:
		return true
	default:
		return false
	}
}

// Customer represents a customer entity (Aggregate Root)
type Customer struct {
	// Identity
	ID string
	
	// Personal Information
	Email     Email
	FirstName string
	LastName  string
	Phone     string
	
	// Account Status
	Status    CustomerStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	
	// Shipping Addresses (managed by the aggregate)
	ShippingAddresses []ShippingAddress
}

// NewCustomer creates a new customer with validation
func NewCustomer(email, firstName, lastName, phone string) (*Customer, error) {
	// Validate and create email
	emailObj, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	
	// Validate required fields
	if firstName == "" {
		return nil, ErrEmptyFirstName
	}
	
	if lastName == "" {
		return nil, ErrEmptyLastName
	}
	
	// Validate phone if provided
	if phone != "" && !isValidPhone(phone) {
		return nil, ErrInvalidPhone
	}
	
	now := time.Now()
	
	customer := &Customer{
		ID:                uuid.New().String(),
		Email:             emailObj,
		FirstName:         firstName,
		LastName:          lastName,
		Phone:             phone,
		Status:            StatusActive,
		CreatedAt:         now,
		UpdatedAt:         now,
		ShippingAddresses: make([]ShippingAddress, 0),
	}
	
	return customer, nil
}

// UpdateEmail updates the customer's email address
func (c *Customer) UpdateEmail(email string) error {
	emailObj, err := NewEmail(email)
	if err != nil {
		return err
	}
	
	c.Email = emailObj
	c.UpdatedAt = time.Now()
	
	return nil
}

// UpdatePersonalInfo updates the customer's personal information
func (c *Customer) UpdatePersonalInfo(firstName, lastName, phone string) error {
	if firstName == "" {
		return ErrEmptyFirstName
	}
	
	if lastName == "" {
		return ErrEmptyLastName
	}
	
	// Validate phone if provided
	if phone != "" && !isValidPhone(phone) {
		return ErrInvalidPhone
	}
	
	c.FirstName = firstName
	c.LastName = lastName
	c.Phone = phone
	c.UpdatedAt = time.Now()
	
	return nil
}

// Activate activates the customer account
func (c *Customer) Activate() error {
	if c.Status == StatusActive {
		return ErrCustomerAlreadyActive
	}
	
	c.Status = StatusActive
	c.UpdatedAt = time.Now()
	
	return nil
}

// Deactivate deactivates the customer account
func (c *Customer) Deactivate() error {
	if c.Status == StatusInactive {
		return ErrCustomerAlreadyInactive
	}
	
	// Note: In a real application, you might want to check for pending orders
	// This would require coordination with the Order domain
	
	c.Status = StatusInactive
	c.UpdatedAt = time.Now()
	
	return nil
}

// Suspend suspends the customer account
func (c *Customer) Suspend() error {
	c.Status = StatusSuspended
	c.UpdatedAt = time.Now()
	
	return nil
}

// AddShippingAddress adds a new shipping address to the customer
func (c *Customer) AddShippingAddress(label, firstName, lastName, company, addressLine1, addressLine2, city, state, postalCode, country, phone string, isDefault bool) error {
	// If this is set as default, ensure no other address is default
	if isDefault {
		for i := range c.ShippingAddresses {
			c.ShippingAddresses[i].UnsetAsDefault()
		}
	}
	
	// If this is the first address, make it default regardless of the flag
	if len(c.ShippingAddresses) == 0 {
		isDefault = true
	}
	
	address, err := NewShippingAddress(
		c.ID, label, firstName, lastName, company,
		addressLine1, addressLine2, city, state, postalCode, country, phone, isDefault,
	)
	if err != nil {
		return err
	}
	
	// Generate ID for the address
	address.ID = uuid.New().String()
	
	c.ShippingAddresses = append(c.ShippingAddresses, address)
	c.UpdatedAt = time.Now()
	
	return nil
}

// UpdateShippingAddress updates an existing shipping address
func (c *Customer) UpdateShippingAddress(addressID, label, firstName, lastName, company, addressLine1, addressLine2, city, state, postalCode, country, phone string) error {
	for i, addr := range c.ShippingAddresses {
		if addr.ID == addressID {
			updatedAddress, err := NewShippingAddress(
				c.ID, label, firstName, lastName, company,
				addressLine1, addressLine2, city, state, postalCode, country, phone, addr.IsDefault,
			)
			if err != nil {
				return err
			}
			
			updatedAddress.ID = addressID
			c.ShippingAddresses[i] = updatedAddress
			c.UpdatedAt = time.Now()
			
			return nil
		}
	}
	
	return ErrAddressNotFound
}

// RemoveShippingAddress removes a shipping address from the customer
func (c *Customer) RemoveShippingAddress(addressID string) error {
	for i, addr := range c.ShippingAddresses {
		if addr.ID == addressID {
			// Cannot delete the default address if it's the only one
			if addr.IsDefault && len(c.ShippingAddresses) == 1 {
				return ErrCustomerHasNoAddresses
			}
			
			// Remove the address
			c.ShippingAddresses = append(c.ShippingAddresses[:i], c.ShippingAddresses[i+1:]...)
			
			// If we removed the default address, make the first remaining address default
			if addr.IsDefault && len(c.ShippingAddresses) > 0 {
				c.ShippingAddresses[0].SetAsDefault()
			}
			
			c.UpdatedAt = time.Now()
			return nil
		}
	}
	
	return ErrAddressNotFound
}

// SetDefaultShippingAddress sets a shipping address as the default
func (c *Customer) SetDefaultShippingAddress(addressID string) error {
	var found bool
	
	// First, unset all addresses as default
	for i := range c.ShippingAddresses {
		if c.ShippingAddresses[i].ID == addressID {
			found = true
		}
		c.ShippingAddresses[i].UnsetAsDefault()
	}
	
	if !found {
		return ErrAddressNotFound
	}
	
	// Set the specified address as default
	for i := range c.ShippingAddresses {
		if c.ShippingAddresses[i].ID == addressID {
			c.ShippingAddresses[i].SetAsDefault()
			break
		}
	}
	
	c.UpdatedAt = time.Now()
	return nil
}

// GetDefaultShippingAddress returns the default shipping address
func (c *Customer) GetDefaultShippingAddress() (*ShippingAddress, error) {
	for _, addr := range c.ShippingAddresses {
		if addr.IsDefault {
			return &addr, nil
		}
	}
	
	return nil, ErrAddressNotFound
}

// GetShippingAddress returns a specific shipping address by ID
func (c *Customer) GetShippingAddress(addressID string) (*ShippingAddress, error) {
	for _, addr := range c.ShippingAddresses {
		if addr.ID == addressID {
			return &addr, nil
		}
	}
	
	return nil, ErrAddressNotFound
}

// Query methods

// IsActive checks if the customer account is active
func (c *Customer) IsActive() bool {
	return c.Status == StatusActive
}

// IsInactive checks if the customer account is inactive
func (c *Customer) IsInactive() bool {
	return c.Status == StatusInactive
}

// IsSuspended checks if the customer account is suspended
func (c *Customer) IsSuspended() bool {
	return c.Status == StatusSuspended
}

// GetFullName returns the customer's full name
func (c *Customer) GetFullName() string {
	return c.FirstName + " " + c.LastName
}

// GetEmailAddress returns the customer's email address as a string
func (c *Customer) GetEmailAddress() string {
	return c.Email.Address
}

// HasShippingAddresses checks if the customer has any shipping addresses
func (c *Customer) HasShippingAddresses() bool {
	return len(c.ShippingAddresses) > 0
}

// GetShippingAddressCount returns the number of shipping addresses
func (c *Customer) GetShippingAddressCount() int {
	return len(c.ShippingAddresses)
}

// CanPlaceOrder checks if the customer can place an order
func (c *Customer) CanPlaceOrder() bool {
	return c.IsActive() && c.HasShippingAddresses()
}
