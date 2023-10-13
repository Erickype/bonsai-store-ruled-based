package product

import "context"

type CategoryAddResponse struct {
	RowsAffected int64 `json:"rowsAffected"`
}

type CategoryAddRequest struct {
	Name        string `json:"name"`
	Description string `json:"desciption"`
}

// CategoryAdd insert a product category in the database
//
//encore:api method=POST path=/product/category
func CategoryAdd(ctx context.Context, category *CategoryAddRequest) (*CategoryAddResponse, error) {
	response := &CategoryAddResponse{}
	query := `insert into category (name, description)
			  values ($1, $2)`
	res, err := categoryDb.Exec(ctx, query, category.Name, category.Description)
	if err != nil {
		return nil, err
	}
	response.RowsAffected = res.RowsAffected()
	return response, nil
}
