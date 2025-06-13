package customer

import (
	"errors"
	"testing"

	domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Tests for UpdateCustomerUseCase

func TestUpdateCustomerUseCase(t *testing.T) {
	t.Run("successfully update customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewUpdateCustomerUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		cmd := UpdateCustomerCommand{
			ID:        testCustomer.ID,
			FirstName: "Jane",
			LastName:  "Smith",
			Phone:     "+1987654321",
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, testCustomer.ID, response.ID)
		assert.Equal(t, testCustomer.Email.Address, response.Email)
		assert.Equal(t, cmd.FirstName, response.FirstName)
		assert.Equal(t, cmd.LastName, response.LastName)
		assert.Equal(t, cmd.Phone, response.Phone)
		assert.Equal(t, string(testCustomer.Status), response.Status)
		assert.NotEmpty(t, response.UpdatedAt)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("customer not found", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewUpdateCustomerUseCase(mockRepo)
		
		cmd := UpdateCustomerCommand{
			ID:        "non-existent-id",
			FirstName: "Jane",
			LastName:  "Smith",
		}
		
		mockRepo.On("FindByID", "non-existent-id").Return(nil, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrCustomerNotFound, err)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("cannot update with empty first name", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewUpdateCustomerUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		cmd := UpdateCustomerCommand{
			ID:        testCustomer.ID,
			FirstName: "",
			LastName:  "Smith",
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrEmptyFirstName, err)
		
		mockRepo.AssertNotCalled(t, "Update")
	})
	
	t.Run("cannot update with empty last name", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewUpdateCustomerUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		cmd := UpdateCustomerCommand{
			ID:        testCustomer.ID,
			FirstName: "Jane",
			LastName:  "",
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrEmptyLastName, err)
		
		mockRepo.AssertNotCalled(t, "Update")
	})
	
	t.Run("cannot update with invalid phone", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewUpdateCustomerUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		cmd := UpdateCustomerCommand{
			ID:        testCustomer.ID,
			FirstName: "Jane",
			LastName:  "Smith",
			Phone:     "invalid-phone",
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, domainCustomer.ErrInvalidPhone, err)
		
		mockRepo.AssertNotCalled(t, "Update")
	})
	
	t.Run("update customer without phone", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewUpdateCustomerUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		cmd := UpdateCustomerCommand{
			ID:        testCustomer.ID,
			FirstName: "Jane",
			LastName:  "Smith",
			Phone:     "", // Remove phone
		}
		
		mockRepo.On("FindByID", testCustomer.ID).Return(testCustomer, nil)
		mockRepo.On("Update", mock.AnythingOfType("*customer.Customer")).Return(nil)
		
		response, err := useCase.Execute(cmd)
		
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, "", response.Phone)
		
		mockRepo.AssertExpectations(t)
	})
	
	t.Run("handle repository error when finding customer", func(t *testing.T) {
		mockRepo := new(MockCustomerRepository)
		useCase := NewUpdateCustomerUseCase(mockRepo)
		
		cmd := UpdateCustomerCommand{
			ID:        "test-id",
			FirstName: "Jane",
			LastName:  "Smith",
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
		useCase := NewUpdateCustomerUseCase(mockRepo)
		
		testCustomer := createTestCustomerDomain()
		cmd := UpdateCustomerCommand{
			ID:        testCustomer.ID,
			FirstName: "Jane",
			LastName:  "Smith",
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
