package repository_postgres

import (
	"database/sql"

	"go.uber.org/zap"

	"gophermarket/internal/app/config/env"
)

type Postgres struct {
	l      *zap.Logger
	config env_config.Config
	DB     *sql.DB
}
