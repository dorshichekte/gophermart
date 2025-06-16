package balance_handler

import (
	"go.uber.org/zap"

	"gophermarket/internal/app/application/usecase"
	v "gophermarket/internal/libs/validator"
)

type Handler struct {
	Logger    *zap.Logger
	Service   *usecase.UseCases
	validator *v.Validator
}
