package customer

import "errors"

// Customer domain errors organized by category

// === Validation Errors ===
var (
	ErrInvalidEmail       = errors.New("email address is invalid")
	ErrEmptyEmail         = errors.New("email address cannot be empty")
	ErrEmptyFirstName     = errors.New("first name cannot be empty")
	ErrEmptyLastName      = errors.New("last name cannot be empty")
	ErrInvalidPhone       = errors.New("phone number is invalid")
	ErrEmptyCustomerID    = errors.New("customer ID cannot be empty")
)

// === Address Validation Errors ===
var (
	ErrEmptyAddressLine1  = errors.New("address line 1 cannot be empty")
	ErrEmptyCity          = errors.New("city cannot be empty")
	ErrEmptyState         = errors.New("state cannot be empty")
	ErrEmptyPostalCode    = errors.New("postal code cannot be empty")
	ErrInvalidCountryCode = errors.New("country code must be 2 characters")
	ErrEmptyCountryCode   = errors.New("country code cannot be empty")
	ErrInvalidAddressLabel = errors.New("address label is invalid")
)

// === Business Rule Errors ===
var (
	ErrCustomerNotFound      = errors.New("customer not found")
	ErrDuplicateEmail        = errors.New("email address already exists")
	ErrAddressNotFound       = errors.New("shipping address not found")
	ErrCannotDeleteDefaultAddress = errors.New("cannot delete default shipping address")
	ErrMultipleDefaultAddresses  = errors.New("customer cannot have multiple default addresses")
	ErrCustomerHasNoAddresses    = errors.New("customer must have at least one shipping address")
)

// === State Transition Errors ===
var (
	ErrCustomerAlreadyActive   = errors.New("customer is already active")
	ErrCustomerAlreadyInactive = errors.New("customer is already inactive")
	ErrCannotDeactivateCustomer = errors.New("cannot deactivate customer with pending orders")
)

// === External Service Errors ===
var (
	ErrAddressValidationFailed = errors.New("address validation service failed")
	ErrEmailServiceUnavailable = errors.New("email service is unavailable")
	ErrGeolocationServiceError = errors.New("geolocation service error")
)
