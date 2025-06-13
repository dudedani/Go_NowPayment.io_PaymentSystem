package customer

import domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"

// UpdateShippingAddressCommand represents the input for updating a shipping address
type UpdateShippingAddressCommand struct {
	CustomerID   string `json:"customer_id" validate:"required"`
	AddressID    string `json:"address_id" validate:"required"`
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
}

// UpdateShippingAddressResponse represents the output after updating a shipping address
type UpdateShippingAddressResponse struct {
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

// UpdateShippingAddressUseCase handles updating shipping addresses
type UpdateShippingAddressUseCase struct {
	customerRepo CustomerRepository
}

// NewUpdateShippingAddressUseCase creates a new instance of UpdateShippingAddressUseCase
func NewUpdateShippingAddressUseCase(customerRepo CustomerRepository) *UpdateShippingAddressUseCase {
	return &UpdateShippingAddressUseCase{
		customerRepo: customerRepo,
	}
}

// Execute updates a shipping address for a customer
func (uc *UpdateShippingAddressUseCase) Execute(cmd UpdateShippingAddressCommand) (*UpdateShippingAddressResponse, error) {
	// Find customer by ID
	customer, err := uc.customerRepo.FindByID(cmd.CustomerID)
	if err != nil {
		return nil, err
	}
	
	if customer == nil {
		return nil, domainCustomer.ErrCustomerNotFound
	}
	
	// Update shipping address
	if err := customer.UpdateShippingAddress(
		cmd.AddressID,
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
	); err != nil {
		return nil, err
	}
	
	// Save updated customer
	if err := uc.customerRepo.Update(customer); err != nil {
		return nil, err
	}
	
	// Get the updated address
	updatedAddress, err := customer.GetShippingAddress(cmd.AddressID)
	if err != nil {
		return nil, err
	}
	
	// Return response
	return &UpdateShippingAddressResponse{
		ID:           updatedAddress.ID,
		CustomerID:   updatedAddress.CustomerID,
		Label:        updatedAddress.Label,
		FirstName:    updatedAddress.FirstName,
		LastName:     updatedAddress.LastName,
		Company:      updatedAddress.Company,
		AddressLine1: updatedAddress.AddressLine1,
		AddressLine2: updatedAddress.AddressLine2,
		City:         updatedAddress.City,
		State:        updatedAddress.State,
		PostalCode:   updatedAddress.PostalCode,
		Country:      updatedAddress.Country,
		Phone:        updatedAddress.Phone,
		IsDefault:    updatedAddress.IsDefault,
	}, nil
}
