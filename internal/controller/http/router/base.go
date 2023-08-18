package router

import (
	v1 "go-clean-template/internal/controller/http/api/v1"
	"go-clean-template/internal/repository"
	"go-clean-template/internal/usecase"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func NewBaseRouter(group *gin.RouterGroup, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	apiV1 := &v1.BaseApi{
		LoginUsecase: usecase.NewLoginUsecase(ur),
	}

	router := group.Group("")
	router.POST("v1/base/login", apiV1.Login)
}
