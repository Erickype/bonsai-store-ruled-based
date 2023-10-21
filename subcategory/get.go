package subcategory

import (
	"context"

	"github.com/Erickype/bonsai-store-ruled-based/category"
)

// GetByID retuns a Subcateogry by a param ID
//
//encore:api method=GET public path=/subcategory/id/:id
func GetByID(ctx context.Context, id int64) (*Subcategory, error) {
	subcategory := &Subcategory{}
	query := `select id, category, name, description
			  from subcategory
			  where id = $1`
	row := subcategoryDb.QueryRow(ctx, query, id)
	err := row.Scan(&subcategory.ID, &subcategory.Category, &subcategory.Name, &subcategory.Description)
	if err != nil {
		return nil, err
	}
	return subcategory, nil
}

type SubcategoriesResponse struct {
	Subcategories []*Subcategory
}

//encore:api method=GET path=/subcategory/category/:categoryID
func GetByCategory(ctx context.Context, categoryID int64) (*SubcategoriesResponse, error) {
	response := &SubcategoriesResponse{}
	_, err := category.GetById(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	query := `select id, category, name, description
			  from subcategory
			  where category = $1`
	rows, err := subcategoryDb.Query(ctx, query, categoryID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		subcategory := &Subcategory{}
		err = rows.Scan(&subcategory.ID, &subcategory.Category, &subcategory.Name, &subcategory.Description)
		if err != nil {
			return nil, err
		}
		response.Subcategories = append(response.Subcategories, subcategory)
	}
	return response, nil
}
