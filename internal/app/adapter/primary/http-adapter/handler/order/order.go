package orderhandler

import (
	"go.uber.org/zap"

	orderusecase "gophermarket/internal/app/application/usecase/order"
	v "gophermarket/internal/libs/validator"
)

func New(logger *zap.Logger, serviceOrder *orderusecase.OrderUseCase, validator *v.Validator) *Handler {
	return &Handler{
		ServiceOrder: serviceOrder,
		Logger:       logger,
		validator:    validator,
	}
}
