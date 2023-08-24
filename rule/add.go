package rule

import (
	"context"
	"encoding/json"
)

// Add inserts a Rule in the database
//
//encore:api public method=POST path=/rule
func Add(ctx context.Context, rule *Rule) error {
	ruleMarshal, err := json.Marshal(rule)
	if err != nil {
		return err
	}
	query := `insert into rule (name, "desc", salience, "json_data")
			values($1, $2, $3, $4)`
	_, err = ruleDb.Exec(ctx, query, rule.Name, rule.Desc, rule.Salience, ruleMarshal)
	if err != nil {
		return err
	}
	return nil
}
