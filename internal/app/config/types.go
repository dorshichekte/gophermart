package config

import (
	"gophermarket/internal/app/config/env"

	http_adapter "gophermarket/internal/app/adapter/primary/http-adapter"
)

type Adapters struct {
	HTTPAdapter http_adapter.Config
}

type Config struct {
	Env      envconfig.Config
	Adapters Adapters
}
