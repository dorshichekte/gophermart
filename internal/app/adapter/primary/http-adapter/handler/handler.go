package handler

import (
	"go.uber.org/zap"

	balance_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/balance"
	base_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/base"
	order_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/order"
	user_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/user"
	withdraw_hanlder "gophermarket/internal/app/adapter/primary/http-adapter/handler/withdrawal"
	"gophermarket/internal/app/application/usecase"
	v "gophermarket/internal/libs/validator"
)

func New(logger *zap.Logger, service *usecase.UseCases, validator *v.Validator) *Handlers {
	bh := base_handler.New(logger)

	return &Handlers{
		Balance:    balance_handler.New(logger, service.Balance, validator),
		Order:      order_handler.New(logger, service.Order, validator),
		User:       user_handler.New(bh, logger, service.User, validator),
		Withdrawal: withdraw_hanlder.New(bh, logger, service.Withdrawal, validator),
	}
}
