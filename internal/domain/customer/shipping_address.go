package customer

import (
	"strings"
)

// ShippingAddress represents a complete shipping address value object
type ShippingAddress struct {
	ID           string
	CustomerID   string
	Label        string // "Home", "Work", "Office", etc.
	FirstName    string
	LastName     string
	Company      string
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	PostalCode   string
	Country      string // ISO 2-letter country code
	Phone        string
	IsDefault    bool
}

// NewShippingAddress creates a new shipping address with validation
func NewShippingAddress(customerID, label, firstName, lastName, company, addressLine1, addressLine2, city, state, postalCode, country, phone string, isDefault bool) (ShippingAddress, error) {
	// Validate required fields
	if customerID == "" {
		return ShippingAddress{}, ErrEmptyCustomerID
	}
	
	if strings.TrimSpace(firstName) == "" {
		return ShippingAddress{}, ErrEmptyFirstName
	}
	
	if strings.TrimSpace(lastName) == "" {
		return ShippingAddress{}, ErrEmptyLastName
	}
	
	if strings.TrimSpace(addressLine1) == "" {
		return ShippingAddress{}, ErrEmptyAddressLine1
	}
	
	if strings.TrimSpace(city) == "" {
		return ShippingAddress{}, ErrEmptyCity
	}
	
	if strings.TrimSpace(state) == "" {
		return ShippingAddress{}, ErrEmptyState
	}
	
	if strings.TrimSpace(postalCode) == "" {
		return ShippingAddress{}, ErrEmptyPostalCode
	}
	
	if strings.TrimSpace(country) == "" {
		return ShippingAddress{}, ErrEmptyCountryCode
	}
	
	// Validate country code format (2 characters)
	country = strings.ToUpper(strings.TrimSpace(country))
	if len(country) != 2 {
		return ShippingAddress{}, ErrInvalidCountryCode
	}
	
	// Validate address label if provided
	if label != "" && !isValidAddressLabel(label) {
		return ShippingAddress{}, ErrInvalidAddressLabel
	}
	
	// Validate phone if provided
	if phone != "" && !isValidPhone(phone) {
		return ShippingAddress{}, ErrInvalidPhone
	}
	
	return ShippingAddress{
		CustomerID:   customerID,
		Label:        strings.TrimSpace(label),
		FirstName:    strings.TrimSpace(firstName),
		LastName:     strings.TrimSpace(lastName),
		Company:      strings.TrimSpace(company),
		AddressLine1: strings.TrimSpace(addressLine1),
		AddressLine2: strings.TrimSpace(addressLine2),
		City:         strings.TrimSpace(city),
		State:        strings.TrimSpace(state),
		PostalCode:   strings.TrimSpace(postalCode),
		Country:      country,
		Phone:        strings.TrimSpace(phone),
		IsDefault:    isDefault,
	}, nil
}

// isValidAddressLabel validates address label
func isValidAddressLabel(label string) bool {
	validLabels := []string{"Home", "Work", "Office", "Business", "Other"}
	for _, validLabel := range validLabels {
		if strings.EqualFold(label, validLabel) {
			return true
		}
	}
	return false
}

// isValidPhone validates phone number format (simple validation)
func isValidPhone(phone string) bool {
	// Remove common phone number characters
	cleaned := strings.ReplaceAll(phone, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	cleaned = strings.ReplaceAll(cleaned, "(", "")
	cleaned = strings.ReplaceAll(cleaned, ")", "")
	cleaned = strings.ReplaceAll(cleaned, "+", "")
	
	// Check if remaining characters are digits and length is reasonable
	if len(cleaned) < 10 || len(cleaned) > 15 {
		return false
	}
	
	for _, char := range cleaned {
		if char < '0' || char > '9' {
			return false
		}
	}
	
	return true
}

// GetFullName returns the full name for the address
func (sa ShippingAddress) GetFullName() string {
	return sa.FirstName + " " + sa.LastName
}

// GetFullAddress returns the complete address as a formatted string
func (sa ShippingAddress) GetFullAddress() string {
	var parts []string
	
	if sa.Company != "" {
		parts = append(parts, sa.Company)
	}
	
	parts = append(parts, sa.AddressLine1)
	
	if sa.AddressLine2 != "" {
		parts = append(parts, sa.AddressLine2)
	}
	
	cityStateZip := sa.City + ", " + sa.State + " " + sa.PostalCode
	parts = append(parts, cityStateZip)
	parts = append(parts, sa.Country)
	
	return strings.Join(parts, "\n")
}

// UpdateLabel updates the address label
func (sa *ShippingAddress) UpdateLabel(label string) error {
	if label != "" && !isValidAddressLabel(label) {
		return ErrInvalidAddressLabel
	}
	
	sa.Label = strings.TrimSpace(label)
	return nil
}

// SetAsDefault marks this address as the default address
func (sa *ShippingAddress) SetAsDefault() {
	sa.IsDefault = true
}

// UnsetAsDefault removes the default flag from this address
func (sa *ShippingAddress) UnsetAsDefault() {
	sa.IsDefault = false
}

// ValidateForShipping checks if the address has all required fields for shipping
func (sa ShippingAddress) ValidateForShipping() error {
	if sa.FirstName == "" {
		return ErrEmptyFirstName
	}
	
	if sa.LastName == "" {
		return ErrEmptyLastName
	}
	
	if sa.AddressLine1 == "" {
		return ErrEmptyAddressLine1
	}
	
	if sa.City == "" {
		return ErrEmptyCity
	}
	
	if sa.State == "" {
		return ErrEmptyState
	}
	
	if sa.PostalCode == "" {
		return ErrEmptyPostalCode
	}
	
	if sa.Country == "" {
		return ErrEmptyCountryCode
	}
	
	return nil
}
