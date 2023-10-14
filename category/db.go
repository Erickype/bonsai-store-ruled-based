package category

import "encore.dev/storage/sqldb"

var categoryDb = sqldb.NewDatabase("category", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})
