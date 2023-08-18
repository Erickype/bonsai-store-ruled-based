package rule

type ExtendedRule struct {
	Name     string   `json:"name"`
	Desc     string   `json:"desc"`
	Salience int      `json:"salience"`
	When     When     `json:"when"`
	Then     []Action `json:"then"`
}

type When struct {
	And []Condition `json:"and"`
}

type Condition struct {
	Eq []Expression `json:"eq"`
	Lt []Expression `json:"lt"`
	// Add other condition types as needed
}

type Expression struct {
	Obj   string       `json:"obj"`
	Const interface{}  `json:"const"`
	Plus  []Expression `json:"plus"`
	// Add other expression types as needed
}

type Action struct {
	Set  []Expression  `json:"set"`
	Call []interface{} `json:"call"` // Assumes the second value is an object with const field
	// Add other action types as needed
}
