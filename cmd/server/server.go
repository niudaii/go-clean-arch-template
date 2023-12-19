package main

import (
	"flag"
	serverConfig "go-clean-template/config/server"
	serverEngine "go-clean-template/internal/app/server"
	"go-clean-template/pkg/config"
	"log"
)

func main() {
	// Initialize config
	file := flag.String("c", "server.yaml", "config file")
	flag.Parse()
	var conf serverConfig.Server
	err := config.New(*file, &conf)
	if err != nil {
		log.Printf("error 配置文件 %v 解析失败: %v\n", file, err)
		return
	}
	serverConfig.ProcessConfig(&conf)
	// Run
	serverEngine.Run(&conf)
}
