package usecase

import (
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/worker"
	"sync"
	"time"
)

var (
	workerStatusMap      = make(map[string]*worker.Status)
	workerStatusMapMutex sync.Mutex
	exitMap              = make(map[string]struct{})
	exitMapMutex         sync.Mutex
)

type workerUsecase struct{}

func NewWorkerUsecase() entity.WorkerUsecase {
	return &workerUsecase{}
}

func (w workerUsecase) CheckExit(uuid string) bool {
	for exitUUID := range exitMap {
		if exitUUID == uuid {
			delete(exitMap, uuid)
			delete(workerStatusMap, uuid)
			return true
		}
	}
	return false
}

func (w workerUsecase) Update(workerStatus *worker.Status) {
	workerStatusMapMutex.Lock()
	defer workerStatusMapMutex.Unlock()
	workerStatusMap[workerStatus.UUID] = workerStatus
}

func (w workerUsecase) AddExit(uuid string) {
	exitMapMutex.Lock()
	defer exitMapMutex.Unlock()
	exitMap[uuid] = struct{}{}
}

func (w workerUsecase) FindList() (list []*worker.Status, total int64, err error) {
	workerStatusMapMutex.Lock()
	defer workerStatusMapMutex.Unlock()
	list = make([]*worker.Status, 0)
	for _, workerStatus := range workerStatusMap {
		if time.Since(workerStatus.UpdatedAt).Minutes() > 3 {
			delete(workerStatusMap, workerStatus.UUID)
			continue
		}
		list = append(list, workerStatus)
	}
	total = int64(len(list))
	return
}
