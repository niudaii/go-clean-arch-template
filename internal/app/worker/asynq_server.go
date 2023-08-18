package worker

import (
	"go-clean-template/internal/controller/asynq/middleware"
	"go-clean-template/internal/controller/asynq/router"
	"go-clean-template/pkg/task"
	"time"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func RunAsynqServer(addr, password string) {
	logger := zap.L().Named("[asynq]")
	newLogger := middleware.NewLogger(logger)
	// 连接消息中间件
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     addr,
			Password: password,
		},
		asynq.Config{
			ErrorHandler: middleware.ErrorLogger(logger),
			Logger:       newLogger,
			Queues: map[string]int{
				task.TypeMathQueue: 1,
			},
			StrictPriority: true,
			Concurrency:    1,
			RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
				return 10 * time.Second
			},
		},
	)
	mux := asynqHandler()
	err := srv.Run(mux)
	logger.Error("程序退出", zap.Error(err))
}

func asynqHandler() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	mux.Use(middleware.Process)

	router.NewMathRouter(mux)

	return mux
}
