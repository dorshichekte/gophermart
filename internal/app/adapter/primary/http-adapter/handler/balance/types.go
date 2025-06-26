package balancehandler

import (
	"go.uber.org/zap"

	balanceusecase "gophermarket/internal/app/application/usecase/balance"
	v "gophermarket/internal/libs/validator"
)

type Handler struct {
	Logger         *zap.Logger
	ServiceBalance *balanceusecase.BalanceUseCase
	validator      *v.Validator
}
