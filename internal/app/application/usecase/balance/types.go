package balanceusecase

import (
	"go.uber.org/zap"

	balance_repository "gophermarket/internal/app/domain/repository/balance"
)

type BalanceUseCase struct {
	logger            *zap.Logger
	balanceRepository balance_repository.BalanceRepository
}
