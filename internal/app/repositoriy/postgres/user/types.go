package userrepositorypostgres

import "database/sql"

type userPostgresRepository struct {
	db *sql.DB
}
