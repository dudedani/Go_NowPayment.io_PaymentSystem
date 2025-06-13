package customer

import (
	domainCustomer "github.com/dudedani/Go_NowPayment.io_PaymentSystem/internal/domain/customer"
)

// GetCustomerQuery represents the input for retrieving a customer
type GetCustomerQuery struct {
	ID string `json:"id" validate:"required"`
}

// GetCustomerResponse represents the customer details response
type GetCustomerResponse struct {
	ID                string                    `json:"id"`
	Email             string                    `json:"email"`
	FirstName         string                    `json:"first_name"`
	LastName          string                    `json:"last_name"`
	Phone             string                    `json:"phone,omitempty"`
	Status            string                    `json:"status"`
	ShippingAddresses []ShippingAddressResponse `json:"shipping_addresses"`
	CreatedAt         string                    `json:"created_at"`
	UpdatedAt         string                    `json:"updated_at"`
}

// ShippingAddressResponse represents a shipping address in the response
type ShippingAddressResponse struct {
	ID           string `json:"id"`
	Label        string `json:"label"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Company      string `json:"company,omitempty"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2,omitempty"`
	City         string `json:"city"`
	State        string `json:"state"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Phone        string `json:"phone,omitempty"`
	IsDefault    bool   `json:"is_default"`
}

// GetCustomerUseCase handles retrieving customer details
type GetCustomerUseCase struct {
	customerRepo CustomerRepository
}

// NewGetCustomerUseCase creates a new instance of GetCustomerUseCase
func NewGetCustomerUseCase(customerRepo CustomerRepository) *GetCustomerUseCase {
	return &GetCustomerUseCase{
		customerRepo: customerRepo,
	}
}

// Execute retrieves customer details by ID
func (uc *GetCustomerUseCase) Execute(query GetCustomerQuery) (*GetCustomerResponse, error) {
	// Find customer by ID
	customer, err := uc.customerRepo.FindByID(query.ID)
	if err != nil {
		return nil, err
	}
	
	if customer == nil {
		return nil, domainCustomer.ErrCustomerNotFound
	}
	
	// Convert shipping addresses
	addresses := make([]ShippingAddressResponse, len(customer.ShippingAddresses))
	for i, addr := range customer.ShippingAddresses {
		addresses[i] = ShippingAddressResponse{
			ID:           addr.ID,
			Label:        addr.Label,
			FirstName:    addr.FirstName,
			LastName:     addr.LastName,
			Company:      addr.Company,
			AddressLine1: addr.AddressLine1,
			AddressLine2: addr.AddressLine2,
			City:         addr.City,
			State:        addr.State,
			PostalCode:   addr.PostalCode,
			Country:      addr.Country,
			Phone:        addr.Phone,
			IsDefault:    addr.IsDefault,
		}
	}
	
	// Return response
	return &GetCustomerResponse{
		ID:                customer.ID,
		Email:             customer.Email.Address,
		FirstName:         customer.FirstName,
		LastName:          customer.LastName,
		Phone:             customer.Phone,
		Status:            string(customer.Status),
		ShippingAddresses: addresses,
		CreatedAt:         customer.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:         customer.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
