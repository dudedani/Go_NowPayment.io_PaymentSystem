package customer

import (
	"errors"
	"testing"

	domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Tests for CustomerService

func TestCustomerService(t *testing.T) {
	t.Run("register customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		cmd := RegisterCustomerCommand{
			Email:     "test@example.com",
			FirstName: "John",
			LastName:  "Doe",
			Phone:     "+1234567890",
		}
		
		mockRepo.On("ExistsByEmail", cmd.Email).Return(false, nil)
		mockRepo.On("Save", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := service.RegisterCustomer(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, cmd.Email, response.Email)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("get customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		query := GetCustomerQuery{ID: testCustomer.ID}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := service.GetCustomer(query)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, testCustomer.ID, response.ID)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("get customer by email", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		email := testCustomer.Email.Address
		
		mockRepo.On("FindByEmail", email).Return(testCustomer, nil)
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := service.GetCustomerByEmail(email)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, testCustomer.ID, response.ID)
		assert.Equal(t, email, response.Email)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("get customer by email - not found", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		email := "notfound@example.com"
		
		mockRepo.On("FindByEmail", email).Return(nil, nil)
		
		response, err := service.GetCustomerByEmail(email)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrCustomerNotFound, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("update customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		cmd := UpdateCustomerCommand{
			ID:        testCustomer.ID,
			FirstName: "Jane",
			LastName:  "Smith",
			Phone:     "+1987654321",
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := service.UpdateCustomer(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, cmd.FirstName, response.FirstName)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("deactivate customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		err := service.DeactivateCustomer(testCustomer.ID)
		
		assert.NoError(t, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("activate customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		testCustomer.Status = domainCustomer.StatusInactive
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		err := service.ActivateCustomer(testCustomer.ID)
		
		assert.NoError(t, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("suspend customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		err := service.SuspendCustomer(testCustomer.ID)
		
		assert.NoError(t, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("can customer place order - yes", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		testCustomer := createTestCustomerWithAddress() // Active customer with address
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		canPlace, err := service.CanCustomerPlaceOrder(testCustomer.ID)
		
		assert.NoError(t, err)
		assert.True(t, canPlace)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("can customer place order - no addresses", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		testCustomer := createTestCustomerDomain() // Active customer without addresses
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		canPlace, err := service.CanCustomerPlaceOrder(testCustomer.ID)
		
		assert.NoError(t, err)
		assert.False(t, canPlace)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("can customer place order - inactive", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		testCustomer := createTestCustomerWithAddress()
		testCustomer.Status = domainCustomer.StatusInactive // Inactive customer
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		canPlace, err := service.CanCustomerPlaceOrder(testCustomer.ID)
		
		assert.NoError(t, err)
		assert.False(t, canPlace)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("can customer place order - customer not found", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		mockRepo.On("FindByID", "non-existent-id").Return(nil, nil)
		
		canPlace, err := service.CanCustomerPlaceOrder("non-existent-id")
		
		assert.Error(t, err)
		assert.False(t, canPlace)
		assert.Equal(t, domainCustomer.ErrCustomerNotFound, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("add shipping address", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		cmd := AddShippingAddressCommand{
			CustomerID:   testCustomer.ID,
			Label:        "Home",
			FirstName:    "John",
			LastName:     "Doe",
			AddressLine1: "123 Main St",
			City:         "New York",
			State:        "NY",
			PostalCode:   "10001",
			Country:      "US",
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := service.AddShippingAddress(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, cmd.Label, response.Label)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("handle repository errors", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		service := NewCustomerService(mockRepo)
		
		repoError := errors.New("database connection error")
		
		// Test deactivate customer with repository error
		mockRepo.On("FindByID", "test-id").Return(nil, repoError)
		
		err := service.DeactivateCustomer("test-id")
		
		assert.Error(t, err)
		assert.Equal(t, repoError, err)
		
		mockRepo.AssertExpectations(t)
	})
}
