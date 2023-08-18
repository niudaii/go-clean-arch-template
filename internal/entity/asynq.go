package entity

import "github.com/hibiken/asynq"

type AsynqUsecase interface {
	NewTask(newTask Task) (asynqTask *asynq.Task, err error)
	EnqueueTask(newTask Task) (err error)
	DeleteTask(taskUUIDs []string) (err error)
}

type AsynqPayload struct {
	TaskUUID string `json:"taskUUID"`
	TaskType string `json:"taskType"`
	Inputs   string `json:"inputs"`
}
