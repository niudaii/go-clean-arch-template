package middleware

import (
	"fmt"
	"go-clean-template/pkg/casbin"
	"go-clean-template/pkg/jwt"
	"go-clean-template/pkg/response"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Casbin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户的角色
		sub := fmt.Sprintf("%v", jwt.GetAuthorityID(c)) // 需要转为字符串
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		e := casbin.GetCachedEnforcer(db)
		//判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if gin.Mode() == gin.DebugMode || success {
			c.Next()
		} else {
			response.OkWithMessage("权限不足", c)
			c.Abort()
		}
	}
}
