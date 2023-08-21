package rule

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
	When When `json:"when"`
	// Then represents the action when the condition is true
	Then []Then `json:"then"`
}

type Eq struct {
	Obj   string `json:"obj,omitempty"`
	Const bool   `json:"const,omitempty"`
}
type Lt struct {
	Obj string `json:"obj,omitempty"`
}
type And struct {
	Eq []Eq `json:"eq,omitempty"`
	Lt []Lt `json:"lt,omitempty"`
}
type When struct {
	And []And `json:"and,omitempty"`
}
type Plus struct {
	Obj string `json:"obj,omitempty"`
}
type Set struct {
	Obj  string `json:"obj,omitempty"`
	Plus []Plus `json:"plus,omitempty"`
}

type Then struct {
	Set []Set `json:"set,omitempty"`
}
