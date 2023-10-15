package subcategory

import "encore.dev/storage/sqldb"

var subcategoryDb = sqldb.NewDatabase("subcategory", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})
