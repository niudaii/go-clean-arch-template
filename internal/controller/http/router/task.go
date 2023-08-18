package router

import (
	v1 "go-clean-template/internal/controller/http/api/v1"
	"go-clean-template/internal/controller/http/middleware"
	"go-clean-template/internal/repository"
	"go-clean-template/internal/usecase"

	"github.com/hibiken/asynq"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTaskRouter(group *gin.RouterGroup, db *gorm.DB, redisClientOpt asynq.RedisClientOpt) {
	tr := repository.NewTaskRepository(db)
	apiV1 := &v1.TaskApi{
		TaskUsecase:  usecase.NewTaskUsecase(tr),
		AsynqUsecase: usecase.NewAsynqUsecase(redisClientOpt),
	}

	router := group.Group("")
	router.POST("v1/tasks", apiV1.FindList)

	routerWithRecord := group.Group("").Use(middleware.Operation())
	routerWithRecord.POST("v1/task/create", apiV1.Create)
	routerWithRecord.POST("v1/task/delete", apiV1.Delete)
}
