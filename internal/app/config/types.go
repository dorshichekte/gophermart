package config

import (
	adapter "gophermarket/internal/app/config/adapter"
	"gophermarket/internal/app/config/env"
)

type Config struct {
	Env         config.Env
	HTTPAdapter adapter.HTTPAdapter
}
