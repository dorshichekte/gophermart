package accrual

import (
	"go.uber.org/zap"

	envconfig "gophermarket/internal/app/config/env"
	order_repository "gophermarket/internal/app/domain/repository/order"
)

func New(logger *zap.Logger, config envconfig.Config, orderRepository order_repository.OrderRepository) *AccrualUseCase {
	return &AccrualUseCase{logger: logger, config: config, orderRepository: orderRepository}
}
