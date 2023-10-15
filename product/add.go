package product

import (
	"context"

	"github.com/Erickype/bonsai-store-ruled-based/subcategory"
)

// AddRequest is the response struct for Add method
type AddRequest struct {
	// Subcategory is the subcategory the product belongs to.
	Subcategory int64 `json:"subcategory,omitempty"`
	// Name is a descriptve name of the product to insert
	Name string `json:"name,omitempty"`
	// Description is a short description about the product
	Description string `json:"description,omitempty"`
}

// AddResponse is the response struct for Add method
type AddResponse struct {
	// ProductID is the generated id from the product inserted
	ProductID int64 `json:"product_id,omitempty"`
}

// Add insert a product in a subcategory, the subcategory must exist.
//
//encore:api method=POST path=/product
func Add(ctx context.Context, product *AddRequest) (*AddResponse, error) {
	response := &AddResponse{}
	_, err := subcategory.GetByID(ctx, product.Subcategory)
	if err != nil {
		return nil, err
	}
	query := `insert into product (subcategory, name, description)
			  values ($1, $2, $3)
			  returning id`
	row := productDb.QueryRow(ctx, query, product.Subcategory, product.Name, product.Description)
	err = row.Scan(&response.ProductID)
	if err != nil {
		return nil, err
	}
	return response, nil
}
