package config

import (
	"flag"
	"os"

	customerror "gophermarket/internal/error"
)

func NewEnvConfig() (Env, error) {
	var cfg Env

	err := cfg.init()
	if err != nil {
		return Env{}, err
	}

	return cfg, nil
}

func (c *Env) valid() error {
	var missingVariables []string

	if c.DatabaseDSN == "" {
		missingVariables = append(missingVariables, "DATABASE_URI")
	}
	if c.ServerAddress == "" {
		missingVariables = append(missingVariables, "RUN_ADDRESS")
	}
	if c.AccessSecretKey == "" {
		missingVariables = append(missingVariables, "ACCESS_SECRET_KEY")
	}

	hasEmptyVariables := len(missingVariables) != 0
	if hasEmptyVariables {
		return customerror.NewWithData(errEnvMissingVariables, missingVariables)
	}

	return nil
}

func (c *Env) initEnv() {
	c.DatabaseDSN = os.Getenv("DATABASE_URI")
	c.ServerAddress = os.Getenv("RUN_ADDRESS")
	c.AccrualAddress = os.Getenv("ACCRUAL_SYSTEM_ADDRESS")
	c.AccessSecretKey = os.Getenv("ACCESS_SECRET_KEY")
}

func (c *Env) initFlags() {
	flag.StringVar(&c.ServerAddress, "a", c.ServerAddress, "Адрес сервера")
	flag.StringVar(&c.DatabaseDSN, "d", c.DatabaseDSN, "Строка подключения к бд")
	flag.StringVar(&c.AccrualAddress, "r", c.AccrualAddress, "Адрес запуска Accrual system")
	flag.Parse()
}

func (c *Env) initDefaultValues() {
	c.AccessSecretKey = defaultAccessSecret
}

func (c *Env) init() (err error) {
	c.initEnv()
	c.initFlags()
	c.initDefaultValues()

	err = c.valid()

	return err
}
