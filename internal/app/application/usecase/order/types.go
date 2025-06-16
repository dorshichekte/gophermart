package orderusecase

import (
	"go.uber.org/zap"

	order_repository "gophermarket/internal/app/domain/repository/order"
)

type OrderUseCase struct {
	logger          *zap.Logger
	orderRepository order_repository.OrderRepository
}
