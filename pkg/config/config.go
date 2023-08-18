package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func New(filename string, val interface{}) (err error) {
	v := viper.New()
	v.SetConfigFile(filename)
	v.SetConfigType("yaml")
	err = v.ReadInConfig()
	if err != nil {
		return
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("info %v 配置改动\n", e.Name)
		if err = v.Unmarshal(&val); err != nil {
			return
		}
	})
	v.WatchConfig()
	err = v.Unmarshal(&val)
	return
}
