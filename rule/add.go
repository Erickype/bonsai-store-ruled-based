package rule

import "context"

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
