package withdrawalusecase

import (
	"context"
	"gophermarket/internal/constants"
	util "gophermarket/internal/util/order"

	customerror "gophermarket/internal/error"
)

func (wCase *WithdrawalUseCase) Make(ctx context.Context, orderNumber string, userID int, sum float64) error {
	isValid := util.IsValidOrderNumber(orderNumber)
	if !isValid {
		return customerror.New(string(constants.ErrInvalidOrderNumber))
	}

	balance, balanceErr := wCase.balanceRepository.Get(ctx, userID)
	if balanceErr != nil {
		return balanceErr
	}

	if balance.Current < sum {
		return customerror.New(errInsufficientBalance)
	}

	makeErr := wCase.withdrawalRepository.Make(ctx, orderNumber, userID, sum)
	if makeErr != nil {
		return makeErr
	}

	return nil
}
