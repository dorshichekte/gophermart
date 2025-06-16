package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Router struct {
	router *chi.Mux
	config Config
	logger *zap.Logger
}

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}

type Config struct {
	Shutdown time.Duration
	Timeout  time.Duration
}
