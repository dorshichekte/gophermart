package order_repository

import (
	"context"

	entity "gophermarket/internal/app/domain/entity/order"
	model "gophermarket/internal/app/repositoriy/model/order"
)

type OrderRepository interface {
	GetAll(ctx context.Context, userID int) ([]model.Order, error)
	GetByNumber(ctx context.Context, orderNumber string) (model.Order, error)
	Upload(ctx context.Context, order entity.Order) error
}
