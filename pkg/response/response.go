package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(data interface{}, msg string, c *gin.Context) {
	Result(http.StatusOK, data, msg, c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(http.StatusOK, struct{}{}, msg, c)
}

func ErrorWithMessage(msg string, err error, c *gin.Context) {
	if err != nil {
		zap.L().Named("[gin]").Error(
			msg,
			zap.Error(err),
		)
	}
	Result(http.StatusInternalServerError, struct{}{}, msg, c)
}

func UnAuthWithMessage(msg string, c *gin.Context) {
	Result(http.StatusUnauthorized, struct{}{}, msg, c)
}

func BadRequestWithMessage(msg string, c *gin.Context) {
	Result(http.StatusBadRequest, struct{}{}, msg, c)
}
