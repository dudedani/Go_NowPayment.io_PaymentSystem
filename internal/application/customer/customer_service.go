package customer

import domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"

// CustomerService provides high-level customer operations
type CustomerService struct {
	customerRepo CustomerRepository
	
	// Use cases
	registerCustomer           *RegisterCustomerUseCase
	getCustomer               *GetCustomerUseCase
	updateCustomer            *UpdateCustomerUseCase
	addShippingAddress        *AddShippingAddressUseCase
	updateShippingAddress     *UpdateShippingAddressUseCase
	removeShippingAddress     *RemoveShippingAddressUseCase
	setDefaultShippingAddress *SetDefaultShippingAddressUseCase
}

// NewCustomerService creates a new instance of CustomerService
func NewCustomerService(customerRepo CustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepo:              customerRepo,
		registerCustomer:          NewRegisterCustomerUseCase(customerRepo),
		getCustomer:              NewGetCustomerUseCase(customerRepo),
		updateCustomer:           NewUpdateCustomerUseCase(customerRepo),
		addShippingAddress:       NewAddShippingAddressUseCase(customerRepo),
		updateShippingAddress:    NewUpdateShippingAddressUseCase(customerRepo),
		removeShippingAddress:    NewRemoveShippingAddressUseCase(customerRepo),
		setDefaultShippingAddress: NewSetDefaultShippingAddressUseCase(customerRepo),
	}
}

// RegisterCustomer registers a new customer
func (s *CustomerService) RegisterCustomer(cmd RegisterCustomerCommand) (*RegisterCustomerResponse, error) {
	return s.registerCustomer.Execute(cmd)
}

// GetCustomer retrieves customer details by ID
func (s *CustomerService) GetCustomer(query GetCustomerQuery) (*GetCustomerResponse, error) {
	return s.getCustomer.Execute(query)
}

// UpdateCustomer updates customer information
func (s *CustomerService) UpdateCustomer(cmd UpdateCustomerCommand) (*UpdateCustomerResponse, error) {
	return s.updateCustomer.Execute(cmd)
}

// AddShippingAddress adds a shipping address to a customer
func (s *CustomerService) AddShippingAddress(cmd AddShippingAddressCommand) (*AddShippingAddressResponse, error) {
	return s.addShippingAddress.Execute(cmd)
}

// UpdateShippingAddress updates a shipping address for a customer
func (s *CustomerService) UpdateShippingAddress(cmd UpdateShippingAddressCommand) (*UpdateShippingAddressResponse, error) {
	return s.updateShippingAddress.Execute(cmd)
}

// RemoveShippingAddress removes a shipping address from a customer
func (s *CustomerService) RemoveShippingAddress(cmd RemoveShippingAddressCommand) (*RemoveShippingAddressResponse, error) {
	return s.removeShippingAddress.Execute(cmd)
}

// SetDefaultShippingAddress sets a shipping address as the default for a customer
func (s *CustomerService) SetDefaultShippingAddress(cmd SetDefaultShippingAddressCommand) (*SetDefaultShippingAddressResponse, error) {
	return s.setDefaultShippingAddress.Execute(cmd)
}

// GetCustomerByEmail retrieves customer by email address
func (s *CustomerService) GetCustomerByEmail(email string) (*GetCustomerResponse, error) {
	// Find customer by email
	customer, err := s.customerRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	
	if customer == nil {
		return nil, domainCustomer.ErrCustomerNotFound
	}
	
	// Convert to response using GetCustomerUseCase logic
	return s.getCustomer.Execute(GetCustomerQuery{ID: customer.ID})
}

// DeactivateCustomer deactivates a customer account
func (s *CustomerService) DeactivateCustomer(customerID string) error {
	// Find customer by ID
	customer, err := s.customerRepo.FindByID(customerID)
	if err != nil {
		return err
	}
	
	if customer == nil {
		return domainCustomer.ErrCustomerNotFound
	}
	
	// Deactivate customer
	if err := customer.Deactivate(); err != nil {
		return err
	}
	
	// Save updated customer
	return s.customerRepo.Update(customer)
}

// ActivateCustomer activates a customer account
func (s *CustomerService) ActivateCustomer(customerID string) error {
	// Find customer by ID
	customer, err := s.customerRepo.FindByID(customerID)
	if err != nil {
		return err
	}
	
	if customer == nil {
		return domainCustomer.ErrCustomerNotFound
	}
	
	// Activate customer
	if err := customer.Activate(); err != nil {
		return err
	}
	
	// Save updated customer
	return s.customerRepo.Update(customer)
}

// SuspendCustomer suspends a customer account
func (s *CustomerService) SuspendCustomer(customerID string) error {
	// Find customer by ID
	customer, err := s.customerRepo.FindByID(customerID)
	if err != nil {
		return err
	}
	
	if customer == nil {
		return domainCustomer.ErrCustomerNotFound
	}
	
	// Suspend customer
	if err := customer.Suspend(); err != nil {
		return err
	}
	
	// Save updated customer
	return s.customerRepo.Update(customer)
}

// CanCustomerPlaceOrder checks if a customer can place an order
func (s *CustomerService) CanCustomerPlaceOrder(customerID string) (bool, error) {
	// Find customer by ID
	customer, err := s.customerRepo.FindByID(customerID)
	if err != nil {
		return false, err
	}
	
	if customer == nil {
		return false, domainCustomer.ErrCustomerNotFound
	}
	
	return customer.CanPlaceOrder(), nil
}
