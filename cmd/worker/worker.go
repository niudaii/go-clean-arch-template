package main

import (
	workerConfig "go-clean-template/config/worker"
	"go-clean-template/internal/app/worker"
	"go-clean-template/pkg/config"
	"log"
)

func main() {
	// Initialize config
	file := "worker.yaml"
	var conf workerConfig.Worker
	err := config.New(file, &conf)
	if err != nil {
		log.Printf("error 配置文件 %v 解析失败: %v\n", file, err)
		return
	}
	workerConfig.ProcessConfig(&conf)
	log.Printf("info 配置文件 %v 解析成功\n%v\n", file, conf.String())
	// Run
	worker.Run(&conf)
}
