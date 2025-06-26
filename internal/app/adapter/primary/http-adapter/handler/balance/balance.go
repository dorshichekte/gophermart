package balancehandler

import (
	"go.uber.org/zap"

	balanceusecase "gophermarket/internal/app/application/usecase/balance"
	v "gophermarket/internal/libs/validator"
)

func New(logger *zap.Logger, serviceBalance *balanceusecase.BalanceUseCase, validator *v.Validator) *Handler {
	return &Handler{
		ServiceBalance: serviceBalance,
		Logger:         logger,
		validator:      validator,
	}
}
