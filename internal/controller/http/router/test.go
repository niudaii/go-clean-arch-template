package router

import (
	v1 "go-clean-template/internal/controller/http/api/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTestRouter(group *gin.RouterGroup, db *gorm.DB) {
	apiV1 := &v1.TestApi{}

	router := group.Group("")
	router.POST("v1/test", apiV1.TestError)
}
