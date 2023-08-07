package rule

import (
	"context"
	"encore.dev/beta/errs"
)

// Rule is a representation of the GRule json format, used only the basic format.
// https://github.com/hyperjumptech/grule-rule-engine/blob/master/docs/en/GRL_JSON_en.md
type Rule struct {
	// ID the identifier
	ID int `json:"id,omitempty"`
	// Name represents the rule name
	Name string `json:"name"`
	// Desc is the rule description
	Desc string `json:"desc"`
	// Salience represent the relevance of the rule
	Salience int64 `json:"salience"`
	// When are the conditions of the rule
	When string `json:"when"`
	// Then represents the action when the condition is true
	Then []string `json:"then"`
}

// Add inserts a Rule in the database
//
//encore:api public method=POST path=/rule
func Add(ctx context.Context, rule *Rule) error {
	query := `insert into rule (name, "desc", salience, "when", "then")
			values($1, $2, $3, $4, $5)`
	_, err := ruleDb.Exec(ctx, query, rule.Name, rule.Desc, rule.Salience, rule.When, rule.Then)
	if err != nil {
		return err
	}
	return nil
}

// GetById return a rule based on its ID
//
//encore:api public method=GET path=/rule/:id
func GetById(ctx context.Context, id int) (*Rule, error) {
	query := `select * from rule
			where id = $1`
	row := ruleDb.QueryRow(ctx, query, id)
	rule := &Rule{}
	err := row.Scan(&rule.ID, &rule.Name, &rule.Desc, &rule.Salience, &rule.When, &rule.Then)
	if err != nil {
		return nil, err
	}
	return rule, nil
}

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
