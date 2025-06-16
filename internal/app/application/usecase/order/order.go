package order_usecase

import (
	"go.uber.org/zap"

	order_repository "gophermarket/internal/app/domain/repository/order"
)

func New(logger *zap.Logger, orderRepository order_repository.OrderRepository) *OrderUseCase {
	return &OrderUseCase{logger: logger, orderRepository: orderRepository}
}
