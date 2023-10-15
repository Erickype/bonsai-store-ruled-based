package subcategory

// Subcategory is representation of child from a category, the category must exist.
type Subcategory struct {
	// ID is the identifier of the category
	ID int64 `json:"id,omitempty"`
	// Category is the identifier of the parent category
	Category int64 `json:"category,omitempty"`
	// Name is a descriotive name of the category
	Name string `json:"name,omitempty"`
	// Description is a short description of the subacategory
	Description string `json:"description,omitempty"`
}
