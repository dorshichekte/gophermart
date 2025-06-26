package userhandler

import (
	"go.uber.org/zap"

	base_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/base"
	userusecase "gophermarket/internal/app/application/usecase/user"
	v "gophermarket/internal/libs/validator"
)

func New(bh base_handler.BaseHandler, logger *zap.Logger, service *userusecase.UserUseCase, validator *v.Validator) *Handler {
	return &Handler{
		BaseHandler: bh,
		ServiceUser: service,
		Logger:      logger,
		validator:   validator,
	}
}
