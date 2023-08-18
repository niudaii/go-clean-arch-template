package router

import (
	v1 "go-clean-template/internal/controller/http/api/v1"
	"go-clean-template/internal/repository"
	"go-clean-template/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRouter(group *gin.RouterGroup, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	apiV1 := &v1.UserApi{
		UserUsecase: usecase.NewUserUsecase(ur),
	}

	router := group.Group("")
	router.POST("v1/user/getInfo", apiV1.GetInfo)
	router.POST("v1/user/logout", apiV1.Logout)
}
