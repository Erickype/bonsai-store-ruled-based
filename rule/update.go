package rule

import (
	"context"
	"encoding/json"

	"encore.dev/beta/errs"
)

// Update modifies the values of a rule by its ID
//
//encore:api public method=PUT path=/rule
func Update(ctx context.Context, rule *Rule) error {
	id := rule.ID
	rule.ID = 0
	json_rule, err := json.Marshal(rule)
	if err != nil {
		return err
	}
	query := `update rule 
			SET name = $1, "desc" = $2, salience = $3, json_data = $4
			WHERE id = $5`
	result, err := ruleDb.Exec(ctx, query, rule.Name, rule.Desc, rule.Salience, json_rule, id)
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
