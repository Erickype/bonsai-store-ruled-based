package rules

// Rule is a representation of the GRule json format, used only the basic format.
// https://github.com/hyperjumptech/grule-rule-engine/blob/master/docs/en/GRL_JSON_en.md
type Rule struct {
	Name     string   `json:"name"`
	Desc     string   `json:"desc"`
	Salience int64    `json:"salience"`
	When     string   `json:"when"`
	Then     []string `json:"then"`
}
