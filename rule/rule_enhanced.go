package rule

type BetterRule struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Salience int    `json:"salience"`
	When     WhenB   `json:"when"`
	Then     []ThenB `json:"then"`
}
type Expresion struct {
	Obj   string `json:"obj,omitempty"`
	Const bool   `json:"const,omitempty"`
}
type Operator struct {
	Name      string      `json:"name"`
	Expresion []Expresion `json:"expresion"`
}
type Conditions struct {
	Operator Operator `json:"operator"`
}
type CompoundCondition struct {
	Name       string       `json:"name"`
	Conditions []Conditions `json:"conditions"`
}
type WhenB struct {
	CompoundCondition CompoundCondition `json:"compoundCondition"`
}
type SetB struct {
	Obj      string   `json:"obj,omitempty"`
	Operator Operator `json:"operator,omitempty"`
}
type Call struct {
	Expresion []Expresion `json:"expresion"`
}
type ThenB struct {
	Set  []Set  `json:"set,omitempty"`
	Call []Call `json:"call,omitempty"`
}
