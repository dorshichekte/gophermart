package orderusecase

import (
	"context"
	"database/sql"
	"errors"

	entity "gophermarket/internal/app/domain/entity/order"
)

func (o *OrderUseCase) Upload(ctx context.Context, userID int, orderNumber string) error {
	_, err := o.orderRepository.GetByNumber(ctx, orderNumber)
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if errors.Is(err, sql.ErrNoRows) {
		ord := entity.NewOrder(orderNumber, userID)
		uploadErr := o.orderRepository.Upload(ctx, ord)
		if uploadErr != nil {
			return uploadErr
		}
	}

	return nil
}
