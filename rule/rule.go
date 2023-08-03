package rule

import (
	"context"
	"fmt"
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
func Add(_ context.Context, rule *Rule) error {
	fmt.Print(rule)
	return nil
}
