package withdrawalrepositorypostgres

import "database/sql"

type withdrawalPostgresRepository struct {
	db *sql.DB
}
