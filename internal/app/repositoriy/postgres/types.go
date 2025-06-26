package repositorypostgres

import (
	"database/sql"

	"go.uber.org/zap"

	"gophermarket/internal/app/config/env"
)

type Postgres struct {
	l      *zap.Logger
	config config.Env
	DB     *sql.DB
}
