package middleware

import (
	"fmt"

	amqprpc "github.com/0x4b53/amqp-rpc/v3"
	"go.uber.org/zap"
)

func DebugLogger(logger *zap.Logger) amqprpc.LogFunc {
	return func(format string, args ...interface{}) {
		logger.Debug(fmt.Sprintf(format, args...))
	}
}

func ErrorLogger(logger *zap.Logger) amqprpc.LogFunc {
	return func(format string, args ...interface{}) {
		logger.Error(fmt.Sprintf(format, args...))
	}
}
