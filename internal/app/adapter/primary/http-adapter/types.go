package http_adapter

import (
	"gophermarket/internal/app/adapter/primary/http-adapter/router"
	"gophermarket/internal/libs/http-server"
)

type Config struct {
	Server http_server.Config
	Router router.Config
}

type HttpAdapter struct {
	server *http_server.Server
}
