package order_repository_postgres

import "database/sql"

type orderPostgresRepository struct {
	db *sql.DB
}
