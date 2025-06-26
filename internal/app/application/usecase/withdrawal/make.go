package withdrawalusecase

import (
	"context"
	"gophermarket/internal/constants"

	customerror "gophermarket/internal/error"
	util "gophermarket/internal/util/order"
)

func (wCase *WithdrawalUseCase) Make(ctx context.Context, orderNumber string, userID int, sum float64) error {
	isValid := util.IsValidOrderNumber(orderNumber)
	if !isValid {
		return constants.ErrInvalidOrderNumber
	}

	balance, balanceErr := wCase.balanceRepository.Get(ctx, userID)
	if balanceErr != nil {
		return balanceErr
	}

	if balance.Current < sum {
		return customerror.New(InsufficientBalance)
	}

	makeErr := wCase.withdrawalRepository.Make(ctx, orderNumber, userID, sum)
	if makeErr != nil {
		return makeErr
	}

	return nil
}
