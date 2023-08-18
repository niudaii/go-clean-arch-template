package server

import (
	"fmt"
	"go-clean-template/pkg/config"
	"go-clean-template/pkg/db"
	"go-clean-template/pkg/jwt"
	"go-clean-template/pkg/logger"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Root string `yaml:"root"`

	HTTPServer HTTPServer `yaml:"http-server" mapstructure:"http-server"`

	Jwt jwt.Config `yaml:"jwt"`

	Logger logger.Zap `yaml:"logger"`

	DB db.Config `yaml:"db"`

	AMQP  config.AMQP  `yaml:"amqp"`
	Redis config.Redis `yaml:"redis"`
}

type HTTPServer struct {
	Mode string `yaml:"mode"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (s *Server) String() string {
	_, err := yaml.Marshal(*s)
	if err != nil {
		return fmt.Sprintf("%+v", *s)
	}
	var bf []byte
	bf, err = yaml.Marshal(s)
	if err != nil {
		return ""
	}
	return string(bf)
}

func ProcessConfig(conf *Server) {
	// 设置 root
	if conf.Root == "" {
		conf.Root, _ = os.Getwd()
	}
}
