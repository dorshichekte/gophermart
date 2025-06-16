package httpadapter

import (
	"context"
	"gophermarket/internal/libs/auth"
	"net/http"

	"go.uber.org/zap"

	"gophermarket/internal/app/adapter/primary/http-adapter/handler"
	"gophermarket/internal/app/adapter/primary/http-adapter/router"
	"gophermarket/internal/app/application/usecase"
	"gophermarket/internal/libs/http-server"
	"gophermarket/internal/libs/validator"
)

func New(logger *zap.Logger, auth auth.Auth, config Config, uc *usecase.UseCases, validator *validator.Validator) *HTTPAdapter {
	rtr := newRouter(logger, auth, config, uc, validator)

	s := httpserver.New(logger, config.Server, rtr)

	return &HTTPAdapter{
		server: s,
	}
}

func newRouter(logger *zap.Logger, auth auth.Auth, config Config, uc *usecase.UseCases, validator *validator.Validator) http.Handler {
	r := router.New(logger)

	h := handler.New(logger, uc, validator)

	r.AppendRoutes(config.Router, h, auth)

	return r.Router()
}

func (a HTTPAdapter) Start(ctx context.Context) error {
	return a.server.Start(ctx)
}
