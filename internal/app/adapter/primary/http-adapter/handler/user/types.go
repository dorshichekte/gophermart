package user_handler

import (
	"go.uber.org/zap"

	base_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/base"
	"gophermarket/internal/app/application/usecase"
	v "gophermarket/internal/libs/validator"
)

type Handler struct {
	base_handler.BaseHandler
	Logger    *zap.Logger
	Service   *usecase.UseCases
	validator *v.Validator
}
