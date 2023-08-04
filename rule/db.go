package rule

import "encore.dev/storage/sqldb"

var ruleDb = sqldb.NewDatabase("rule", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})
