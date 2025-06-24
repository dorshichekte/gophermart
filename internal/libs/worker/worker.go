package worker

import (
	envconfig "gophermarket/internal/app/config/env"

	"go.uber.org/zap"
)

func New(cfg envconfig.Config, logger *zap.Logger) *Worker {
	return &Worker{Cfg: cfg, Logger: logger}
}
