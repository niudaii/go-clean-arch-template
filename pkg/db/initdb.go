package db

import (
	"log"
)

const (
	InitDataExist   = "%v 表的初始数据已存在\n"
	InitDataFailed  = "%v 表初始化数据失败: %v\n"
	InitDataSuccess = "%v 表初始化数据成功\n"
)

type InitData interface {
	TableName() string
	Initialize() (err error)
	CheckDataExist() bool
}

func InitTableData(inits ...InitData) error {
	for i := 0; i < len(inits); i++ {
		if inits[i].CheckDataExist() {
			log.Printf(InitDataExist, inits[i].TableName())
			continue
		}
		if err := inits[i].Initialize(); err != nil {
			log.Printf(InitDataFailed, inits[i].TableName(), err)
		}
		log.Printf(InitDataSuccess, inits[i].TableName())
	}
	return nil
}
