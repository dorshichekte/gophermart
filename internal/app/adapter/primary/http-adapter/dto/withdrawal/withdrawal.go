package dto

import entity "gophermarket/internal/app/domain/entity/withdrawal"

func NewWithdrawalResponse(withdrawal entity.Withdrawal) WithdrawalResponseDTO {
	return WithdrawalResponseDTO{
		OrderNumber: withdrawal.OrderNumber,
		Amount:      withdrawal.Amount,
		CreatedAt:   withdrawal.CreatedAt,
	}
}
