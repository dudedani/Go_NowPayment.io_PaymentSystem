package customer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test helper functions

// createTestEmail creates a valid email for testing
func createTestEmail() Email {
	email, _ := NewEmail("test@example.com")
	return email
}

// createTestCustomer creates a valid customer for testing
func createTestCustomer() *Customer {
	customer, _ := NewCustomer("test@example.com", "John", "Doe", "+1234567890")
	return customer
}

// createTestShippingAddress creates a valid shipping address for testing
func createTestShippingAddress(customerID string) ShippingAddress {
	addr, _ := NewShippingAddress(
		customerID,
		"Home",
		"John",
		"Doe",
		"Acme Corp",
		"123 Main St",
		"Apt 4B",
		"New York",
		"NY",
		"10001",
		"US",
		"+1234567890",
		true,
	)
	return addr
}

// Tests for Email Value Object

func TestEmail(t *testing.T) {
	t.Run("create valid email", func(t *testing.T) {
		email, err := NewEmail("test@example.com")
		
		assert.NoError(t, err)
		assert.Equal(t, "test@example.com", email.Address)
	})
	
	t.Run("create email with uppercase", func(t *testing.T) {
		email, err := NewEmail("TEST@EXAMPLE.COM")
		
		assert.NoError(t, err)
		assert.Equal(t, "test@example.com", email.Address) // Should be normalized to lowercase
	})
	
	t.Run("create email with whitespace", func(t *testing.T) {
		email, err := NewEmail("  test@example.com  ")
		
		assert.NoError(t, err)
		assert.Equal(t, "test@example.com", email.Address)
	})
	
	t.Run("cannot create email with empty string", func(t *testing.T) {
		_, err := NewEmail("")
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyEmail, err)
	})
	
	t.Run("cannot create email with invalid format", func(t *testing.T) {
		invalidEmails := []string{
			"invalid",
			"@example.com",
			"test@",
			"test.example.com",
			"test@.com",
			"test@com",
		}
		
		for _, invalidEmail := range invalidEmails {
			_, err := NewEmail(invalidEmail)
			assert.Error(t, err)
			assert.Equal(t, ErrInvalidEmail, err)
		}
	})
	
	t.Run("get domain", func(t *testing.T) {
		email := createTestEmail()
		
		assert.Equal(t, "example.com", email.GetDomain())
	})
	
	t.Run("get username", func(t *testing.T) {
		email := createTestEmail()
		
		assert.Equal(t, "test", email.GetUsername())
	})
	
	t.Run("email equality", func(t *testing.T) {
		email1, _ := NewEmail("test@example.com")
		email2, _ := NewEmail("test@example.com")
		email3, _ := NewEmail("other@example.com")
		
		assert.True(t, email1.Equals(email2))
		assert.False(t, email1.Equals(email3))
	})
	
	t.Run("email string representation", func(t *testing.T) {
		email := createTestEmail()
		
		assert.Equal(t, "test@example.com", email.String())
	})
}

// Tests for ShippingAddress Value Object

func TestShippingAddress(t *testing.T) {
	t.Run("create valid shipping address", func(t *testing.T) {
		addr, err := NewShippingAddress(
			"customer-123",
			"Home",
			"John",
			"Doe",
			"Acme Corp",
			"123 Main St",
			"Apt 4B",
			"New York",
			"NY",
			"10001",
			"US",
			"+1234567890",
			true,
		)
		
		assert.NoError(t, err)
		assert.Equal(t, "customer-123", addr.CustomerID)
		assert.Equal(t, "Home", addr.Label)
		assert.Equal(t, "John", addr.FirstName)
		assert.Equal(t, "Doe", addr.LastName)
		assert.Equal(t, "Acme Corp", addr.Company)
		assert.Equal(t, "123 Main St", addr.AddressLine1)
		assert.Equal(t, "Apt 4B", addr.AddressLine2)
		assert.Equal(t, "New York", addr.City)
		assert.Equal(t, "NY", addr.State)
		assert.Equal(t, "10001", addr.PostalCode)
		assert.Equal(t, "US", addr.Country)
		assert.Equal(t, "+1234567890", addr.Phone)
		assert.True(t, addr.IsDefault)
	})
	
	t.Run("cannot create address with empty customer ID", func(t *testing.T) {
		_, err := NewShippingAddress(
			"", "Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyCustomerID, err)
	})
	
	t.Run("cannot create address with empty first name", func(t *testing.T) {
		_, err := NewShippingAddress(
			"customer-123", "Home", "", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyFirstName, err)
	})
	
	t.Run("cannot create address with empty last name", func(t *testing.T) {
		_, err := NewShippingAddress(
			"customer-123", "Home", "John", "", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyLastName, err)
	})
	
	t.Run("cannot create address with empty address line 1", func(t *testing.T) {
		_, err := NewShippingAddress(
			"customer-123", "Home", "John", "Doe", "",
			"", "", "New York", "NY", "10001", "US", "", false,
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyAddressLine1, err)
	})
	
	t.Run("cannot create address with empty city", func(t *testing.T) {
		_, err := NewShippingAddress(
			"customer-123", "Home", "John", "Doe", "",
			"123 Main St", "", "", "NY", "10001", "US", "", false,
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyCity, err)
	})
	
	t.Run("cannot create address with empty state", func(t *testing.T) {
		_, err := NewShippingAddress(
			"customer-123", "Home", "John", "Doe", "",
			"123 Main St", "", "New York", "", "10001", "US", "", false,
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyState, err)
	})
	
	t.Run("cannot create address with empty postal code", func(t *testing.T) {
		_, err := NewShippingAddress(
			"customer-123", "Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "", "US", "", false,
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyPostalCode, err)
	})
	
	t.Run("cannot create address with empty country code", func(t *testing.T) {
		_, err := NewShippingAddress(
			"customer-123", "Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "", "", false,
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyCountryCode, err)
	})
	
	t.Run("cannot create address with invalid country code", func(t *testing.T) {
		_, err := NewShippingAddress(
			"customer-123", "Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "USA", "", false,
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidCountryCode, err)
	})
	
	t.Run("cannot create address with invalid phone", func(t *testing.T) {
		_, err := NewShippingAddress(
			"customer-123", "Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "invalid-phone", false,
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidPhone, err)
	})
	
	t.Run("cannot create address with invalid label", func(t *testing.T) {
		_, err := NewShippingAddress(
			"customer-123", "InvalidLabel", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidAddressLabel, err)
	})
	
	t.Run("get full name", func(t *testing.T) {
		addr := createTestShippingAddress("customer-123")
		
		assert.Equal(t, "John Doe", addr.GetFullName())
	})
	
	t.Run("get full address", func(t *testing.T) {
		addr := createTestShippingAddress("customer-123")
		
		fullAddress := addr.GetFullAddress()
		assert.Contains(t, fullAddress, "Acme Corp")
		assert.Contains(t, fullAddress, "123 Main St")
		assert.Contains(t, fullAddress, "Apt 4B")
		assert.Contains(t, fullAddress, "New York, NY 10001")
		assert.Contains(t, fullAddress, "US")
	})
	
	t.Run("update label", func(t *testing.T) {
		addr := createTestShippingAddress("customer-123")
		
		err := addr.UpdateLabel("Work")
		assert.NoError(t, err)
		assert.Equal(t, "Work", addr.Label)
	})
	
	t.Run("cannot update to invalid label", func(t *testing.T) {
		addr := createTestShippingAddress("customer-123")
		
		err := addr.UpdateLabel("InvalidLabel")
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidAddressLabel, err)
	})
	
	t.Run("set and unset default", func(t *testing.T) {
		addr := createTestShippingAddress("customer-123")
		
		addr.UnsetAsDefault()
		assert.False(t, addr.IsDefault)
		
		addr.SetAsDefault()
		assert.True(t, addr.IsDefault)
	})
	
	t.Run("validate for shipping", func(t *testing.T) {
		addr := createTestShippingAddress("customer-123")
		
		err := addr.ValidateForShipping()
		assert.NoError(t, err)
	})
}

// Tests for Customer Status

func TestCustomerStatus(t *testing.T) {
	t.Run("valid statuses", func(t *testing.T) {
		validStatuses := []CustomerStatus{
			StatusActive, StatusInactive, StatusSuspended,
		}
		
		for _, status := range validStatuses {
			assert.True(t, status.IsValid())
		}
	})
	
	t.Run("invalid status", func(t *testing.T) {
		invalidStatus := CustomerStatus("INVALID")
		assert.False(t, invalidStatus.IsValid())
	})
}

// Tests for Customer Entity

func TestNewCustomer(t *testing.T) {
	t.Run("create valid customer", func(t *testing.T) {
		customer, err := NewCustomer("test@example.com", "John", "Doe", "+1234567890")
		
		assert.NoError(t, err)
		assert.NotEmpty(t, customer.ID)
		assert.Equal(t, "test@example.com", customer.Email.Address)
		assert.Equal(t, "John", customer.FirstName)
		assert.Equal(t, "Doe", customer.LastName)
		assert.Equal(t, "+1234567890", customer.Phone)
		assert.Equal(t, StatusActive, customer.Status)
		assert.False(t, customer.CreatedAt.IsZero())
		assert.False(t, customer.UpdatedAt.IsZero())
		assert.Empty(t, customer.ShippingAddresses)
	})
	
	t.Run("cannot create customer with invalid email", func(t *testing.T) {
		_, err := NewCustomer("invalid-email", "John", "Doe", "")
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidEmail, err)
	})
	
	t.Run("cannot create customer with empty first name", func(t *testing.T) {
		_, err := NewCustomer("test@example.com", "", "Doe", "")
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyFirstName, err)
	})
	
	t.Run("cannot create customer with empty last name", func(t *testing.T) {
		_, err := NewCustomer("test@example.com", "John", "", "")
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyLastName, err)
	})
	
	t.Run("cannot create customer with invalid phone", func(t *testing.T) {
		_, err := NewCustomer("test@example.com", "John", "Doe", "invalid-phone")
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidPhone, err)
	})
	
	t.Run("create customer without phone", func(t *testing.T) {
		customer, err := NewCustomer("test@example.com", "John", "Doe", "")
		
		assert.NoError(t, err)
		assert.Equal(t, "", customer.Phone)
	})
}

func TestCustomerEmailUpdate(t *testing.T) {
	t.Run("update email", func(t *testing.T) {
		customer := createTestCustomer()
		oldUpdateTime := customer.UpdatedAt
		time.Sleep(1 * time.Millisecond)
		
		err := customer.UpdateEmail("newemail@example.com")
		
		assert.NoError(t, err)
		assert.Equal(t, "newemail@example.com", customer.Email.Address)
		assert.True(t, customer.UpdatedAt.After(oldUpdateTime))
	})
	
	t.Run("cannot update to invalid email", func(t *testing.T) {
		customer := createTestCustomer()
		
		err := customer.UpdateEmail("invalid-email")
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidEmail, err)
	})
}

func TestCustomerPersonalInfoUpdate(t *testing.T) {
	t.Run("update personal info", func(t *testing.T) {
		customer := createTestCustomer()
		oldUpdateTime := customer.UpdatedAt
		time.Sleep(1 * time.Millisecond)
		
		err := customer.UpdatePersonalInfo("Jane", "Smith", "+9876543210")
		
		assert.NoError(t, err)
		assert.Equal(t, "Jane", customer.FirstName)
		assert.Equal(t, "Smith", customer.LastName)
		assert.Equal(t, "+9876543210", customer.Phone)
		assert.True(t, customer.UpdatedAt.After(oldUpdateTime))
	})
	
	t.Run("cannot update with empty first name", func(t *testing.T) {
		customer := createTestCustomer()
		
		err := customer.UpdatePersonalInfo("", "Smith", "")
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyFirstName, err)
	})
	
	t.Run("cannot update with empty last name", func(t *testing.T) {
		customer := createTestCustomer()
		
		err := customer.UpdatePersonalInfo("Jane", "", "")
		
		assert.Error(t, err)
		assert.Equal(t, ErrEmptyLastName, err)
	})
	
	t.Run("cannot update with invalid phone", func(t *testing.T) {
		customer := createTestCustomer()
		
		err := customer.UpdatePersonalInfo("Jane", "Smith", "invalid-phone")
		
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidPhone, err)
	})
}

func TestCustomerStatusTransitions(t *testing.T) {
	t.Run("activate customer", func(t *testing.T) {
		customer := createTestCustomer()
		customer.Status = StatusInactive
		
		err := customer.Activate()
		
		assert.NoError(t, err)
		assert.Equal(t, StatusActive, customer.Status)
	})
	
	t.Run("cannot activate already active customer", func(t *testing.T) {
		customer := createTestCustomer()
		
		err := customer.Activate()
		
		assert.Error(t, err)
		assert.Equal(t, ErrCustomerAlreadyActive, err)
	})
	
	t.Run("deactivate customer", func(t *testing.T) {
		customer := createTestCustomer()
		
		err := customer.Deactivate()
		
		assert.NoError(t, err)
		assert.Equal(t, StatusInactive, customer.Status)
	})
	
	t.Run("cannot deactivate already inactive customer", func(t *testing.T) {
		customer := createTestCustomer()
		customer.Status = StatusInactive
		
		err := customer.Deactivate()
		
		assert.Error(t, err)
		assert.Equal(t, ErrCustomerAlreadyInactive, err)
	})
	
	t.Run("suspend customer", func(t *testing.T) {
		customer := createTestCustomer()
		
		err := customer.Suspend()
		
		assert.NoError(t, err)
		assert.Equal(t, StatusSuspended, customer.Status)
	})
}

func TestCustomerShippingAddresses(t *testing.T) {
	t.Run("add first shipping address", func(t *testing.T) {
		customer := createTestCustomer()
		
		err := customer.AddShippingAddress(
			"Home", "John", "Doe", "Acme Corp",
			"123 Main St", "Apt 4B", "New York", "NY", "10001", "US", "+1234567890", false,
		)
		
		assert.NoError(t, err)
		assert.Len(t, customer.ShippingAddresses, 1)
		
		// First address should be default regardless of the flag
		addr := customer.ShippingAddresses[0]
		assert.True(t, addr.IsDefault)
		assert.NotEmpty(t, addr.ID)
		assert.Equal(t, customer.ID, addr.CustomerID)
	})
	
	t.Run("add second shipping address", func(t *testing.T) {
		customer := createTestCustomer()
		
		// Add first address
		customer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		
		// Add second address as default
		err := customer.AddShippingAddress(
			"Work", "John", "Doe", "Work Corp",
			"456 Work Ave", "", "Boston", "MA", "02101", "US", "", true,
		)
		
		assert.NoError(t, err)
		assert.Len(t, customer.ShippingAddresses, 2)
		
		// First address should no longer be default
		assert.False(t, customer.ShippingAddresses[0].IsDefault)
		// Second address should be default
		assert.True(t, customer.ShippingAddresses[1].IsDefault)
	})
	
	t.Run("update shipping address", func(t *testing.T) {
		customer := createTestCustomer()
		customer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		
		addressID := customer.ShippingAddresses[0].ID
		
		err := customer.UpdateShippingAddress(
			addressID, "Work", "Jane", "Smith", "New Corp",
			"789 New St", "Suite 10", "Chicago", "IL", "60601", "US", "+9876543210",
		)
		
		assert.NoError(t, err)
		
		addr := customer.ShippingAddresses[0]
		assert.Equal(t, "Work", addr.Label)
		assert.Equal(t, "Jane", addr.FirstName)
		assert.Equal(t, "Smith", addr.LastName)
		assert.Equal(t, "New Corp", addr.Company)
		assert.Equal(t, "789 New St", addr.AddressLine1)
		assert.Equal(t, "Suite 10", addr.AddressLine2)
		assert.Equal(t, "Chicago", addr.City)
		assert.Equal(t, "IL", addr.State)
		assert.Equal(t, "60601", addr.PostalCode)
		assert.Equal(t, "+9876543210", addr.Phone)
	})
	
	t.Run("cannot update non-existent address", func(t *testing.T) {
		customer := createTestCustomer()
		
		err := customer.UpdateShippingAddress(
			"non-existent", "Work", "Jane", "Smith", "",
			"789 New St", "", "Chicago", "IL", "60601", "US", "",
		)
		
		assert.Error(t, err)
		assert.Equal(t, ErrAddressNotFound, err)
	})
	
	t.Run("remove shipping address", func(t *testing.T) {
		customer := createTestCustomer()
		
		// Add two addresses
		customer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		customer.AddShippingAddress(
			"Work", "John", "Doe", "",
			"456 Work Ave", "", "Boston", "MA", "02101", "US", "", false,
		)
		
		addressID := customer.ShippingAddresses[1].ID
		
		err := customer.RemoveShippingAddress(addressID)
		
		assert.NoError(t, err)
		assert.Len(t, customer.ShippingAddresses, 1)
		assert.Equal(t, "Home", customer.ShippingAddresses[0].Label)
	})
	
	t.Run("cannot remove only shipping address", func(t *testing.T) {
		customer := createTestCustomer()
		customer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		
		addressID := customer.ShippingAddresses[0].ID
		
		err := customer.RemoveShippingAddress(addressID)
		
		assert.Error(t, err)
		assert.Equal(t, ErrCustomerHasNoAddresses, err)
	})
	
	t.Run("removing default address makes first remaining address default", func(t *testing.T) {
		customer := createTestCustomer()
		
		// Add two addresses, second as default
		customer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		customer.AddShippingAddress(
			"Work", "John", "Doe", "",
			"456 Work Ave", "", "Boston", "MA", "02101", "US", "", true,
		)
		
		// Remove the default address (work)
		workAddressID := customer.ShippingAddresses[1].ID
		err := customer.RemoveShippingAddress(workAddressID)
		
		assert.NoError(t, err)
		assert.Len(t, customer.ShippingAddresses, 1)
		assert.True(t, customer.ShippingAddresses[0].IsDefault) // Home should now be default
	})
	
	t.Run("set default shipping address", func(t *testing.T) {
		customer := createTestCustomer()
		
		// Add two addresses
		customer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		customer.AddShippingAddress(
			"Work", "John", "Doe", "",
			"456 Work Ave", "", "Boston", "MA", "02101", "US", "", false,
		)
		
		workAddressID := customer.ShippingAddresses[1].ID
		
		err := customer.SetDefaultShippingAddress(workAddressID)
		
		assert.NoError(t, err)
		assert.False(t, customer.ShippingAddresses[0].IsDefault)
		assert.True(t, customer.ShippingAddresses[1].IsDefault)
	})
	
	t.Run("cannot set non-existent address as default", func(t *testing.T) {
		customer := createTestCustomer()
		
		err := customer.SetDefaultShippingAddress("non-existent")
		
		assert.Error(t, err)
		assert.Equal(t, ErrAddressNotFound, err)
	})
	
	t.Run("get default shipping address", func(t *testing.T) {
		customer := createTestCustomer()
		customer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		
		addr, err := customer.GetDefaultShippingAddress()
		
		assert.NoError(t, err)
		assert.NotNil(t, addr)
		assert.Equal(t, "Home", addr.Label)
		assert.True(t, addr.IsDefault)
	})
	
	t.Run("get shipping address by ID", func(t *testing.T) {
		customer := createTestCustomer()
		customer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		
		addressID := customer.ShippingAddresses[0].ID
		addr, err := customer.GetShippingAddress(addressID)
		
		assert.NoError(t, err)
		assert.NotNil(t, addr)
		assert.Equal(t, addressID, addr.ID)
	})
	
	t.Run("cannot get non-existent address", func(t *testing.T) {
		customer := createTestCustomer()
		
		_, err := customer.GetShippingAddress("non-existent")
		
		assert.Error(t, err)
		assert.Equal(t, ErrAddressNotFound, err)
	})
}

func TestCustomerQueryMethods(t *testing.T) {
	t.Run("query methods on active customer", func(t *testing.T) {
		customer := createTestCustomer()
		
		assert.True(t, customer.IsActive())
		assert.False(t, customer.IsInactive())
		assert.False(t, customer.IsSuspended())
		assert.Equal(t, "John Doe", customer.GetFullName())
		assert.Equal(t, "test@example.com", customer.GetEmailAddress())
		assert.False(t, customer.HasShippingAddresses())
		assert.Equal(t, 0, customer.GetShippingAddressCount())
		assert.False(t, customer.CanPlaceOrder()) // No shipping addresses
	})
	
	t.Run("query methods on customer with addresses", func(t *testing.T) {
		customer := createTestCustomer()
		customer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		
		assert.True(t, customer.HasShippingAddresses())
		assert.Equal(t, 1, customer.GetShippingAddressCount())
		assert.True(t, customer.CanPlaceOrder()) // Active and has addresses
	})
	
	t.Run("inactive customer cannot place order", func(t *testing.T) {
		customer := createTestCustomer()
		customer.Status = StatusInactive
		customer.AddShippingAddress(
			"Home", "John", "Doe", "",
			"123 Main St", "", "New York", "NY", "10001", "US", "", false,
		)
		
		assert.False(t, customer.CanPlaceOrder()) // Has addresses but inactive
	})
}
