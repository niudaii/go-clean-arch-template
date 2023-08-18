package asynq

import (
	"github.com/hibiken/asynq"
)

var redisClientOpt asynq.RedisClientOpt

func Init(addr, password string) {
	redisClientOpt = asynq.RedisClientOpt{
		Addr:     addr,
		Password: password,
	}
}

func GetRedisClientOpt() asynq.RedisClientOpt {
	return redisClientOpt
}
