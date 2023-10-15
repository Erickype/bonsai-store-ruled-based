package category

import (
	"context"
)

// GetById returns the information of a category based on the id from path
// returns a Category struct.
//
//encore:api public method=GET path=/category/:id
func GetById(ctx context.Context, id int64) (*Category, error) {
	category := &Category{}
	query := `select *
			  from public.category
			  where id = $1`
	row := categoryDb.QueryRow(ctx, query, id)

	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}

	return category, nil
}
