package repository

import (
	"go-clean-template/internal/entity"

	"gorm.io/gorm"
)

type resultRepository struct {
	db *gorm.DB
}

func NewResultRepository(db *gorm.DB) entity.ResultRepository {
	return &resultRepository{
		db: db,
	}
}

func (r resultRepository) buildWhere(f *entity.ResultFilter) (db *gorm.DB) {
	db = r.db.Model(&entity.Result{})
	if f.TaskUUID != "" {
		db = db.Where("task_uuid = ?", f.TaskUUID)
	}
	return
}
func (r resultRepository) Create(result entity.Result) (err error) {
	err = r.db.Create(&result).Error
	return
}

func (r resultRepository) FindList(f *entity.ResultFilter) (list []entity.Result, total int64, err error) {
	db := r.buildWhere(f)
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
