package entity

import "github.com/hibiken/asynq"

type AsynqResult interface {
	WaitForResult(redisClient asynq.RedisClientOpt)
}
