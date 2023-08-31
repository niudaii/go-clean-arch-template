package server

import (
	serverConfig "go-clean-template/config/server"
	"go-clean-template/pkg/asynq"
	"go-clean-template/pkg/db"
	"go-clean-template/pkg/jwt"
	"go-clean-template/pkg/logger"
	"log"
)

func Run(conf *serverConfig.Server) {
	// Logger
	logger.Init(conf.Logger)
	// Jwt
	jwt.Init(conf.Jwt)
	// Asynq
	asynq.Init(conf.Redis.Addr(), conf.Redis.Password)
	// DB
	err := db.Init(conf.DB)
	if err != nil {
		log.Printf("error 连接数据库失败: %v\n", err)
		return
	}
	err = registerTables(db.GetDB())
	if err != nil {
		log.Printf("error 创建表失败: %v\n", err)
		return
	}
	err = initTableData(db.GetDB())
	if err != nil {
		log.Printf("error 初始化表数据失败: %v\n", err)
		return
	}
	// Asynq Client
	go RunAsynqClient(db.GetDB(), asynq.GetRedisClientOpt())
	// RPC Server
	go RunRPCServer(conf.AMQP.URL())
	// HTTP Server
	RunHTTPServer(conf.HTTPServer.Mode, conf.HTTPServer.Port, db.GetDB(), asynq.GetRedisClientOpt())
}
