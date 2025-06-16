package balance_repository_postgres

import "database/sql"

type balancePostgresRepository struct {
	db *sql.DB
}
