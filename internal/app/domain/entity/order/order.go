package entity

import (
	"time"
)

func NewOrder(number string, userID int) Order {
	return Order{
		Number: number,
		Status: StatusNew,
		UserID: userID,
	}
}

func NewOrderResponse(number string, status Status, accrual float64, uploadedAt time.Time) OrderResponse {
	return OrderResponse{
		Number:     number,
		Status:     status,
		Accrual:    accrual,
		UploadedAt: uploadedAt,
	}
}
