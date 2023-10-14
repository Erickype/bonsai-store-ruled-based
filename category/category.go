package category

// Category represents a product parent category
type Category struct {
	// ID is a serial id
	ID int64 `json:"id"`
	// Name is a name for the category
	Name string `json:"name"`
	// Description is a short description for the category
	Description string `json:"description"`
}
