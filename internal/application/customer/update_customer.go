package customer

import domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"

// UpdateCustomerCommand represents the input for updating customer details
type UpdateCustomerCommand struct {
	ID        string `json:"id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone,omitempty"`
}

// UpdateCustomerResponse represents the output after updating a customer
type UpdateCustomerResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone,omitempty"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}

// UpdateCustomerUseCase handles customer information updates
type UpdateCustomerUseCase struct {
	customerRepo CustomerRepository
}

// NewUpdateCustomerUseCase creates a new instance of UpdateCustomerUseCase
func NewUpdateCustomerUseCase(customerRepo CustomerRepository) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{
		customerRepo: customerRepo,
	}
}

// Execute updates customer information
func (uc *UpdateCustomerUseCase) Execute(cmd UpdateCustomerCommand) (*UpdateCustomerResponse, error) {
	// Find customer by ID
	existingCustomer, err := uc.customerRepo.FindByID(cmd.ID)
	if err != nil {
		return nil, err
	}
	
	if existingCustomer == nil {
		return nil, domainCustomer.ErrCustomerNotFound
	}
	
	// Update personal information
	if err := existingCustomer.UpdatePersonalInfo(cmd.FirstName, cmd.LastName, cmd.Phone); err != nil {
		return nil, err
	}
	
	// Save updated customer
	if err := uc.customerRepo.Update(existingCustomer); err != nil {
		return nil, err
	}
	
	// Return response
	return &UpdateCustomerResponse{
		ID:        existingCustomer.ID,
		Email:     existingCustomer.Email.Address,
		FirstName: existingCustomer.FirstName,
		LastName:  existingCustomer.LastName,
		Phone:     existingCustomer.Phone,
		Status:    string(existingCustomer.Status),
		UpdatedAt: existingCustomer.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
