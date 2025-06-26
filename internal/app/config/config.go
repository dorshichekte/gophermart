package config

import (
	"fmt"

	"github.com/joho/godotenv"

	adapter "gophermarket/internal/app/config/adapter"
	"gophermarket/internal/app/config/env"
)

func New() (Config, error) {
	err := godotenv.Load()

	envCfg, err := config.NewEnvConfig()
	if err != nil {
		return Config{}, err
	}

	httpAdapterCfg := adapter.NewAdapterConfig(envCfg.ServerAddress)

	fmt.Println(envCfg, httpAdapterCfg)

	return Config{
		Env:         envCfg,
		HTTPAdapter: httpAdapterCfg,
	}, nil
}
