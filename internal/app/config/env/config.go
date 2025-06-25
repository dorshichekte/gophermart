package envconfig

import (
	"flag"
	"os"

	"github.com/joho/godotenv"

	customerror "gophermarket/internal/error"
)

func New() (Config, error) {
	var cfg Config

	err := cfg.init()
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) valid() error {
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

func (c *Config) initEnv() {
	_ = godotenv.Load()

	c.DatabaseDSN = os.Getenv("DATABASE_URI")
	c.ServerAddress = os.Getenv("RUN_ADDRESS")
	c.AccrualAddress = os.Getenv("ACCRUAL_SYSTEM_ADDRESS")
	c.AccessSecretKey = os.Getenv("ACCESS_SECRET_KEY")
}

func (c *Config) initFlags() {
	flag.StringVar(&c.ServerAddress, "a", c.ServerAddress, "Адрес сервера")
	flag.StringVar(&c.DatabaseDSN, "d", c.DatabaseDSN, "Строка подключения к бд")
	flag.StringVar(&c.AccrualAddress, "r", c.AccrualAddress, "Адрес запуска Accrual system")
	flag.Parse()
}

func (c *Config) initDefaultValues() {
	c.AccessSecretKey = defaultAccessSecret
}

func (c *Config) init() (err error) {
	c.initEnv()
	c.initFlags()
	c.initDefaultValues()

	err = c.valid()

	return err
}
