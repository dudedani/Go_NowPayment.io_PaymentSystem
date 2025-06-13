package customer

import (
	"errors"
	"testing"

	domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCustomerRepository is a mock implementation of CustomerRepository
type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) Save(customer *domainCustomer.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) FindByID(id string) (*domainCustomer.Customer, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainCustomer.Customer), args.Error(1)
}

func (m *MockCustomerRepository) FindByEmail(email string) (*domainCustomer.Customer, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainCustomer.Customer), args.Error(1)
}

func (m *MockCustomerRepository) Update(customer *domainCustomer.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCustomerRepository) ExistsByEmail(email string) (bool, error) {
	args := m.Called(email)
	return args.Bool(0), args.Error(1)
}

// Test helper functions

func createTestCustomerDomain() *domainCustomer.Customer {
	customer, _ := domainCustomer.NewCustomer("test@example.com", "John", "Doe", "+1234567890")
	return customer
}

func createTestCustomerWithAddress() *domainCustomer.Customer {
	customer := createTestCustomerDomain()
	customer.AddShippingAddress(
		"Home", "John", "Doe", "",
		"123 Main St", "", "New York", "NY", "10001", "US", "", false,
	)
	return customer
}

// Tests for RegisterCustomerUseCase

func TestRegisterCustomerUseCase(t *testing.T) {
	t.Run("successfully register new customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewRegisterCustomerUseCase(mockRepo)
		
		cmd := RegisterCustomerCommand{
			Email:     "newuser@example.com",
			FirstName: "Jane",
			LastName:  "Smith",
			Phone:     "+1987654321",
		}
		
		mockRepo.On("ExistsByEmail", cmd.Email).Return(false, nil)
		mockRepo.On("Save", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotEmpty(t, response.ID)
		assert.Equal(t, cmd.Email, response.Email)
		assert.Equal(t, cmd.FirstName, response.FirstName)
		assert.Equal(t, cmd.LastName, response.LastName)
		assert.Equal(t, cmd.Phone, response.Phone)
		assert.Equal(t, "ACTIVE", response.Status)
		assert.NotEmpty(t, response.CreatedAt)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("cannot register customer with existing email", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewRegisterCustomerUseCase(mockRepo)
		
		cmd := RegisterCustomerCommand{
			Email:     "existing@example.com",
			FirstName: "Jane",
			LastName:  "Smith",
		}
		
		mockRepo.On("ExistsByEmail", cmd.Email).Return(true, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrDuplicateEmail, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("cannot register customer with invalid email", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewRegisterCustomerUseCase(mockRepo)
		
		cmd := RegisterCustomerCommand{
			Email:     "invalid-email",
			FirstName: "Jane",
			LastName:  "Smith",
		}
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrInvalidEmail, err)
		
		mockRepo.AssertNotCalled(t, "ExistsByEmail")
		mockRepo.AssertNotCalled(t, "Save")
	})
	
	t.Run("cannot register customer with empty first name", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewRegisterCustomerUseCase(mockRepo)
		
		cmd := RegisterCustomerCommand{
			Email:     "test@example.com",
			FirstName: "",
			LastName:  "Smith",
		}
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrEmptyFirstName, err)
	})
	
	t.Run("cannot register customer with invalid phone", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewRegisterCustomerUseCase(mockRepo)
		
		cmd := RegisterCustomerCommand{
			Email:     "test@example.com",
			FirstName: "Jane",
			LastName:  "Smith",
			Phone:     "invalid-phone",
		}
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrInvalidPhone, err)
	})
	
	t.Run("handle repository error when checking email existence", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewRegisterCustomerUseCase(mockRepo)
		
		cmd := RegisterCustomerCommand{
			Email:     "test@example.com",
			FirstName: "Jane",
			LastName:  "Smith",
		}
		
		repoError := errors.New("database connection error")
		mockRepo.On("ExistsByEmail", cmd.Email).Return(false, repoError)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, repoError, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("handle repository error when saving customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewRegisterCustomerUseCase(mockRepo)
		
		cmd := RegisterCustomerCommand{
			Email:     "test@example.com",
			FirstName: "Jane",
			LastName:  "Smith",
		}
		
		saveError := errors.New("failed to save customer")
		mockRepo.On("ExistsByEmail", cmd.Email).Return(false, nil)
		mockRepo.On("Save", mock.AnythingOfType("*customer.Customer")).Return(saveError)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, saveError, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("register customer without phone", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewRegisterCustomerUseCase(mockRepo)
		
		cmd := RegisterCustomerCommand{
			Email:     "nophone@example.com",
			FirstName: "Jane",
			LastName:  "Smith",
			Phone:     "", // No phone provided
		}
		
		mockRepo.On("ExistsByEmail", cmd.Email).Return(false, nil)
		mockRepo.On("Save", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, "", response.Phone)
		
		mockRepo.AssertExpectations(t)
	})
}
