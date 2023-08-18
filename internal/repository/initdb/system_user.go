package initdb

import (
	"go-clean-template/internal/entity"

	"gorm.io/gorm"
)

func NewUser(db *gorm.DB) User {
	return User{
		db: db,
	}
}

type User struct {
	db *gorm.DB
}

func (User) TableName() string {
	return "system_user"
}

func (u User) Initialize() (err error) {
	entities := []entity.User{
		{
			Username:    "admin",
			Password:    "e10adc3949ba59abbe56e057f20f883e",
			AuthorityID: 1,
		},
		{
			Username:    "user",
			Password:    "e10adc3949ba59abbe56e057f20f883e",
			AuthorityID: 2,
		},
	}
	err = u.db.Create(&entities).Error
	return
}

func (u User) CheckDataExist() bool {
	var total int64
	u.db.Model(&entity.User{}).Count(&total)
	return total > 0
}
