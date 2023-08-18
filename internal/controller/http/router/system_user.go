package router

import (
	v1 "go-clean-template/internal/controller/http/api/v1"
	"go-clean-template/internal/controller/http/middleware"
	"go-clean-template/internal/repository"
	"go-clean-template/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSystemUserRouter(group *gin.RouterGroup, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	apiV1 := &v1.SystemUserApi{
		SystemUserUsecase: usecase.NewSystemUserUsecase(ur),
	}

	router := group.Group("")
	router.POST("v1/system/users", apiV1.FindList)

	routerWithRecord := group.Group("").Use(middleware.Operation())
	routerWithRecord.POST("v1/system/user/create", apiV1.Create)
	routerWithRecord.POST("v1/system/user/delete", apiV1.Delete)
}
