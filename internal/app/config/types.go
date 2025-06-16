package config

import (
	http_adapter "gophermarket/internal/app/adapter/primary/http-adapter"
	"gophermarket/internal/app/config/env"
)

//	type Databases struct {
//		Postgres repository.Postgres
//	}
type Adapters struct {
	HttpAdapter http_adapter.Config
	//Databases   Databases
}

type Config struct {
	Env      env_config.Config
	Adapters Adapters
}
