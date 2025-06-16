package usecase

import (
	"go.uber.org/zap"

	balance_usecase "gophermarket/internal/app/application/usecase/balance"
	order_usecase "gophermarket/internal/app/application/usecase/order"
	user_usecase "gophermarket/internal/app/application/usecase/user"
	withdrawal_usecase "gophermarket/internal/app/application/usecase/withdrawal"
	"gophermarket/internal/libs/auth"
	"gophermarket/internal/libs/hasher"
)

func New(logger *zap.Logger, hasher hasher.Hasher, auth auth.Auth, repos Repositories) *UseCases {
	return &UseCases{
		User:       user_usecase.New(logger, hasher, auth, repos.User),
		Balance:    balance_usecase.New(logger, repos.Balance),
		Order:      order_usecase.New(logger, repos.Order),
		Withdrawal: withdrawal_usecase.New(logger, repos.Withdrawal, repos.Balance),
	}
}
