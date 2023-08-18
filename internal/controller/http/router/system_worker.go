package router

import (
	v1 "go-clean-template/internal/controller/http/api/v1"
	"go-clean-template/internal/controller/http/middleware"
	"go-clean-template/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewSystemWorkerRouter(group *gin.RouterGroup) {
	apiV1 := &v1.SystemWorkerApi{
		WorkerUsecase: usecase.NewWorkerUsecase(),
	}

	router := group.Group("")
	router.POST("v1/system/workers", apiV1.FindList)

	routerWithRecord := group.Group("").Use(middleware.Operation())
	routerWithRecord.POST("v1/system/worker/exit", apiV1.Exit)
}
