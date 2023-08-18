package initdb

import (
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func NewCasbin(db *gorm.DB) Casbin {
	return Casbin{
		db: db,
	}
}

type Casbin struct {
	db *gorm.DB
}

func (c Casbin) TableName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (c Casbin) Initialize() (err error) {
	err = c.db.Where("1 = 1").Delete(&adapter.CasbinRule{}).Error
	if err != nil {
		return
	}
	entities := []adapter.CasbinRule{
		// user
		{Ptype: "p", V0: "1", V1: "/api/v1/user/getInfo", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/v1/user/logout", V2: "POST"},
		// system - user
		{Ptype: "p", V0: "1", V1: "/api/v1/system/users", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/v1/system/user/delete", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/v1/system/user/create", V2: "POST"},
		// system - worker
		{Ptype: "p", V0: "1", V1: "/api/v1/system/workers", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/v1/system/worker/exit", V2: "POST"},
		// task
		{Ptype: "p", V0: "1", V1: "/api/v1/tasks", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/v1/task/create", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/api/v1/task/delete", V2: "POST"},
		// result
		{Ptype: "p", V0: "1", V1: "/api/v1/results", V2: "POST"},

		// user
		{Ptype: "p", V0: "2", V1: "/api/v1/user/getInfo", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/api/v1/user/logout", V2: "POST"},
		// task
		{Ptype: "p", V0: "2", V1: "/api/v1/tasks", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/api/v1/task/create", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/api/v1/task/delete", V2: "POST"},
		// result
		{Ptype: "p", V0: "2", V1: "/api/v1/results", V2: "POST"},
	}
	err = c.db.Create(&entities).Error
	return
}

func (c Casbin) CheckDataExist() bool {
	return false
}
