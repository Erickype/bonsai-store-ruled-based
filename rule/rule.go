package rule

import "encoding/json"

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
	// When are the conditions of the rule, there is a compoud condition that are a map of operators with a PairExpresion
	When map[string][]map[string][]PairExpresion `json:"when,omitempty"`
	// Then represents the action when the condition is true
	Then []Then `json:"then,omitempty"`
}

// PairExpresion represents the inner values of an operator, this could be the convination of "obj" and "const"
type PairExpresion struct {
	// Obj is the object to apply the rule, it can acces a property of an object using [object].[syntax]
	Obj string `json:"obj,omitempty"`
	// Const value to be applied. Values accepted: string, int, float, bool
	Const json.RawMessage `json:"const,omitempty"`
}

type Then struct {
	Set  []SetExp          `json:"set,omitempty"`
	Call []json.RawMessage `json:"call,omitempty"`
}

type SetExp struct {
	Obj   string          `json:"obj,omitempty"`
	Plus  []PairExpresion `json:"plus,omitempty"`
	Minus []PairExpresion `json:"minus,omitempty"`
	Div   []PairExpresion `json:"div,omitempty"`
	Mul   []PairExpresion `json:"mul,omitempty"`
	Mod   []PairExpresion `json:"mod,omitempty"`
}
