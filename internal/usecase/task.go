package usecase

import "go-clean-template/internal/entity"

type taskUsecase struct {
	taskRepository entity.TaskRepository
}

func NewTaskUsecase(taskRepository entity.TaskRepository) entity.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
	}
}

func (t taskUsecase) Create(task entity.Task) (err error) {
	return t.taskRepository.Create(task)
}

func (t taskUsecase) Delete(uuids []string) (err error) {
	return t.taskRepository.Delete(uuids)
}

func (t taskUsecase) FindList(searchTask *entity.SearchTask) (list []entity.Task, total int64, err error) {
	f := &entity.TaskFilter{}
	f.PageInfo = searchTask.PageInfo
	f.TaskName = searchTask.TaskName
	f.TaskType = searchTask.TaskType
	return t.taskRepository.FindList(f)
}
