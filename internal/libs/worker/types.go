package worker

import (
	envconfig "gophermarket/internal/app/config/env"

	"go.uber.org/zap"
)

type Worker struct {
	Cfg    envconfig.Env
	Logger *zap.Logger
}
