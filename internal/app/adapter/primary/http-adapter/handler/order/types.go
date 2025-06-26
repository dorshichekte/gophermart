package orderhandler

import (
	"go.uber.org/zap"

	orderusecase "gophermarket/internal/app/application/usecase/order"
	v "gophermarket/internal/libs/validator"
)

type Handler struct {
	Logger       *zap.Logger
	ServiceOrder *orderusecase.OrderUseCase
	validator    *v.Validator
}
