package order_repository_postgres

import (
	"database/sql"

	order_repository "gophermarket/internal/app/domain/repository/order"
)

func New(db *sql.DB) order_repository.OrderRepository {
	return &orderPostgresRepository{db: db}
}
