package rule

import "context"

// ListResponse is the response of the Get method
type ListResponse struct {
	// Rules holds the list of rules
	Rules []*Rule
}

// Get return the list of rules, just the ID, name, desc and salience
//
//encore:api public method=GET path=/rule
func Get(ctx context.Context) (*ListResponse, error) {
	query := "select ID, name, 'des', salience from rule"
	rows, err := ruleDb.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	response := &ListResponse{Rules: []*Rule{}}

	for rows.Next() {
		rule := &Rule{}
		err = rows.Scan(&rule.ID, &rule.Name, &rule.Desc, &rule.Salience)
		if err != nil {
			return nil, err
		}
		response.Rules = append(response.Rules, rule)
	}

	return response, nil
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
