package app

import (
	"gophermarket/internal/app/adapter/primary/http-adapter"
)

type App struct {
	HttpAdapter *http_adapter.HttpAdapter
}
