package product

import "encore.dev/storage/sqldb"

var productDb = sqldb.NewDatabase("product", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})
