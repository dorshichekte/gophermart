package basehandler

import (
	"go.uber.org/zap"
)

func New(logger *zap.Logger) BaseHandler {
	return BaseHandler{
		logger: logger,
	}
}
