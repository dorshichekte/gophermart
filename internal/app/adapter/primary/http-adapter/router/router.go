package router

import (
	"net/http"

	"gophermarket/internal/app/adapter/primary/http-adapter/middleware"

	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
)

func New(logger *zap.Logger) *Router {
	router := chi.NewRouter()

	r := Router{
		router: router,
		logger: logger,
	}

	return &r
}

func (r *Router) Router() http.Handler {
	return r.router
}

func (r *Router) appendRoutesToRouter(subrouter *chi.Mux, routes []Route) {
	globalMiddlewares := chi.Middlewares{middleware.Log(r.logger), middleware.Gzip, middleware.Decompress}

	subrouter.Use(globalMiddlewares...)

	for _, route := range routes {
		subrouter.Method(route.Method, route.Path, route.Handler)
	}
}
