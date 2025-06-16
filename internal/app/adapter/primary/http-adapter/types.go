package httpadapter

import (
	"gophermarket/internal/app/adapter/primary/http-adapter/router"
	"gophermarket/internal/libs/http-server"
)

type Config struct {
	Server httpserver.Config
	Router router.Config
}

type HTTPAdapter struct {
	server *httpserver.Server
}
