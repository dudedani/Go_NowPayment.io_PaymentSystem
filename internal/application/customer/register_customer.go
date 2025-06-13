package customer

import (
	domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"
)

// RegisterCustomerCommand represents the input for registering a new customer
type RegisterCustomerCommand struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone,omitempty"`
}

// RegisterCustomerResponse represents the output after registering a customer
type RegisterCustomerResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone,omitempty"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// CustomerRepository defines the interface for customer persistence
type CustomerRepository interface {
	Save(customer *domainCustomer.Customer) error
	FindByID(id string) (*domainCustomer.Customer, error)
	FindByEmail(email string) (*domainCustomer.Customer, error)
	Update(customer *domainCustomer.Customer) error
	Delete(id string) error
	ExistsByEmail(email string) (bool, error)
}

// RegisterCustomerUseCase handles customer registration
type RegisterCustomerUseCase struct {
	customerRepo CustomerRepository
}

// NewRegisterCustomerUseCase creates a new instance of RegisterCustomerUseCase
func NewRegisterCustomerUseCase(customerRepo CustomerRepository) *RegisterCustomerUseCase {
	return &RegisterCustomerUseCase{
		customerRepo: customerRepo,
	}
}

// Execute registers a new customer
func (uc *RegisterCustomerUseCase) Execute(cmd RegisterCustomerCommand) (*RegisterCustomerResponse, error) {
	// Create new customer (this will validate email format)
	newCustomer, err := domainCustomer.NewCustomer(cmd.Email, cmd.FirstName, cmd.LastName, cmd.Phone)
	if err != nil {
		return nil, err
	}
	
	// Check if email already exists
	exists, err := uc.customerRepo.ExistsByEmail(cmd.Email)
	if err != nil {
		return nil, err
	}
	
	if exists {
		return nil, domainCustomer.ErrDuplicateEmail
	}
	
	// Save customer
	if err := uc.customerRepo.Save(newCustomer); err != nil {
		return nil, err
	}
	
	// Return response
	return &RegisterCustomerResponse{
		ID:        newCustomer.ID,
		Email:     newCustomer.Email.Address,
		FirstName: newCustomer.FirstName,
		LastName:  newCustomer.LastName,
		Phone:     newCustomer.Phone,
		Status:    string(newCustomer.Status),
		CreatedAt: newCustomer.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
