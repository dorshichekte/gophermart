package orderrepositorypostgres

import "database/sql"

type orderPostgresRepository struct {
	db *sql.DB
}
