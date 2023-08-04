package rule

import (
	"context"
)

// Rule is a representation of the GRule json format, used only the basic format.
// https://github.com/hyperjumptech/grule-rule-engine/blob/master/docs/en/GRL_JSON_en.md
type Rule struct {
	Name     string   `json:"name"`
	Desc     string   `json:"desc"`
	Salience int64    `json:"salience"`
	When     string   `json:"when"`
	Then     []string `json:"then"`
}

// Add inserts a Rule in the database
//
//encore:api public method=POST path=/rule
func Add(ctx context.Context, rule *Rule) error {
	query := `insert into rule (name, "desc", salience, "when", "then")
			values($1, $2, $3, $4, $5)`
	_, err := ruleDb.Exec(ctx, query, rule.Name, rule.Desc, rule.Salience, rule.When, rule.Then)
	return err
}
