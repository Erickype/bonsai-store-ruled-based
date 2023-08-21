package rule

import (
	"context"

	"encore.dev/beta/errs"
)

// Update modifies the values of a rule by its ID
//
//encore:api public method=PUT path=/rule
func Update(ctx context.Context, rule *Rule) error {
	query := `update rule 
			SET name = $1, "desc" = $2, salience = $3, "when" = $4, "then" = $5 
			WHERE id = $6`
	result, err := ruleDb.Exec(ctx, query, rule.Name, rule.Desc, rule.Salience, rule.When, rule.Then, rule.ID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return &errs.Error{
			Code:    errs.NotFound,
			Message: "rule not updated",
		}
	}
	return nil
}
