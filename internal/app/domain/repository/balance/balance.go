package balancerepository

import (
	"context"

	model "gophermarket/internal/app/repositoriy/model/balance"
)

//go:generate mockgen -package mocks -source balance.go -destination ../../mock/balance_repository.go BalanceRepository
type BalanceRepository interface {
	Get(ctx context.Context, userID int) (model.Balance, error)
}
