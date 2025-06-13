package customer

import (
	"errors"
	"testing"

	domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Tests for AddShippingAddressUseCase

func TestAddShippingAddressUseCase(t *testing.T) {
	t.Run("successfully add shipping address", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewAddShippingAddressUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		cmd := AddShippingAddressCommand{
			CustomerID:   testCustomer.ID,
			Label:        "Home",
			FirstName:    "John",
			LastName:     "Doe",
			Company:      "",
			AddressLine1: "123 Main St",
			AddressLine2: "",
			City:         "New York",
			State:        "NY",
			PostalCode:   "10001",
			Country:      "US",
			Phone:        "",
			IsDefault:    true,
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.NotEmpty(t, response.ID)
		assert.Equal(t, testCustomer.ID, response.CustomerID)
		assert.Equal(t, cmd.Label, response.Label)
		assert.Equal(t, cmd.FirstName, response.FirstName)
		assert.Equal(t, cmd.LastName, response.LastName)
		assert.Equal(t, cmd.AddressLine1, response.AddressLine1)
		assert.Equal(t, cmd.City, response.City)
		assert.Equal(t, cmd.State, response.State)
		assert.Equal(t, cmd.PostalCode, response.PostalCode)
		assert.Equal(t, cmd.Country, response.Country)
		assert.Equal(t, cmd.IsDefault, response.IsDefault)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("customer not found", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewAddShippingAddressUseCase(mockRepo)
		
		cmd := AddShippingAddressCommand{
			CustomerID:   "non-existent-id",
			Label:        "Home",
			FirstName:    "John",
			LastName:     "Doe",
			AddressLine1: "123 Main St",
			City:         "New York",
			State:        "NY",
			PostalCode:   "10001",
			Country:      "US",
		}
		
		mockRepo.On("FindByID", "non-existent-id").Return(nil, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrCustomerNotFound, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("cannot add address with empty address line 1", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewAddShippingAddressUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		cmd := AddShippingAddressCommand{
			CustomerID:   testCustomer.ID,
			Label:        "Home",
			FirstName:    "John",
			LastName:     "Doe",
			AddressLine1: "", // Empty address line 1
			City:         "New York",
			State:        "NY",
			PostalCode:   "10001",
			Country:      "US",
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrEmptyAddressLine1, err)
		
		mockRepo.AssertNotCalled(t, "Update")
	})
	
	t.Run("cannot add address with invalid country code", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewAddShippingAddressUseCase(mockRepo)
		
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
			Country:      "USA", // Invalid - should be 2 characters
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrInvalidCountryCode, err)
		
		mockRepo.AssertNotCalled(t, "Update")
	})
	
	t.Run("add address with all optional fields", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewAddShippingAddressUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		cmd := AddShippingAddressCommand{
			CustomerID:   testCustomer.ID,
			Label:        "Work",
			FirstName:    "John",
			LastName:     "Doe",
			Company:      "Acme Corp",
			AddressLine1: "456 Work Ave",
			AddressLine2: "Suite 100",
			City:         "Boston",
			State:        "MA",
			PostalCode:   "02101",
			Country:      "US",
			Phone:        "+1987654321",
			IsDefault:    false,
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, cmd.Company, response.Company)
		assert.Equal(t, cmd.AddressLine2, response.AddressLine2)
		assert.Equal(t, cmd.Phone, response.Phone)
		assert.True(t, response.IsDefault) // First address is always default
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("handle repository error when finding customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewAddShippingAddressUseCase(mockRepo)
		
		cmd := AddShippingAddressCommand{
			CustomerID:   "test-id",
			Label:        "Home",
			FirstName:    "John",
			LastName:     "Doe",
			AddressLine1: "123 Main St",
			City:         "New York",
			State:        "NY",
			PostalCode:   "10001",
			Country:      "US",
		}
		
		repoError := errors.New("database connection error")
		mockRepo.On("FindByID", "test-id").Return(nil, repoError)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, repoError, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("handle repository error when updating customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewAddShippingAddressUseCase(mockRepo)
		
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
		
		updateError := errors.New("failed to update customer")
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(updateError)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, updateError, err)
		
		mockRepo.AssertExpectations(t)
	})
}
