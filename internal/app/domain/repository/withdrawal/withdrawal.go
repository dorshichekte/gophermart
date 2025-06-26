package withdrawalrepository

import (
	"context"
	model "gophermarket/internal/app/repositoriy/model/withdrawal"
)

//go:generate mockgen -package mocks -source withdrawal.go -destination ../../mock/withdrawal_repository.go WithdrawalRepository
type WithdrawalRepository interface {
	Get(ctx context.Context, userID int) ([]model.Withdrawal, error)
	Make(ctx context.Context, orderNum string, userID int, amount float64) error
}
