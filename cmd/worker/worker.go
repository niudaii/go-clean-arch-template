package main

import (
	"flag"
	workerConfig "go-clean-template/config/worker"
	workerEngine "go-clean-template/internal/app/worker"
	"go-clean-template/pkg/config"
	"log"
)

func main() {
	// Initialize config
	file := flag.String("c", "worker.yaml", "config file")
	flag.Parse()
	var conf workerConfig.Worker
	err := config.New(*file, &conf)
	if err != nil {
		log.Printf("error 配置文件 %v 解析失败: %v\n", file, err)
		return
	}
	workerConfig.ProcessConfig(&conf)
	// Run
	workerEngine.Run(&conf)
}
