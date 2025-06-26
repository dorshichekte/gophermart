package handler

import (
	balance_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/balance"
	order_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/order"
	user_handler "gophermarket/internal/app/adapter/primary/http-adapter/handler/user"
	withdraw_hanlder "gophermarket/internal/app/adapter/primary/http-adapter/handler/withdrawal"
)

type Handlers struct {
	Balance    *balance_handler.Handler
	Order      *order_handler.Handler
	User       *user_handler.Handler
	Withdrawal *withdraw_hanlder.Handler
}
