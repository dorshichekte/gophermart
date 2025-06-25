package accrual

import (
	"go.uber.org/zap"

	envconfig "gophermarket/internal/app/config/env"
	order_repository "gophermarket/internal/app/domain/repository/order"
)

type AccrualUseCase struct {
	logger          *zap.Logger
	config          envconfig.Config
	orderRepository order_repository.OrderRepository
}
