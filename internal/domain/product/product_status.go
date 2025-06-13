package product

// ProductStatus represents the current state of a product
type ProductStatus string

const (
    StatusActive       ProductStatus = "ACTIVE"
    StatusInactive     ProductStatus = "INACTIVE"
    StatusOutOfStock   ProductStatus = "OUT_OF_STOCK"
    StatusDiscontinued ProductStatus = "DISCONTINUED"
)

// IsValid checks if the product status is valid
func (ps ProductStatus) IsValid() bool {
    switch ps {
    case StatusActive, StatusInactive, StatusOutOfStock, StatusDiscontinued:
        return true
    default:
        return false
    }
}

// CanBeOrdered checks if products with this status can be added to orders
func (ps ProductStatus) CanBeOrdered() bool {
    return ps == StatusActive
}