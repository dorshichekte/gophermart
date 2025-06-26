package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	adapter "gophermarket/internal/app/config/adapter"
)

type Router struct {
	router *chi.Mux
	config adapter.Router
	logger *zap.Logger
}

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}
