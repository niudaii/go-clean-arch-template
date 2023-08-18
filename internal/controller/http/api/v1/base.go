package v1

import (
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/response"

	"github.com/gin-gonic/gin"
)

type BaseApi struct {
	LoginUsecase entity.LoginUsecase
}

func (a BaseApi) Login(c *gin.Context) {
	var req entity.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequestWithMessage(err.Error(), c)
		return
	}
	// 判断用户是否存在
	user, err := a.LoginUsecase.GetUserByUsername(req.Username)
	if err != nil {
		response.ErrorWithMessage("用户名或密码错误", nil, c)
		return
	}
	// 密码校验
	if user.Password != req.Password {
		response.ErrorWithMessage("用户名或密码错误", nil, c)
		return
	}
	// 登录成功,返回 token
	token, err := a.LoginUsecase.CreateAccessToken(&user)
	if err != nil {
		response.ErrorWithMessage("生成 token 失败", err, c)
	} else {
		loginResponse := entity.LoginResponse{
			Token: token,
		}
		response.Ok(loginResponse, "登录成功", c)
	}
}
