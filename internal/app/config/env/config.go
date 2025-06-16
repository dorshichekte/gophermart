package env_config

import (
	"flag"
	customerror "gophermarket/internal/error"
	"os"

	"github.com/joho/godotenv"
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
	//if c.AccrualAddress == "" {
	//	missingVariables = append(missingVariables, "ACCRUAL_SYSTEM_ADDRESS")
	//}
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
	flag.StringVar(&c.ServerAddress, "a", c.ServerAddress, "Адресс сервера")
	flag.StringVar(&c.DatabaseDSN, "d", c.DatabaseDSN, "Строка подключения к бд")
	flag.StringVar(&c.AccrualAddress, "r", c.AccrualAddress, "Адресс запуска Accrual system")
	flag.Parse()
}

func (c *Config) initDefaultValues() {
	c.AccessSecretKey = defaultAccessTokenSecret
}

func (c *Config) init() (err error) {
	c.initEnv()
	c.initFlags()

	err = c.valid()
	return err
}
