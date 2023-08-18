package v1

import (
	"github.com/gin-gonic/gin"
)

type TestApi struct {
}

func (a TestApi) TestError(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"code":   500,
		"msg":    "xxxx错误",
		"result": "",
	})
}
