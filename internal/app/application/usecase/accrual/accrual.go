package accrual

import "go.uber.org/zap"

func New(logger *zap.Logger) *AccrualUseCase {
	return &AccrualUseCase{logger: logger}
}
