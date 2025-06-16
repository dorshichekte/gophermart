package balancerepositorypostgres

import "database/sql"

type balancePostgresRepository struct {
	db *sql.DB
}
