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

// GetResponse is the response of Get method
type GetResponse struct {
	// Products an array of products
	Products []*Product `json:"products"`
}

// Get returns the list of products
//
//encore:api public method=GET path=/product
func Get(ctx context.Context) (*GetResponse, error) {
	response := &GetResponse{}
	query := `select id, subcategory, name, description
			  from product`
	rows, err := productDb.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		product := &Product{}
		err = rows.Scan(&product.ID, &product.Subcategory, &product.Name, &product.Description)
		if err != nil {
			return nil, err
		}
		response.Products = append(response.Products, product)
	}
	return response, nil
}
