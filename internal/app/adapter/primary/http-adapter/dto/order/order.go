package dto

import (
	"time"

	entity "gophermarket/internal/app/domain/entity/order"
)

func NewOrderResponse(number string, status entity.Status, accrual float64, uploadedAt time.Time) OrderResponseDTO {
	return OrderResponseDTO{
		Number:     number,
		Status:     status,
		Accrual:    accrual,
		UploadedAt: uploadedAt,
	}
}
