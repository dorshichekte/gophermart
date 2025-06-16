package dto

import (
	"time"

	"gophermarket/internal/app/domain/entity/order"
)

type OrderResponseDTO struct {
	Number     string        `json:"number"`
	Status     entity.Status `json:"status"`
	Accrual    float64       `json:"accrual,omitempty"`
	UploadedAt time.Time     `json:"uploaded_at"`
}
