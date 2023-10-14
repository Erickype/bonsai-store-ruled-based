package category

import (
	"context"
)

// AddResponse is the response struct for Add method
type AddResponse struct {
	// RowsAffected return 1 if the category was created
	RowsAffected int64 `json:"rowsAffected"`
}

// AddRequest is the request struct for Add method
type AddRequest struct {
	// Name is the name of the category
	Name string `json:"name"`
	// Description is a small description of the category
	Description string `json:"description"`
}

// Add insert a category in the database
//
//encore:api method=POST path=/subcategory/category
func Add(ctx context.Context, category *AddRequest) (*AddResponse, error) {
	response := &AddResponse{}
	query := `insert into public.category (name, description)
			  values ($1, $2)`
	res, err := categoryDb.Exec(ctx, query, category.Name, category.Description)
	if err != nil {
		return nil, err
	}
	response.RowsAffected = res.RowsAffected()
	return response, nil
}
