package worker

import (
	"fmt"
	"go-clean-template/pkg/config"
	"go-clean-template/pkg/logger"
	"os"

	"gopkg.in/yaml.v3"
)

type Worker struct {
	Root string `yaml:"root"`

	Logger logger.Zap `yaml:"logger"`

	AMQP  config.AMQP  `yaml:"amqp"`
	Redis config.Redis `yaml:"redis"`
}

func (w *Worker) String() string {
	_, err := yaml.Marshal(*w)
	if err != nil {
		return fmt.Sprintf("%+v", *w)
	}
	var bf []byte
	bf, err = yaml.Marshal(w)
	if err != nil {
		return ""
	}
	return string(bf)
}

func ProcessConfig(conf *Worker) {
	// 设置 root
	if conf.Root == "" {
		conf.Root, _ = os.Getwd()
	}
}
