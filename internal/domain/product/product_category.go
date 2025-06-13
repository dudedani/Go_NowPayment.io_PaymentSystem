package product

import (
	"strings"

	"github.com/google/uuid"
)

// Category represents a product category
type Category struct {
	ID          uuid.UUID
	Name        string
	Description string
	ParentID    *uuid.UUID // Optional parent category for hierarchical categories
}

// NewCategory creates a new category with validation
func NewCategory(name, description string, parentID *uuid.UUID) (Category, error) {
	if strings.TrimSpace(name) == "" {
		return Category{}, ErrInvalidCategoryName
	}
	
	category := Category{
		ID:          uuid.New(),
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		ParentID:    parentID,
	}
	
	return category, nil
}

// IsRoot checks if this is a root category (no parent)
func (c Category) IsRoot() bool {
	return c.ParentID == nil
}

// HasParent checks if this category has a parent
func (c Category) HasParent() bool {
	return c.ParentID != nil
}

// GetParentID returns the parent ID if it exists
func (c Category) GetParentID() *uuid.UUID {
	return c.ParentID
}

// UpdateName updates the category name with validation
func (c *Category) UpdateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrInvalidCategoryName
	}
	
	c.Name = strings.TrimSpace(name)
	return nil
}

// UpdateDescription updates the category description
func (c *Category) UpdateDescription(description string) {
	c.Description = strings.TrimSpace(description)
}

// SetParent sets the parent category
func (c *Category) SetParent(parentID uuid.UUID) error {
	// Prevent circular reference (category cannot be parent of itself)
	if parentID == c.ID {
		return ErrCircularReference
	}
	
	c.ParentID = &parentID
	return nil
}

// RemoveParent removes the parent category (makes it a root category)
func (c *Category) RemoveParent() {
	c.ParentID = nil
}

// CategoryPath represents a category path for breadcrumbs
type CategoryPath struct {
	Categories []Category
}

// NewCategoryPath creates a category path
func NewCategoryPath(categories []Category) CategoryPath {
	return CategoryPath{
		Categories: categories,
	}
}

// GetBreadcrumb returns a breadcrumb string
func (cp CategoryPath) GetBreadcrumb(separator string) string {
	if len(cp.Categories) == 0 {
		return ""
	}
	
	var names []string
	for _, category := range cp.Categories {
		names = append(names, category.Name)
	}
	
	return strings.Join(names, separator)
}

// GetDepth returns the depth of the category path
func (cp CategoryPath) GetDepth() int {
	return len(cp.Categories)
}

// IsEmpty checks if the path is empty
func (cp CategoryPath) IsEmpty() bool {
	return len(cp.Categories) == 0
}