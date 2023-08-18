package entity

type Task struct {
	BasicModel
	BasicAccessModel
	TaskName string `json:"taskName"`
	TaskType string `json:"taskType"`
	Inputs   string `json:"inputs"`
	Process  string `json:"process"`
	EndTime  string `json:"endTime"`
}

func (Task) TableName() string {
	return "task"
}

type TaskRepository interface {
	Create(task Task) (err error)
	Delete(uuids []string) (err error)
	Update(task Task) (err error)
	Find(f *TaskFilter) (task Task, err error)
	FindList(f *TaskFilter) (list []Task, total int64, err error)
}

type TaskFilter struct {
	PageInfo
	Task
}

type TaskUsecase interface {
	Create(task Task) (err error)
	Delete(uuids []string) (err error)
	FindList(f *SearchTask) (list []Task, total int64, err error)
}

type CreateTask struct {
	TaskName string `json:"taskName"`
	Inputs   string `json:"inputs"`
	TaskType string `json:"taskType"`
}

type DeleteTask struct {
	UUIDs []string `json:"uuids"`
}

type SearchTask struct {
	PageInfo
	Task
}
