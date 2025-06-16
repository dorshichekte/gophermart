package config

import (
	http_adapter "gophermarket/internal/app/adapter/primary/http-adapter"
	"gophermarket/internal/app/adapter/primary/http-adapter/router"
	"gophermarket/internal/app/config/env"
	"gophermarket/internal/constants"
	http_server "gophermarket/internal/libs/http-server"
)

func New() (Config, error) {
	envCfg, err := envconfig.New()
	if err != nil {
		return Config{}, err
	}

	adapterCfg := Adapters{
		HTTPAdapter: http_adapter.Config{
			Router: router.Config{
				Shutdown: constants.DefaultTimeRequest,
				Timeout:  constants.DefaultTimeRequest,
			},
			Server: http_server.Config{
				Address:           envCfg.ServerAddress,
				ReadHeaderTimeout: constants.DefaultTimeRequest,
				WriteTimeout:      constants.DefaultTimeRequest,
				ReadTimeout:       constants.DefaultTimeRequest,
				ShutdownTimeout:   constants.DefaultTimeRequest,
			},
		},
	}

	return Config{
		Env:      envCfg,
		Adapters: adapterCfg,
	}, nil
}
