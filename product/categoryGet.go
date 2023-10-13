package product

import "context"

// CategoryGetById returns the information of a category based on the id from path
//
//encore:api public method=GET path=/product/category/:id
func CategoryGetById(ctx context.Context, id int64) (*Category, error) {
	category := &Category{}
	query := `select *
			  from category
			  where id = $1`
	row := categoryDb.QueryRow(ctx, query, id)

	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}

	return category, nil
}
