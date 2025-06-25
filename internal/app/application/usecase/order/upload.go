package orderusecase

import (
	"context"
	"database/sql"
	"errors"

	entity "gophermarket/internal/app/domain/entity/order"
	"gophermarket/internal/constants"
	customerror "gophermarket/internal/error"
	util "gophermarket/internal/util/order"
)

func (oc *OrderUseCase) Upload(ctx context.Context, userID int, orderNumber string) error {
	isValid := util.IsValidOrderNumber(orderNumber)
	if !isValid {
		return customerror.New(string(constants.ErrInvalidOrderNumber))
	}

	order, err := oc.orderRepository.GetByNumber(ctx, orderNumber)
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if err == nil && order.UserID == userID {
		return ErrOrderExists
	}

	if err == nil && order.UserID != userID {
		return ErrOrderExistsByAnotherUser
	}

	if errors.Is(err, sql.ErrNoRows) {
		ord := entity.NewOrder(orderNumber, userID)
		uploadErr := oc.orderRepository.Upload(ctx, ord)

		if uploadErr != nil {
			return uploadErr
		}
	}

	return nil
}
