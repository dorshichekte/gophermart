package orderrepository

import (
	"context"

	entity "gophermarket/internal/app/domain/entity/order"
	model "gophermarket/internal/app/repositoriy/model/order"
)

//go:generate mockgen -package mocks -source order.go -destination ../../mock/order_repository.go OrderRepository
type OrderRepository interface {
	GetAll(ctx context.Context, userID int) ([]model.Order, error)
	GetByNumber(ctx context.Context, orderNumber string) (model.Order, error)
	Upload(ctx context.Context, order entity.Order) error
	GetNew(ctx context.Context) ([]model.Order, error)
	MakeAccrualToBalance(ctx context.Context, order model.Order) error
}
