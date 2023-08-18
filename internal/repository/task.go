package repository

import (
	"go-clean-template/internal/entity"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) entity.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (t taskRepository) buildWhere(f *entity.TaskFilter) (db *gorm.DB) {
	db = t.db.Model(&entity.Task{})
	if f.UUID != "" {
		db = db.Where("uuid = ?", f.UUID)
	}
	if f.TaskName != "" {
		db = db.Where("task_name = ?", f.TaskName)
	}
	if f.TaskType != "" {
		db = db.Where("task_type = ?", f.TaskType)
	}
	return
}

func (t taskRepository) Create(task entity.Task) (err error) {
	err = t.db.Create(&task).Error
	return
}

func (t taskRepository) Delete(uuids []string) (err error) {
	err = t.db.Where("uuid in (?)", uuids).Delete(&entity.Task{}).Error
	return
}

func (t taskRepository) Update(task entity.Task) (err error) {
	err = t.db.Save(&task).Error
	return
}

func (t taskRepository) Find(f *entity.TaskFilter) (task entity.Task, err error) {
	err = t.buildWhere(f).First(&task).Error
	return
}

func (t taskRepository) FindList(f *entity.TaskFilter) (list []entity.Task, total int64, err error) {
	db := t.buildWhere(f)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if f.Page > 0 && f.PageSize > 0 {
		limit := f.PageSize
		offset := f.PageSize * (f.Page - 1)
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("id desc").Find(&list).Error
	return
}
