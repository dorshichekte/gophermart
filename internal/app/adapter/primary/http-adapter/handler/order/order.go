package orderhandler

import (
	"go.uber.org/zap"

	"gophermarket/internal/app/application/usecase"
	v "gophermarket/internal/libs/validator"
)

func New(logger *zap.Logger, service *usecase.UseCases, validator *v.Validator) *Handler {
	return &Handler{
		Service:   service,
		Logger:    logger,
		validator: validator,
	}
}
