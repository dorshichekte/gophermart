package router

import (
	"net/http"

	"gophermarket/internal/app/adapter/primary/http-adapter/handler"
	"gophermarket/internal/app/adapter/primary/http-adapter/middleware"
	"gophermarket/internal/libs/auth"
)

func (r *Router) AppendRoutes(config Config, handlers *handler.Handlers, auth auth.Auth) {
	r.config = config

	routes := []Route{
		{
			Path:    "/api/user/register",
			Method:  http.MethodPost,
			Handler: http.HandlerFunc(handlers.User.Register(auth)),
		},
		{
			Path:    "/api/user/login",
			Method:  http.MethodPost,
			Handler: http.HandlerFunc(handlers.User.Login(auth)),
		},
		{
			Path:    "/api/user/orders",
			Method:  http.MethodGet,
			Handler: middleware.Add(middleware.Auth(auth))(http.HandlerFunc(handlers.Order.GetOrders)),
		},
		{
			Path:    "/api/user/orders",
			Method:  http.MethodPost,
			Handler: middleware.Add(middleware.Auth(auth))(http.HandlerFunc(handlers.Order.UploadOrder)),
		},
		{
			Path:    "/api/user/balance",
			Method:  http.MethodGet,
			Handler: middleware.Add(middleware.Auth(auth))(http.HandlerFunc(handlers.Balance.GetBalance)),
		},
		{
			Path:    "/api/user/balance/withdraw",
			Method:  http.MethodPost,
			Handler: middleware.Add(middleware.Auth(auth))(http.HandlerFunc(handlers.Withdrawal.Make)),
		},
		{
			Path:    "/api/user/withdrawals",
			Method:  http.MethodGet,
			Handler: middleware.Add(middleware.Auth(auth))(http.HandlerFunc(handlers.Withdrawal.GetWithdrawals)),
		},
	}

	r.appendRoutesToRouter(r.router, routes)
}
