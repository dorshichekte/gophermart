package withdrawalhandler

import (
	"go.uber.org/zap"

	base_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/base"
	withdrawalusecase "gophermarket/internal/app/application/usecase/withdrawal"
	v "gophermarket/internal/libs/validator"
)

func New(bh base_handler.BaseHandler, logger *zap.Logger, service *withdrawalusecase.WithdrawalUseCase, validator *v.Validator) *Handler {
	return &Handler{
		BaseHandler:       bh,
		ServiceWithdrawal: service,
		Logger:            logger,
		validator:         validator,
	}
}
