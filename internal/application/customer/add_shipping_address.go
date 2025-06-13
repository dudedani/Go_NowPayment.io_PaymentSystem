package customer

import (
	domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"
)

// AddShippingAddressCommand represents the input for adding a shipping address
type AddShippingAddressCommand struct {
	CustomerID   string `json:"customer_id" validate:"required"`
	Label        string `json:"label" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Company      string `json:"company,omitempty"`
	AddressLine1 string `json:"address_line1" validate:"required"`
	AddressLine2 string `json:"address_line2,omitempty"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required"`
	PostalCode   string `json:"postal_code" validate:"required"`
	Country      string `json:"country" validate:"required,len=2"`
	Phone        string `json:"phone,omitempty"`
	IsDefault    bool   `json:"is_default"`
}

// AddShippingAddressResponse represents the output after adding a shipping address
type AddShippingAddressResponse struct {
	ID           string `json:"id"`
	CustomerID   string `json:"customer_id"`
	Label        string `json:"label"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Company      string `json:"company,omitempty"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2,omitempty"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Phone        string `json:"phone,omitempty"`
	IsDefault    bool   `json:"is_default"`
}

// AddShippingAddressUseCase handles adding shipping addresses to customers
type AddShippingAddressUseCase struct {
	customerRepo CustomerRepository
}

// NewAddShippingAddressUseCase creates a new instance of AddShippingAddressUseCase
func NewAddShippingAddressUseCase(customerRepo CustomerRepository) *AddShippingAddressUseCase {
	return &AddShippingAddressUseCase{
		customerRepo: customerRepo,
	}
}

// Execute adds a shipping address to a customer
func (uc *AddShippingAddressUseCase) Execute(cmd AddShippingAddressCommand) (*AddShippingAddressResponse, error) {
	// Find customer by ID
	customer, err := uc.customerRepo.FindByID(cmd.CustomerID)
	if err != nil {
		return nil, err
	}
	
	if customer == nil {
		return nil, domainCustomer.ErrCustomerNotFound
	}
	
	// Add shipping address
	if err := customer.AddShippingAddress(
		cmd.Label,
		cmd.FirstName,
		cmd.LastName,
		cmd.Company,
		cmd.AddressLine1,
		cmd.AddressLine2,
		cmd.City,
		cmd.State,
		cmd.PostalCode,
		cmd.Country,
		cmd.Phone,
		cmd.IsDefault,
	); err != nil {
		return nil, err
	}
	
	// Save updated customer
	if err := uc.customerRepo.Update(customer); err != nil {
		return nil, err
	}
	
	// Get the newly added address (it's the last one in the slice)
	newAddress := customer.ShippingAddresses[len(customer.ShippingAddresses)-1]
	
	// Return response
	return &AddShippingAddressResponse{
		ID:           newAddress.ID,
		CustomerID:   newAddress.CustomerID,
		Label:        newAddress.Label,
		FirstName:    newAddress.FirstName,
		LastName:     newAddress.LastName,
		Company:      newAddress.Company,
		AddressLine1: newAddress.AddressLine1,
		AddressLine2: newAddress.AddressLine2,
		City:         newAddress.City,
		State:        newAddress.State,
		PostalCode:   newAddress.PostalCode,
		Country:      newAddress.Country,
		Phone:        newAddress.Phone,
		IsDefault:    newAddress.IsDefault,
	}, nil
}
