package withdrawal_usecase

import (
	"go.uber.org/zap"

	balance_repository "gophermarket/internal/app/domain/repository/balance"
	withdrawal_repository "gophermarket/internal/app/domain/repository/withdrawal"
)

type WithdrawalUseCase struct {
	logger               *zap.Logger
	withdrawalRepository withdrawal_repository.WithdrawalRepository
	balanceRepository    balance_repository.BalanceRepository
}
