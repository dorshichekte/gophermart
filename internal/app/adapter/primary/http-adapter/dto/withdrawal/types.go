package dto

import "time"

type WithdrawalResponseDTO struct {
	OrderNumber string    `json:"order"`
	Amount      float64   `json:"sum"`
	CreatedAt   time.Time `json:"processed_at"`
}

type WithdrawalRequestDTO struct {
	Order string  `json:"order" validate:"required,min=3"`
	Sum   float64 `json:"sum" validate:"required,gt=0"`
}
