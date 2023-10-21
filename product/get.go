package product

import (
	"context"
	"strconv"
	"strings"

	"encore.dev/storage/sqldb"
	"github.com/Erickype/bonsai-store-ruled-based/subcategory"
)

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
	query := `select id, subcategory, name, description
			  from product`
	rows, err := productDb.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	return processQueryResponse(rows)
}

// GetByCategory returns the list of products by a category ID passed in path
//
//encore:api public method=GET path=/product/category/:categoryID
func GetByCategory(ctx context.Context, categoryID int64) (*ProductsResponse, error) {
	subcatResponse, err := subcategory.GetByCategory(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	var subCatIDs []int64
	for _, subcategory := range subcatResponse.Subcategories {
		subCatIDs = append(subCatIDs, subcategory.ID)
	}
	// Create a parameterized query with the appropriate number of placeholders
	query := "SELECT id, subcategory, name, description FROM product WHERE subcategory IN ("
	placeholders := make([]string, len(subCatIDs))
	args := make([]interface{}, len(subCatIDs))

	for i := range subCatIDs {
		placeholders[i] = "$" + strconv.Itoa(i+1)
		args[i] = subCatIDs[i]
	}
	query += strings.Join(placeholders, ", ") + ")"
	// Now you can execute the query with the IDs
	rows, err := productDb.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return processQueryResponse(rows)
}

func processQueryResponse(rows *sqldb.Rows) (*ProductsResponse, error) {
	response := &ProductsResponse{}
	for rows.Next() {
		product := &Product{}
		err := rows.Scan(&product.ID, &product.Subcategory, &product.Name, &product.Description)
		if err != nil {
			return nil, err
		}
		response.Products = append(response.Products, product)
	}
	return response, nil
}
