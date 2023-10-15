package product

// Products represent a product from the store, It has a subcategory field.
type Product struct {
	// ID is the identifier of the product
	ID int64 `json:"id,omitempty"`
	// Subcategory represents the subcategory of the product
	Subcategory int64 `json:"subcategory,omitempty"`
	// Name is a descriptive name of the product
	Name string `json:"name,omitempty"`
	// Description is a short description of the product
	Description string `json:"description,omitempty"`
}
