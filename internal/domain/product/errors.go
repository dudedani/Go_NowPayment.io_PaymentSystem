package product

import "errors"

// Product domain errors organized by category

// === Validation Errors ===
var (
	ErrEmptyName        = errors.New("product name cannot be empty")
	ErrInvalidPrice     = errors.New("product price must be positive")
	ErrEmptySKU         = errors.New("product SKU cannot be empty")
	ErrInvalidCurrency  = errors.New("product currency is invalid")
	ErrEmptyCategory    = errors.New("product category cannot be empty")
	ErrEmptyDescription = errors.New("product description cannot be empty")
)

// === Business Rule Errors ===
var (
	ErrProductNotActive      = errors.New("product is not active")
	ErrInsufficientStock     = errors.New("insufficient stock available")
	ErrDuplicateSKU         = errors.New("product SKU already exists")
	ErrNegativeStock        = errors.New("stock quantity cannot be negative")
	ErrNegativeReserved     = errors.New("reserved quantity cannot be negative")
	ErrReservedExceedsTotal = errors.New("reserved quantity cannot exceed total quantity")
	ErrNegativeMinimum      = errors.New("minimum stock cannot be negative")
)

// === State Transition Errors ===
var (
	ErrCannotActivateWithoutStock     = errors.New("cannot activate product without stock")
	ErrCannotActivateDiscontinued     = errors.New("cannot activate discontinued product")
	ErrCannotDeactivateInactive       = errors.New("only active products can be deactivated")
	ErrCannotDiscontinueWithReserved  = errors.New("cannot discontinue product with reserved stock")
	ErrCannotUpdateDiscontinued       = errors.New("cannot update discontinued product")
	ErrCannotDeleteActiveProduct      = errors.New("cannot delete active product")
)

// === Inventory Errors ===
var (
	ErrInvalidQuantity                = errors.New("quantity must be positive")
	ErrCannotReserveMoreThanAvailable = errors.New("cannot reserve more than available quantity")
	ErrCannotReleaseMoreThanReserved  = errors.New("cannot release more than reserved quantity")
	ErrCannotFulfillMoreThanReserved  = errors.New("cannot fulfill more than reserved quantity")
	ErrInvalidStockOperation          = errors.New("invalid stock operation")
)

// === Not Found Errors ===
var (
	ErrProductNotFound  = errors.New("product not found")
	ErrCategoryNotFound = errors.New("product category not found")
)

// === Category Errors ===
var (
	ErrInvalidCategoryName = errors.New("category name is invalid")
	ErrCategoryHasProducts = errors.New("cannot delete category with products")
	ErrCircularReference   = errors.New("category cannot be parent of itself")
)