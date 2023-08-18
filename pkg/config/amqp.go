package config

import "fmt"

type AMQP struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (a *AMQP) URL() string {
	return fmt.Sprintf("amqp://%v:%v@%v:%v/", a.Username, a.Password, a.Host, a.Port)
}
