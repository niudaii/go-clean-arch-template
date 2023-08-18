package jwt

import uuid "github.com/satori/go.uuid"

var config = &Config{
	SigningKey:  uuid.NewV4().String(),
	ExpiresTime: 604800,
	Issuer:      "niudaii",
	HeaderName:  "Authorization",
}

type Config struct {
	HeaderName  string `json:"header-name" mapstructure:"header-name"`
	SigningKey  string `yaml:"signing-key" mapstructure:"signing-key"`
	ExpiresTime int64  `yaml:"expires-time" mapstructure:"expires-time"`
	Issuer      string `yaml:"issuer"`
}

func Init(conf Config) {
	config = &conf
}

func GetConfig() *Config {
	return config
}
