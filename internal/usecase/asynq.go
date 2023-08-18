package usecase

import (
	"encoding/json"
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/task"
	"time"

	"github.com/hibiken/asynq"
)

type asynqUsecase struct {
	redisClientOpt asynq.RedisClientOpt
}

func NewAsynqUsecase(redisClientOpt asynq.RedisClientOpt) entity.AsynqUsecase {
	return &asynqUsecase{
		redisClientOpt: redisClientOpt,
	}
}

func (a asynqUsecase) EnqueueTask(newTask entity.Task) (err error) {
	var asynqTask *asynq.Task
	asynqTask, err = a.NewTask(newTask)
	if err != nil {
		return
	}
	client := asynq.NewClient(a.redisClientOpt)
	defer client.Close()
	_, err = client.Enqueue(asynqTask, asynq.TaskID(newTask.UUID))
	if err != nil {
		return
	}
	return
}

func (a asynqUsecase) NewTask(newTask entity.Task) (asynqTask *asynq.Task, err error) {
	var payload []byte
	payload, err = json.Marshal(entity.AsynqPayload{
		TaskUUID: newTask.UUID,
		TaskType: newTask.TaskType,
		Inputs:   newTask.Inputs,
	})
	if err != nil {
		return
	}
	asynqTask = asynq.NewTask(task.TypeMath, payload, asynq.Queue(task.TypeMathQueue), asynq.Retention(1*time.Minute), asynq.MaxRetry(1))
	return
}

func (a asynqUsecase) DeleteTask(taskUUIDs []string) (err error) {
	inspector := asynq.NewInspector(a.redisClientOpt)
	for _, id := range taskUUIDs {
		_ = inspector.DeleteTask(task.TypeMathQueue, id)
	}
	return
}
