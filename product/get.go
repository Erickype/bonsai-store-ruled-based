package product

import "context"

// GetByID return a Product by a param ID
//
//encore:api public method=GET path=/product/:id
func GetByID(ctx context.Context, id int64) (*Product, error) {
	product := &Product{}
	query := `select id, subcategory, name, description
			  from product
			  where id = $1`
	row := productDb.QueryRow(ctx, query, id)
	err := row.Scan(&product.ID, &product.Subcategory, &product.Name, &product.Description)
	if err != nil {
		return nil, err
	}
	return product, nil
}
