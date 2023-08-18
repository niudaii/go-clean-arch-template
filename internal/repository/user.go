package repository

import (
	"go-clean-template/internal/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u userRepository) buildWhere(f *entity.UserFilter) (db *gorm.DB) {
	db = u.db.Model(&entity.User{})
	if f.Username != "" {
		db = db.Where("username = ?", f.Username)
	}
	if f.ID != 0 {
		db = db.Where("id = ?", f.ID)
	}
	return
}

func (u userRepository) Create(user entity.User) (err error) {
	err = u.db.Create(&user).Error
	return err
}

func (u userRepository) Update(user entity.User) (err error) {
	err = u.db.Save(&user).Error
	return err
}

func (u userRepository) Delete(f *entity.UserFilter) (err error) {
	err = u.buildWhere(f).Delete(&entity.User{}).Error
	return err
}

func (u userRepository) Find(f *entity.UserFilter) (user entity.User, err error) {
	err = u.buildWhere(f).Preload("Authority").First(&user).Error
	return
}

func (u userRepository) FindList(f *entity.UserFilter) (list []entity.User, total int64, err error) {
	db := u.buildWhere(f)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if f.Page > 0 && f.PageSize > 0 {
		limit := f.PageSize
		offset := f.PageSize * (f.Page - 1)
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Preload("Authority").Find(&list).Error
	return
}

func (u userRepository) CheckExist(f *entity.UserFilter) bool {
	var count int64
	u.buildWhere(f).Count(&count)
	return count > 0
}
