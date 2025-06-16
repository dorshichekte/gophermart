package balance_repository_postgres

import (
	"database/sql"

	balance_repository "gophermarket/internal/app/domain/repository/balance"
)

func New(db *sql.DB) balance_repository.BalanceRepository {
	return &balancePostgresRepository{db: db}
}
