package customer

import domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"

// RemoveShippingAddressCommand represents the input for removing a shipping address
type RemoveShippingAddressCommand struct {
	CustomerID string `json:"customer_id" validate:"required"`
	AddressID  string `json:"address_id" validate:"required"`
}

// RemoveShippingAddressResponse represents the output after removing a shipping address
type RemoveShippingAddressResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	AddressID string `json:"address_id"`
}

// RemoveShippingAddressUseCase handles removing shipping addresses
type RemoveShippingAddressUseCase struct {
	customerRepo CustomerRepository
}

// NewRemoveShippingAddressUseCase creates a new instance of RemoveShippingAddressUseCase
func NewRemoveShippingAddressUseCase(customerRepo CustomerRepository) *RemoveShippingAddressUseCase {
	return &RemoveShippingAddressUseCase{
		customerRepo: customerRepo,
	}
}

// Execute removes a shipping address from a customer
func (uc *RemoveShippingAddressUseCase) Execute(cmd RemoveShippingAddressCommand) (*RemoveShippingAddressResponse, error) {
	// Find customer by ID
	customer, err := uc.customerRepo.FindByID(cmd.CustomerID)
	if err != nil {
		return nil, err
	}
	
	if customer == nil {
		return nil, domainCustomer.ErrCustomerNotFound
	}
	
	// Remove shipping address
	if err := customer.RemoveShippingAddress(cmd.AddressID); err != nil {
		return nil, err
	}
	
	// Save updated customer
	if err := uc.customerRepo.Update(customer); err != nil {
		return nil, err
	}
	
	// Return response
	return &RemoveShippingAddressResponse{
		Success:   true,
		Message:   "Shipping address removed successfully",
		AddressID: cmd.AddressID,
	}, nil
}

// SetDefaultShippingAddressCommand represents the input for setting a default shipping address
type SetDefaultShippingAddressCommand struct {
	CustomerID string `json:"customer_id" validate:"required"`
	AddressID  string `json:"address_id" validate:"required"`
}

// SetDefaultShippingAddressResponse represents the output after setting a default shipping address
type SetDefaultShippingAddressResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	AddressID string `json:"address_id"`
}

// SetDefaultShippingAddressUseCase handles setting default shipping addresses
type SetDefaultShippingAddressUseCase struct {
	customerRepo CustomerRepository
}

// NewSetDefaultShippingAddressUseCase creates a new instance of SetDefaultShippingAddressUseCase
func NewSetDefaultShippingAddressUseCase(customerRepo CustomerRepository) *SetDefaultShippingAddressUseCase {
	return &SetDefaultShippingAddressUseCase{
		customerRepo: customerRepo,
	}
}

// Execute sets a shipping address as the default for a customer
func (uc *SetDefaultShippingAddressUseCase) Execute(cmd SetDefaultShippingAddressCommand) (*SetDefaultShippingAddressResponse, error) {
	// Find customer by ID
	customer, err := uc.customerRepo.FindByID(cmd.CustomerID)
	if err != nil {
		return nil, err
	}
	
	if customer == nil {
		return nil, domainCustomer.ErrCustomerNotFound
	}
	
	// Set default shipping address
	if err := customer.SetDefaultShippingAddress(cmd.AddressID); err != nil {
		return nil, err
	}
	
	// Save updated customer
	if err := uc.customerRepo.Update(customer); err != nil {
		return nil, err
	}
	
	// Return response
	return &SetDefaultShippingAddressResponse{
		Success:   true,
		Message:   "Default shipping address updated successfully",
		AddressID: cmd.AddressID,
	}, nil
}
