package server

import (
	"go-clean-template/internal/entity"
	"go-clean-template/internal/repository"
	"go-clean-template/internal/usecase"

	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func registerTables(db *gorm.DB) (err error) {
	err = db.AutoMigrate(
		// system
		adapter.CasbinRule{},
		entity.User{},
		entity.Authority{},
		entity.Menu{},
		// task
		entity.Task{},
		// result
		entity.Result{},
	)
	return
}

func initTableData(db *gorm.DB) (err error) {
	sr := repository.NewSourceRepository(db)
	err = usecase.NewInitdbUsecase(sr).InitTableData()
	return
}
