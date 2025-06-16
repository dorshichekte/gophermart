package withdrawalrepositorypostgres

import (
	"database/sql"

	withdrawal_repository "gophermarket/internal/app/domain/repository/withdrawal"
)

func New(db *sql.DB) withdrawal_repository.WithdrawalRepository {
	return &withdrawalPostgresRepository{db: db}
}
