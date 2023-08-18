package api

import (
	"context"
	"encoding/json"
	"fmt"
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/task"
	"strings"

	"go.uber.org/zap"

	"github.com/hibiken/asynq"
)

func HandleMath(ctx context.Context, t *asynq.Task) (err error) {
	var p entity.AsynqPayload
	if err = json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	taskID := t.ResultWriter().TaskID()
	zap.L().Sugar().Named("[asynq]").Infof("开始任务 taskId=%v taskType=%v inputs=%v", taskID, t.Type(), p.Inputs)

	inputs := strings.Split(p.Inputs, ",")

	switch p.TaskType {
	case task.TypeAddition:
		err = runAddition(t, inputs)
	case task.TypeMultiplication:
		err = runMultiplication(t, inputs)
	default:
		err = fmt.Errorf("unknown task type: %v", t.Type())
	}
	if err != nil {
		return fmt.Errorf("task failed: %v: %w", err, asynq.SkipRetry)
	}

	return
}
