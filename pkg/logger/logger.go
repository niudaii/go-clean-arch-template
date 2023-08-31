package logger

import (
	"os"

	"github.com/niudaii/util/files"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init(z Zap) {
	if ok := files.PathExists(z.Director); !ok {
		_ = os.Mkdir(z.Director, os.ModePerm)
	}
	cores := z.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if z.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(logger)
}

func GetLogger() *zap.Logger {
	return logger
}
