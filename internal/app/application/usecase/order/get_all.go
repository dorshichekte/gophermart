package orderusecase

import (
	"context"

	entity "gophermarket/internal/app/domain/entity/order"
)

func (oc *OrderUseCase) GetAll(ctx context.Context, userID int) ([]entity.OrderResponse, error) {
	tmpOrders, err := oc.orderRepository.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	var results []entity.OrderResponse
	for _, tmpOrder := range tmpOrders {
		order := entity.NewOrderResponse(tmpOrder.Number, tmpOrder.Status, tmpOrder.Accrual, tmpOrder.UploadedAt)
		results = append(results, order)
	}

	return results, nil
}
