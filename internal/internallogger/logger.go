package internallogger

import (
	"github.com/youshy/logger"
	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {
	return logger.NewLogger(logger.DEBUG, false)
}
