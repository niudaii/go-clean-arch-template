package repository

import (
	"go-clean-template/internal/entity"
	"go-clean-template/internal/repository/initdb"
	"go-clean-template/pkg/db"

	"gorm.io/gorm"
)

type initDBRepository struct {
	db *gorm.DB
}

func NewSourceRepository(db *gorm.DB) entity.SourceRepository {
	return &initDBRepository{
		db: db,
	}
}

func (i initDBRepository) InitTableData() (err error) {
	err = db.InitTableData(
		initdb.NewUser(i.db),
		initdb.NewAuthority(i.db),
		initdb.NewCasbin(i.db),
	)
	return
}
