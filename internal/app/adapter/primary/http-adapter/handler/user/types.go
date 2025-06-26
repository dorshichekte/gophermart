package userhandler

import (
	"go.uber.org/zap"

	base_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/base"
	userusecase "gophermarket/internal/app/application/usecase/user"
	v "gophermarket/internal/libs/validator"
)

type Handler struct {
	base_handler.BaseHandler
	Logger      *zap.Logger
	ServiceUser *userusecase.UserUseCase
	validator   *v.Validator
}
