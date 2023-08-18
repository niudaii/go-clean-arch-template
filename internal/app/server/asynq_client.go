package server

import (
	"go-clean-template/internal/repository"
	"go-clean-template/internal/usecase"

	"gorm.io/gorm"

	"github.com/hibiken/asynq"
)

func RunAsynqClient(db *gorm.DB, redisClientOpt asynq.RedisClientOpt) {
	rr := repository.NewResultRepository(db)
	tr := repository.NewTaskRepository(db)
	aru := usecase.NewAsynqResultUsecase(rr, tr)

	go aru.WaitForResult(redisClientOpt) // 接收任务结果
}
