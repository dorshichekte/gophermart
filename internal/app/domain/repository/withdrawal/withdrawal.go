package withdrawalrepository

import (
	"context"
	model "gophermarket/internal/app/repositoriy/model/withdrawal"
)

type WithdrawalRepository interface {
	Get(ctx context.Context, userID int) ([]model.Withdrawal, error)
	Make(ctx context.Context, orderNum string, userID int, amount float64) error
}
