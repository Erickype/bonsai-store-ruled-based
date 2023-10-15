package product

// AddRequest is the response struct for Add method
type AddRequest struct {
	// Subcategory is the subcategory the product belongs to.
	Subcategory int64 `json:"subcategory,omitempty"`
	// Name is a descriptve name of the product to insert
	Name string `json:"name,omitempty"`
	// Description is a short description about the product
	Description string `json:"description,omitempty"`
}

// AddResponse is the response struct for Add method
type AddResponse struct {
	// ProductID is the generated id from the product inserted
	ProductID int64 `json:"product_id,omitempty"`
}
