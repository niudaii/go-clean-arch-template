package worker

import (
	workerConfig "go-clean-template/config/worker"
	"go-clean-template/pkg/logger"
	"go-clean-template/pkg/worker"
	"go.uber.org/zap"
)

const (
	InitConfigSuccess = "初始化配置成功:\n%v"
	InitWorkerFail    = "初始化 worker 失败: %v"
)

func Run(conf *workerConfig.Worker) {
	// Init Logger
	logger.Init(conf.Logger)
	zap.L().Sugar().Infof(InitConfigSuccess, conf.String())
	// Init Worker
	err := worker.Init()
	if err != nil {
		zap.L().Sugar().Errorf(InitWorkerFail, err)
		return
	}

	// Run RPC Client
	go RunRPCClient(conf.AMQP.URL())
	// Run Asynq Server
	RunAsynqServer(conf.Redis.Addr(), conf.Redis.Password)
}
