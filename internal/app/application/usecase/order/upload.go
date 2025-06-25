package orderusecase

import (
	"context"
	"database/sql"
	"errors"
	"gophermarket/internal/constants"
	customerror "gophermarket/internal/error"

	entity "gophermarket/internal/app/domain/entity/order"
	util "gophermarket/internal/util/order"
)

func (oc *OrderUseCase) Upload(ctx context.Context, userID int, orderNumber string) error {
	isValid := util.IsValidOrderNumber(orderNumber)
	if !isValid {
		return customerror.New(string(constants.ErrInvalidOrderNumber))
	}

	_, err := oc.orderRepository.GetByNumber(ctx, orderNumber)
	if !errors.Is(err, sql.ErrNoRows) {
		return err
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
