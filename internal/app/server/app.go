package server

import (
	serverConfig "go-clean-template/config/server"
	"go-clean-template/pkg/asynq"
	"go-clean-template/pkg/db"
	"go-clean-template/pkg/jwt"
	"go-clean-template/pkg/logger"
	"go.uber.org/zap"
)

const (
	InitConfigSuccess = "初始化配置成功:\n%v"
	ConnDBFail        = "连接数据库失败: %v"
	CreateTableFail   = "创建表失败: %v"
	InitDataFail      = "初始化表数据失败: %v"
)

func Run(conf *serverConfig.Server) {
	// Init Logger
	logger.Init(conf.Logger)
	zap.L().Sugar().Infof(InitConfigSuccess, conf.String())
	// Init Jwt
	jwt.Init(conf.Jwt)
	// Init AsynqConfig
	asynq.Init(conf.Redis.Addr(), conf.Redis.Password)
	// Init DB
	err := db.Init(conf.DB)
	if err != nil {
		zap.L().Sugar().Errorf(ConnDBFail, err)
		return
	}
	err = registerTables(db.GetDB())
	if err != nil {
		zap.L().Sugar().Errorf(CreateTableFail, err)
		return
	}
	err = initTableData(db.GetDB())
	if err != nil {
		zap.L().Sugar().Errorf(InitDataFail, err)
		return
	}

	// Run Asynq Client
	go RunAsynqClient(db.GetDB(), asynq.GetRedisClientOpt())
	// Run RPC Server
	go RunRPCServer(conf.AMQP.URL())
	// Run HTTP Server
	RunHTTPServer(conf.HTTPServer.Mode, conf.HTTPServer.Port, db.GetDB(), asynq.GetRedisClientOpt())
}
