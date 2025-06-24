package app

import (
	"context"

	"go.uber.org/zap"

	"gophermarket/internal/app/adapter/primary/http-adapter"
	"gophermarket/internal/app/application/usecase"
	"gophermarket/internal/app/config"
	postgres_repository "gophermarket/internal/app/repositoriy/postgres"
	balance_repository_postgres "gophermarket/internal/app/repositoriy/postgres/balance"
	order_repository_postgres "gophermarket/internal/app/repositoriy/postgres/order"
	user_repository_postgres "gophermarket/internal/app/repositoriy/postgres/user"
	withdrawal_repository_postgres "gophermarket/internal/app/repositoriy/postgres/withdrawal"
	"gophermarket/internal/libs/auth"
	"gophermarket/internal/libs/hasher"
	v "gophermarket/internal/libs/validator"
	"gophermarket/internal/libs/worker"
)

func New(l *zap.Logger, cfg config.Config) App {
	validator := v.New()
	h := hasher.New()
	a := auth.New(cfg.Env.AccessSecretKey)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pgDB := postgres_repository.New(l, cfg.Env)

	repos := usecase.Repositories{
		User:       user_repository_postgres.New(pgDB),
		Balance:    balance_repository_postgres.New(pgDB),
		Order:      order_repository_postgres.New(pgDB),
		Withdrawal: withdrawal_repository_postgres.New(pgDB),
	}

	useCases := usecase.New(l, h, a, repos)

	httpAdapter := httpadapter.New(l, a, cfg.Adapters.HTTPAdapter, useCases, validator)

	w := worker.New(cfg.Env, l)
	go w.Start(ctx, useCases.Accrual.PendingOrders)

	return App{
		HTTPAdapter: httpAdapter,
	}
}
