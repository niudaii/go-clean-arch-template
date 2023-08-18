package initdb

import (
	"go-clean-template/internal/entity"

	"gorm.io/gorm"
)

func NewAuthority(db *gorm.DB) Authority {
	return Authority{
		db: db,
	}
}

type Authority struct {
	db *gorm.DB
}

func (a Authority) TableName() string {
	return "system_user_authority"
}

func (a Authority) Initialize() (err error) {
	err = a.db.Where("1 = 1").Delete(&entity.Authority{}).Error
	if err != nil {
		return
	}
	entities := []entity.Authority{
		{
			AuthorityID:   1,
			AuthorityName: "admin",
		},
		{
			AuthorityID:   2,
			AuthorityName: "user",
		},
	}
	err = a.db.Create(&entities).Error
	return
}

func (a Authority) CheckDataExist() bool {
	return false
}
