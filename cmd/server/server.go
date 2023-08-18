package main

import (
	serverConfig "go-clean-template/config/server"
	"go-clean-template/internal/app/server"
	"go-clean-template/pkg/config"
	"log"
)

func main() {
	// Initialize config
	file := "server.yaml"
	var conf serverConfig.Server
	err := config.New(file, &conf)
	if err != nil {
		log.Printf("error 配置文件 %v 解析失败: %v\n", file, err)
		return
	}
	serverConfig.ProcessConfig(&conf)
	log.Printf("info 配置文件 %v 解析成功\n%v\n", file, conf.String())
	// Run
	server.Run(&conf)
}
