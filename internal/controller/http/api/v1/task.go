package v1

import (
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/jwt"
	"go-clean-template/pkg/response"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

type TaskApi struct {
	TaskUsecase  entity.TaskUsecase
	AsynqUsecase entity.AsynqUsecase
}

func (a TaskApi) Create(c *gin.Context) {
	var req entity.CreateTask
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequestWithMessage(err.Error(), c)
		return
	}
	userID := jwt.GetUserID(c)
	newTask := entity.Task{
		BasicAccessModel: entity.BasicAccessModel{
			UUID:   uuid.NewV4().String(),
			UserID: userID,
		},
		TaskName: req.TaskName,
		TaskType: req.TaskType,
		Inputs:   req.Inputs,
		Process:  "未开始",
	}
	err = a.TaskUsecase.Create(newTask)
	if err != nil {
		response.ErrorWithMessage("创建失败", err, c)
		return
	}
	err = a.AsynqUsecase.EnqueueTask(newTask)
	if err != nil {
		response.ErrorWithMessage("创建失败", err, c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (a TaskApi) Delete(c *gin.Context) {
	var req entity.DeleteTask
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequestWithMessage(err.Error(), c)
		return
	}
	err = a.TaskUsecase.Delete(req.UUIDs)
	if err != nil {
		response.ErrorWithMessage("删除失败", err, c)
		return
	}
	err = a.AsynqUsecase.DeleteTask(req.UUIDs)
	if err != nil {
		response.ErrorWithMessage("删除失败", err, c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (a TaskApi) FindList(c *gin.Context) {
	var req entity.SearchTask
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequestWithMessage(err.Error(), c)
		return
	}
	list, total, err := a.TaskUsecase.FindList(&req)
	if err != nil {
		response.ErrorWithMessage("查询失败", err, c)
	} else {
		response.Ok(entity.PageResult{
			List:  list,
			Total: total,
		}, "查询成功", c)
	}
}
