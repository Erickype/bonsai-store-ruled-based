package subcategory

import (
	"context"

	"github.com/Erickype/bonsai-store-ruled-based/category"
)

// AddRequest is the data required for Add endpoint, the Category field
// must exist
type AddRequest struct {
	// Category is the parent category of the product
	Category int64 `json:"category"`
	// Name is a descriptive name of the subcategory
	Name string `json:"name"`
	// Description is a short description of the subcategory
	Description string `json:"description"`
}

// AddResponse is the response for Add endpoint
type AddResponse struct {
	// RowsAffected return 1 if the record was inserted
	RowsAffected int64 `json:"rows_affected,omitempty"`
}

// Add creates a subcategory, the parent category must exist to create a subcategory.
//
// encore:api method=POST path=/subcategory
func Add(ctx context.Context, subcategory *AddRequest) (*AddResponse, error) {
	response := &AddResponse{}
	_, err := category.GetById(ctx, subcategory.Category)
	if err != nil {
		return nil, err
	}
	query := `insert into public.subcategory (category, name, description)
			  values ($1, $2, $3)`
	res, err := subcategoryDb.Exec(ctx, query, subcategory.Category, subcategory.Name, subcategory.Description)
	if err != nil {
		return nil, err
	}
	response.RowsAffected = res.RowsAffected()
	return response, nil
}
