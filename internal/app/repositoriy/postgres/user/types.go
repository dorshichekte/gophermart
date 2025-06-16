package user_repository_postgres

import "database/sql"

type userPostgresRepository struct {
	db *sql.DB
}
