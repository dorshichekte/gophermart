package balancerepository

import (
	"context"

	model "gophermarket/internal/app/repositoriy/model/balance"
)

type BalanceRepository interface {
	Get(ctx context.Context, userID int) (model.Balance, error)
}
