package userrepositorypostgres

import (
	"database/sql"

	user_repository "gophermarket/internal/app/domain/repository/user"
)

func New(db *sql.DB) user_repository.UserRepository {
	return &userPostgresRepository{db: db}
}
