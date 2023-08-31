package server

import (
	"fmt"
	"go-clean-template/internal/controller/http/middleware"
	"go-clean-template/internal/controller/http/router"
	"net/http"
	"time"

	"github.com/hibiken/asynq"

	"gorm.io/gorm"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func RunHTTPServer(mode, port string, db *gorm.DB, redisClientOpt asynq.RedisClientOpt) {
	gin.SetMode(mode)
	handler := httpHandler(db, redisClientOpt)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", port),
		Handler:        handler,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	zap.L().Named("[gin]").Error("程序退出", zap.Error(err))
}

func httpHandler(db *gorm.DB, opt asynq.RedisClientOpt) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.GinRecovery(zap.L(), true))
	r.Use(middleware.Cors())

	g := r.Group("api")

	publicGroup := g.Group("")
	{
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
		// base
		router.NewBaseRouter(publicGroup, db)
		// test
		router.NewTestRouter(publicGroup, db)
	}

	privateGroup := g.Group("")
	privateGroup.Use(middleware.JwtAuth()).Use(middleware.Casbin(db))
	{
		// user
		router.NewUserRouter(privateGroup, db)
		// system
		router.NewSystemUserRouter(privateGroup, db)
		router.NewSystemWorkerRouter(privateGroup)
		// task
		router.NewTaskRouter(privateGroup, db, opt)
		// result
		router.NewResultRouter(privateGroup, db)
	}

	return r
}
