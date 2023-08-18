package v1

import (
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/response"

	"github.com/gin-gonic/gin"
)

type SystemUserApi struct {
	SystemUserUsecase entity.SystemUserUsecase
}

func (a SystemUserApi) Create(c *gin.Context) {
	var req entity.CreateUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequestWithMessage(err.Error(), c)
		return
	}
	err = a.SystemUserUsecase.Create(req)
	if err != nil {
		response.ErrorWithMessage("创建失败", err, c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (a SystemUserApi) Delete(c *gin.Context) {
	var req entity.DeleteUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequestWithMessage(err.Error(), c)
		return
	}
	err = a.SystemUserUsecase.Delete(req.ID)
	if err != nil {
		response.ErrorWithMessage("删除失败", err, c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (a SystemUserApi) FindList(c *gin.Context) {
	var req entity.PageInfo
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequestWithMessage(err.Error(), c)
		return
	}
	list, total, err := a.SystemUserUsecase.FindList(req)
	if err != nil {
		response.ErrorWithMessage("查询失败", err, c)
	} else {
		response.Ok(entity.PageResult{
			List:  list,
			Total: total,
		}, "查询成功", c)
	}
}
