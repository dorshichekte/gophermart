package withdrawal_usecase

import (
	"context"

	entity "gophermarket/internal/app/domain/entity/withdrawal"
)

func (wCase *WithdrawalUseCase) Get(ctx context.Context, userID int) ([]entity.Withdrawal, error) {
	wds, err := wCase.withdrawalRepository.Get(ctx, userID)
	if err != nil {
		return []entity.Withdrawal{}, err
	}

	withdrawals := make([]entity.Withdrawal, 0, len(wds))
	for _, w := range wds {
		withdrawals = append(withdrawals, entity.NewWithdrawal(w))
	}

	return withdrawals, nil
}
