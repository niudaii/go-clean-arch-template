package usecase

import (
	"encoding/json"
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/task"
	"time"

	uuid "github.com/satori/go.uuid"

	"go.uber.org/zap"

	"github.com/hibiken/asynq"
)

type asynqResultUsecase struct {
	resultRepository entity.ResultRepository
	taskRepository   entity.TaskRepository
}

func NewAsynqResultUsecase(resultRepository entity.ResultRepository, taskRepository entity.TaskRepository) entity.AsynqResult {
	return &asynqResultUsecase{
		resultRepository: resultRepository,
		taskRepository:   taskRepository,
	}
}

func (a asynqResultUsecase) WaitForResult(redisClientOpt asynq.RedisClientOpt) {
	inspector := asynq.NewInspector(redisClientOpt)
	for {
		queues, err := inspector.Queues()
		if err != nil {
			zap.L().Sugar().Named("[result]").Error("inspector.Queues() err", zap.Error(err))
			continue
		}
		for _, queue := range queues {
			var taskInfoList []*asynq.TaskInfo
			taskInfoList, err = inspector.ListCompletedTasks(queue)
			if err != nil {
				zap.L().Sugar().Named("[result]").Error("inspector.ListCompletedTasks() err", zap.Error(err))
				continue
			}
			for _, taskInfo := range taskInfoList {
				err = a.handleResult(taskInfo)
				if err != nil {
					zap.L().Named("[result]").Error("任务结果处理错误", zap.Error(err))
				}
				err = inspector.DeleteTask(queue, taskInfo.ID)
				if err != nil {
					zap.L().Sugar().Named("[result]").Error("inspector.DeleteTask() err", zap.Error(err))
				}
			}
		}
		// 等待10秒
		time.Sleep(10 * time.Second)
	}
}

func (a asynqResultUsecase) handleResult(taskInfo *asynq.TaskInfo) (err error) {
	var p entity.AsynqPayload
	err = json.Unmarshal(taskInfo.Payload, &p)
	if err != nil {
		return
	}
	zap.L().Sugar().Named("[result]").Infof("任务完成: taskId=%v, taskType=%v inputs=%v", taskInfo.ID, p.TaskType, p.Inputs)

	f := &entity.TaskFilter{}
	f.UUID = taskInfo.ID
	t, err := a.taskRepository.Find(f)
	if err != nil {
		return
	}
	t.Process = "完成"
	t.EndTime = time.Now().Format("2006-01-02 15:04:05")
	err = a.taskRepository.Update(t)
	if err != nil {
		return
	}

	switch p.TaskType {
	case task.TypeAddition:
		err = a.handleAddition(taskInfo.ID, taskInfo.Result)
	case task.TypeMultiplication:
		err = a.handleMultiplication(taskInfo.ID, taskInfo.Result)
	}

	return
}

func (a asynqResultUsecase) handleAddition(taskID string, resultBytes []byte) (err error) {
	var result entity.Result
	err = json.Unmarshal(resultBytes, &result)
	if err != nil {
		return
	}
	zap.L().Info("[addition]", zap.Any("result", result.Result))

	result.UUID = uuid.NewV4().String()
	result.TaskUUID = taskID
	err = a.resultRepository.Create(result)

	return
}

func (a asynqResultUsecase) handleMultiplication(taskID string, resultBytes []byte) (err error) {
	var result entity.Result
	err = json.Unmarshal(resultBytes, &result)
	if err != nil {
		return
	}
	zap.L().Info("[multiplication]", zap.Any("result", result.Result))

	result.UUID = uuid.NewV4().String()
	result.TaskUUID = taskID
	err = a.resultRepository.Create(result)

	return
}
