package entity

import entity "gophermarket/internal/app/domain/entity/order"

type Accrual struct {
	Order   string        `json:"order"`
	Status  entity.Status `json:"status"`
	Accrual float64       `json:"accrual"`
}
