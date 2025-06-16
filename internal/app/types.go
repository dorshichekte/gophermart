package app

import (
	"gophermarket/internal/app/adapter/primary/http-adapter"
)

type App struct {
	HTTPAdapter *httpadapter.HTTPAdapter
}
