package accrual

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	entity "gophermarket/internal/app/domain/entity/accrual"
	customerror "gophermarket/internal/error"
)

func (ac *AccrualUseCase) Get(orderNumber string) (*entity.Accrual, error) {
	var accrual entity.Accrual
	url := fmt.Sprintf("%s%s%s", ac.config.AccrualAddress, "/api/orders/", orderNumber)
	resp, err := http.Get(url)
	if err != nil {
		return nil, customerror.NewWithData(requestFailed, err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(&accrual); err != nil {
			return nil, customerror.NewWithData("failed to decode response: %w", err)
		}
		return &accrual, nil

	case http.StatusNoContent:
		ac.logger.Info("Order not registered in system", zap.String("orderNumber", orderNumber))
		return nil, nil

	case http.StatusTooManyRequests:
		timeSleep, err := strconv.Atoi(resp.Header.Get("Retry-After"))
		if err != nil {
			time.Sleep(time.Second * 2)
		} else {
			time.Sleep(time.Duration(timeSleep) * time.Second)
		}
		return ac.Get(orderNumber)

	default:
		ac.logger.Error("Accrual server returned unexpected status code", zap.Int("statusCode", resp.StatusCode), zap.String("url", url))
		return nil, customerror.NewWithData("Accrual server error: %s", resp.Status)
	}
}
