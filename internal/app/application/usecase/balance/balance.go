package balance_usecase

import (
	"go.uber.org/zap"

	balance_repository "gophermarket/internal/app/domain/repository/balance"
)

func New(logger *zap.Logger, balanceRepository balance_repository.BalanceRepository) *BalanceUseCase {
	return &BalanceUseCase{logger: logger, balanceRepository: balanceRepository}
}
