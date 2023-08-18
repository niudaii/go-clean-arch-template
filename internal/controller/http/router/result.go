package router

import (
	v1 "go-clean-template/internal/controller/http/api/v1"
	"go-clean-template/internal/repository"
	"go-clean-template/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewResultRouter(group *gin.RouterGroup, db *gorm.DB) {
	rr := repository.NewResultRepository(db)
	apiV1 := &v1.ResultApi{
		ResultUsecase: usecase.NewResultUsecase(rr),
	}

	router := group.Group("")
	router.POST("v1/results", apiV1.FindList)
}
