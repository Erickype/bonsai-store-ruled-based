package subcategory

import "context"

// GetByID retuns a Subcateogry by a param ID
//
//encore:api method=GET public path=/subcategory/:id
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
