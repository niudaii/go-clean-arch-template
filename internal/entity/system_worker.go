package entity

import "go-clean-template/pkg/worker"

type WorkerUsecase interface {
	CheckExit(uuid string) bool
	Update(workerStatus *worker.Status)
	AddExit(uuid string)
	FindList() (list []*worker.Status, total int64, err error)
}

type ExitWorker struct {
	UUID string `json:"uuid"`
}
