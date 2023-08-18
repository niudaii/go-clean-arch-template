package worker

import (
	"os"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/niudaii/util"
)

var (
	status *Status
)

type Status struct {
	mu                 sync.Mutex
	UUID               string    `json:"uuid"`
	IP                 string    `json:"ip"`
	Hostname           string    `json:"hostname"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	TaskExecutedNumber int       `json:"taskExecutedNumber"` // 任务执行数
	IsRunning          bool      `json:"isRunning"`          // 正在执行
}

func Init() (err error) {
	ip := util.GetOutBoundIP()
	if ip == "" {
		ip, err = util.GetClientIP()
		if err != nil {
			return
		}
	}
	var hostname string
	hostname, err = os.Hostname()
	if err != nil {
		return
	}
	status = &Status{
		mu:        sync.Mutex{},
		UUID:      uuid.NewV4().String(),
		IP:        ip,
		Hostname:  hostname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return
}

func GetStatus() *Status {
	return status
}

func AddTaskExecutedNumber() {
	status.mu.Lock()
	defer status.mu.Unlock()
	status.TaskExecutedNumber++
}
