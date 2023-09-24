package worker

import (
	workerConfig "go-clean-template/config/worker"
	"go-clean-template/pkg/logger"
	"go-clean-template/pkg/worker"
	"log"
)

func Run(conf *workerConfig.Worker) {
	// Logger
	logger.Init(conf.Logger)
	// Init Worker
	err := worker.Init()
	if err != nil {
		log.Printf("error 初始化 worker 失败: %v\n", err)
		return
	}
	// RPC Client
	go RunRPCClient(conf.AMQP.URL())
	// Asynq Server
	RunAsynqServer(conf.Redis.Addr(), conf.Redis.Password)
}
