package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/jwt"
	"go-clean-template/pkg/response"
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
			UUID:   uuid.Must(uuid.NewV4()).String(),
			UserID: userID,
		},
		TaskName: req.TaskName,
		TaskType: req.TaskType,
		Inputs:   req.Inputs,
		Process:  "未开始",
	}
	err = a.TaskUsecase.Create(newTask)
	if err != nil {
		response.ErrorWithMessage(CreateFail, err, c)
		return
	}
	err = a.AsynqUsecase.EnqueueTask(newTask)
	if err != nil {
		response.ErrorWithMessage(CreateFail, err, c)
	} else {
		response.OkWithMessage(CreateSuccess, c)
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
		response.ErrorWithMessage(DeleteFail, err, c)
		return
	}
	err = a.AsynqUsecase.DeleteTask(req.UUIDs)
	if err != nil {
		response.ErrorWithMessage(DeleteFail, err, c)
	} else {
		response.OkWithMessage(DeleteSuccess, c)
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
		response.ErrorWithMessage(FindFail, err, c)
	} else {
		response.Ok(entity.PageResult{
			List:  list,
			Total: total,
		}, FindSuccess, c)
	}
}
