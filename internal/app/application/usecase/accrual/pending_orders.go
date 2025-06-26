package accrual

import (
	"context"
	"sync"

	"go.uber.org/zap"

	model "gophermarket/internal/app/repositoriy/model/order"
)

func (ac *AccrualUseCase) PendingOrders(ctx context.Context) {
	orders, err := ac.orderRepository.GetNew(ctx)
	if err != nil {
		ac.logger.Error("Pending orders error fetch orders from DB", zap.Error(err))
		return
	}

	if len(orders) == 0 {
		ac.logger.Info("Orders not found")
		return
	}

	var wg sync.WaitGroup

	jobs := make(chan model.Order, len(orders))

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for order := range jobs {
				resp, err := ac.Get(order.Number)
				if err != nil {
					ac.logger.Error("Pending orders retrieving accrual from accrual service", zap.Error(err))
					continue
				}

				if resp == nil {
					ac.logger.Error("Pending orders skip order without response", zap.Error(err))
					continue
				}

				if order.Status != resp.Status {
					order.Status = resp.Status
					order.Accrual = resp.Accrual
				}

				err = ac.orderRepository.MakeAccrualToBalance(ctx, order)
				if err != nil {
					ac.logger.Error("Pending orders error on updating orders", zap.Error(err))
					continue
				}

				ac.logger.Info("Pending orders update success", zap.String("number", order.Number))
			}
		}()
	}

	for _, order := range orders {
		jobs <- order
	}

	close(jobs)
	wg.Wait()
}
