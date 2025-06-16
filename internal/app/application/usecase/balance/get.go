package balance_usecase

import (
	"context"

	entity "gophermarket/internal/app/domain/entity/balance"
)

func (b *BalanceUseCase) Get(ctx context.Context, userID int) (entity.Balance, error) {
	balance, err := b.balanceRepository.Get(ctx, userID)
	if err != nil {
		return entity.Balance{}, err
	}

	entityBalance := entity.NewBalance(balance.Current, balance.Withdrawn)
	return entityBalance, nil
}
