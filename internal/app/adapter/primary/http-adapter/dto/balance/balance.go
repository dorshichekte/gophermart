package dto

func NewBalanceResponse(current, withdrawn float64) BalanceResponseDTO {
	return BalanceResponseDTO{
		Current:   current,
		Withdrawn: withdrawn,
	}
}
