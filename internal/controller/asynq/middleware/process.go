package middleware

import (
	"context"
	"go-clean-template/pkg/worker"
	"time"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func Process(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		// 任务开始
		start := time.Now()
		worker.AddTaskExecutedNumber()
		err := h.ProcessTask(ctx, t)
		if err != nil {
			return err
		}
		// 任务结束
		taskID := t.ResultWriter().TaskID()
		zap.L().Sugar().Named("[asynq]").Infof("任务结束 taskId=%v spentTime=%v\n", taskID, time.Since(start))
		return nil
	})
}
