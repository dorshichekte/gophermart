package withdrawal_repository_postgres

import "database/sql"

type withdrawalPostgresRepository struct {
	db *sql.DB
}
