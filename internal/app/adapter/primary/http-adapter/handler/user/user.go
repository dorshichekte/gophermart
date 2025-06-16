package userhandler

import (
	"go.uber.org/zap"

	base_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/base"
	"gophermarket/internal/app/application/usecase"
	v "gophermarket/internal/libs/validator"
)

func New(bh base_handler.BaseHandler, logger *zap.Logger, service *usecase.UseCases, validator *v.Validator) *Handler {
	return &Handler{
		BaseHandler: bh,
		Service:     service,
		Logger:      logger,
		validator:   validator,
	}
}
