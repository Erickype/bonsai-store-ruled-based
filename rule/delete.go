package rule

import (
	"context"

	"encore.dev/beta/errs"
)

// Delete removes a rule from the database base on its ID
//
//encore:api public method=DELETE path=/rule/:id
func Delete(ctx context.Context, id int) error {
	query := `delete from rule
			where id = $1`
	result, err := ruleDb.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return &errs.Error{
			Code:    errs.NotFound,
			Message: "rule not deleted",
		}
	}
	return nil
}
