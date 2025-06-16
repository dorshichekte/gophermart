package withdrawal_usecase

import (
	"context"

	customerror "gophermarket/internal/error"
)

func (wCase *WithdrawalUseCase) Make(ctx context.Context, orderNum string, userID int, sum float64) error {
	balance, balanceErr := wCase.balanceRepository.Get(ctx, userID)
	if balanceErr != nil {
		return balanceErr
	}

	if balance.Current < sum {
		return customerror.New(errInsufficientBalance)
	}

	makeErr := wCase.withdrawalRepository.Make(ctx, orderNum, userID, sum)
	if makeErr != nil {
		return makeErr
	}

	return nil
}
