package customer

import (
	"errors"
	"testing"

	domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"
	"github.com/stretchr/testify/assert"
)

// Tests for GetCustomerUseCase

func TestGetCustomerUseCase(t *testing.T) {
	t.Run("successfully get customer by ID", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewGetCustomerUseCase(mockRepo)
		
		testCustomer := createTestCustomerWithAddress()
		query := GetCustomerQuery{ID: testCustomer.ID}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := useCase.Execute(query)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, testCustomer.ID, response.ID)
		assert.Equal(t, testCustomer.Email.Address, response.Email)
		assert.Equal(t, testCustomer.FirstName, response.FirstName)
		assert.Equal(t, testCustomer.LastName, response.LastName)
		assert.Equal(t, testCustomer.Phone, response.Phone)
		assert.Equal(t, string(testCustomer.Status), response.Status)
		assert.Len(t, response.ShippingAddresses, 1)
		assert.NotEmpty(t, response.CreatedAt)
		assert.NotEmpty(t, response.UpdatedAt)
		
		// Check shipping address details
		addr := response.ShippingAddresses[0]
		domainAddr := testCustomer.ShippingAddresses[0]
		assert.Equal(t, domainAddr.ID, addr.ID)
		assert.Equal(t, domainAddr.Label, addr.Label)
		assert.Equal(t, domainAddr.FirstName, addr.FirstName)
		assert.Equal(t, domainAddr.LastName, addr.LastName)
		assert.Equal(t, domainAddr.City, addr.City)
		assert.Equal(t, domainAddr.IsDefault, addr.IsDefault)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("customer not found", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewGetCustomerUseCase(mockRepo)
		
		query := GetCustomerQuery{ID: "non-existent-id"}
		
		mockRepo.On("FindByID", "non-existent-id").Return(nil, nil)
		
		response, err := useCase.Execute(query)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrCustomerNotFound, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("handle repository error", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewGetCustomerUseCase(mockRepo)
		
		query := GetCustomerQuery{ID: "test-id"}
		repoError := errors.New("database connection error")
		
		mockRepo.On("FindByID", "test-id").Return(nil, repoError)
		
		response, err := useCase.Execute(query)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, repoError, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("get customer with no shipping addresses", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewGetCustomerUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		query := GetCustomerQuery{ID: testCustomer.ID}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := useCase.Execute(query)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Empty(t, response.ShippingAddresses)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("get customer with multiple shipping addresses", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewGetCustomerUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		// Add multiple addresses
		testCustomer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", true,
		)
		testCustomer.AddShippingAddress(
			"Work", "John", "Doe", "Company Inc",
			"456 Work Ave", "Suite 100", "Boston", "MA", "02101", "US", "+1987654321", false,
		)
		
		query := GetCustomerQuery{ID: testCustomer.ID}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := useCase.Execute(query)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Len(t, response.ShippingAddresses, 2)
		
		// Check that one address is default
		defaultCount := 0
		for _, addr := range response.ShippingAddresses {
			if addr.IsDefault {
				defaultCount++
			}
		}
		assert.Equal(t, 1, defaultCount)
		
		mockRepo.AssertExpectations(t)
	})
}
