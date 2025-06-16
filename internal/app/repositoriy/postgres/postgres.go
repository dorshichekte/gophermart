package repositorypostgres

import (
	"database/sql"
	"gophermarket/internal/constants"
	customerror "gophermarket/internal/error"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	"gophermarket/internal/app/config/env"
)

func New(l *zap.Logger, cfg envconfig.Config) *sql.DB {
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		l.Fatal(err.Error())
		panic(err)
	}

	err = applyMigrations(cfg)
	if err != nil {
		l.Fatal(err.Error())
		panic(err)
	}

	return db
}

func applyMigrations(cfg envconfig.Config) error {
	wd, err := os.Getwd()
	if err != nil {
		return customerror.NewWithData(constants.ErrPathUnknownFolderPath, err)
	}

	migrationDirPath := "file://" + filepath.Join(wd, "migrations")

	m, err := migrate.New(migrationDirPath, cfg.DatabaseDSN)
	if err != nil {
		return customerror.NewWithData(errDBFailedInitMigrations, err)
	}

	err = m.Up()
	isApplyMigrationFailed := err != nil && err.Error() != "no change"
	if isApplyMigrationFailed {
		return customerror.NewWithData(errDBFailedApplyMigrations, err)
	}

	return nil
}
