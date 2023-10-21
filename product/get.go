package product

import "context"

// GetByID return a Product by a param ID
//
//encore:api public method=GET path=/product/id/:productID
func GetByID(ctx context.Context, productID int64) (*Product, error) {
	product := &Product{}
	query := `select id, subcategory, name, description
			  from product
			  where id = $1`
	row := productDb.QueryRow(ctx, query, productID)
	err := row.Scan(&product.ID, &product.Subcategory, &product.Name, &product.Description)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// ProductsResponse is a response that returns a list of products, used for various methods.
type ProductsResponse struct {
	// Products an array of products
	Products []*Product `json:"products"`
}

// Get returns the list of products
//
//encore:api public method=GET path=/product
func Get(ctx context.Context) (*ProductsResponse, error) {
	response := &ProductsResponse{}
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

// GetByCategory returns the list of products by a category ID passed in path
//
//encore:api public method=GET path=/product/category/:categoryID
func GetByCategory(ctx context.Context, categoryID int64) (*ProductsResponse, error) {
	response := &ProductsResponse{}

	return response, nil
}
