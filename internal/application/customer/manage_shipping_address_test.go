package customer

import (
	"testing"

	domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Tests for ManageShippingAddressUseCases

func TestUpdateShippingAddressUseCase(t *testing.T) {
	t.Run("successfully update shipping address", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewUpdateShippingAddressUseCase(mockRepo)
		
		testCustomer := createTestCustomerWithAddress()
		addressID := testCustomer.ShippingAddresses[0].ID
		
		cmd := UpdateShippingAddressCommand{
			CustomerID:   testCustomer.ID,
			AddressID:    addressID,
			Label:        "Work",
			FirstName:    "Jane",
			LastName:     "Smith",
			Company:      "New Corp",
			AddressLine1: "789 New St",
			AddressLine2: "Suite 10",
			City:         "Chicago",
			State:        "IL",
			PostalCode:   "60601",
			Country:      "US",
			Phone:        "+1987654321",
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, addressID, response.ID)
		assert.Equal(t, testCustomer.ID, response.CustomerID)
		assert.Equal(t, cmd.Label, response.Label)
		assert.Equal(t, cmd.FirstName, response.FirstName)
		assert.Equal(t, cmd.LastName, response.LastName)
		assert.Equal(t, cmd.Company, response.Company)
		assert.Equal(t, cmd.AddressLine1, response.AddressLine1)
		assert.Equal(t, cmd.AddressLine2, response.AddressLine2)
		assert.Equal(t, cmd.City, response.City)
		assert.Equal(t, cmd.State, response.State)
		assert.Equal(t, cmd.PostalCode, response.PostalCode)
		assert.Equal(t, cmd.Country, response.Country)
		assert.Equal(t, cmd.Phone, response.Phone)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("customer not found", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewUpdateShippingAddressUseCase(mockRepo)
		
		cmd := UpdateShippingAddressCommand{
			CustomerID:   "non-existent-id",
			AddressID:    "address-id",
			Label:        "Work",
			FirstName:    "Jane",
			LastName:     "Smith",
			AddressLine1: "789 New St",
			City:         "Chicago",
			State:        "IL",
			PostalCode:   "60601",
			Country:      "US",
		}
		
		mockRepo.On("FindByID", "non-existent-id").Return(nil, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrCustomerNotFound, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("address not found", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewUpdateShippingAddressUseCase(mockRepo)
		
		testCustomer := createTestCustomerWithAddress()
		
		cmd := UpdateShippingAddressCommand{
			CustomerID:   testCustomer.ID,
			AddressID:    "non-existent-address-id",
			Label:        "Work",
			FirstName:    "Jane",
			LastName:     "Smith",
			AddressLine1: "789 New St",
			City:         "Chicago",
			State:        "IL",
			PostalCode:   "60601",
			Country:      "US",
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrAddressNotFound, err)
		
		mockRepo.AssertNotCalled(t, "Update")
	})
}

func TestRemoveShippingAddressUseCase(t *testing.T) {
	t.Run("successfully remove shipping address", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewRemoveShippingAddressUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		// Add two addresses so we can remove one
		testCustomer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		testCustomer.AddShippingAddress(
			"Work", "John", "Doe", "",
			"456 Work Ave", "", "Boston", "MA", "02101", "US", "", false,
		)
		
		addressToRemove := testCustomer.ShippingAddresses[1].ID
		
		cmd := RemoveShippingAddressCommand{
			CustomerID: testCustomer.ID,
			AddressID:  addressToRemove,
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)
		assert.Equal(t, "Shipping address removed successfully", response.Message)
		assert.Equal(t, addressToRemove, response.AddressID)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("cannot remove only shipping address", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewRemoveShippingAddressUseCase(mockRepo)
		
		testCustomer := createTestCustomerWithAddress() // Has only one address
		addressID := testCustomer.ShippingAddresses[0].ID
		
		cmd := RemoveShippingAddressCommand{
			CustomerID: testCustomer.ID,
			AddressID:  addressID,
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrCustomerHasNoAddresses, err)
		
		mockRepo.AssertNotCalled(t, "Update")
	})
}

func TestSetDefaultShippingAddressUseCase(t *testing.T) {
	t.Run("successfully set default shipping address", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewSetDefaultShippingAddressUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		// Add two addresses
		testCustomer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", true,
		)
		testCustomer.AddShippingAddress(
			"Work", "John", "Doe", "",
			"456 Work Ave", "", "Boston", "MA", "02101", "US", "", false,
		)
		
		workAddressID := testCustomer.ShippingAddresses[1].ID
		
		cmd := SetDefaultShippingAddressCommand{
			CustomerID: testCustomer.ID,
			AddressID:  workAddressID,
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.True(t, response.Success)
		assert.Equal(t, "Default shipping address updated successfully", response.Message)
		assert.Equal(t, workAddressID, response.AddressID)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("customer not found", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewSetDefaultShippingAddressUseCase(mockRepo)
		
		cmd := SetDefaultShippingAddressCommand{
			CustomerID: "non-existent-id",
			AddressID:  "address-id",
		}
		
		mockRepo.On("FindByID", "non-existent-id").Return(nil, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrCustomerNotFound, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("address not found", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewSetDefaultShippingAddressUseCase(mockRepo)
		
		testCustomer := createTestCustomerWithAddress()
		
		cmd := SetDefaultShippingAddressCommand{
			CustomerID: testCustomer.ID,
			AddressID:  "non-existent-address-id",
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrAddressNotFound, err)
		
		mockRepo.AssertNotCalled(t, "Update")
	})
}
