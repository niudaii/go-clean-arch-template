package middleware

import (
	"context"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func ErrorLogger(logger *zap.Logger) asynq.ErrorHandlerFunc {
	return func(ctx context.Context, task *asynq.Task, err error) {
		logger.Named("[asynq]").Error("error", zap.Error(err))
	}
}

type Logger struct {
	zapLogger *zap.Logger
}

func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{
		zapLogger: logger,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.zapLogger.Sugar().Debug(v)
	return
}

func (l *Logger) Info(v ...interface{}) {
	l.zapLogger.Sugar().Info(v)
	return
}

func (l *Logger) Warn(v ...interface{}) {
	l.zapLogger.Sugar().Warn(v)
	return
}

func (l *Logger) Error(v ...interface{}) {
	l.zapLogger.Sugar().Error(v)
	return
}

func (l *Logger) Fatal(v ...interface{}) {
	l.zapLogger.Sugar().Fatal(v)
	return
}
