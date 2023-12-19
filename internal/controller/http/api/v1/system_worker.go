package v1

import (
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/response"

	"github.com/gin-gonic/gin"
)

type SystemWorkerApi struct {
	WorkerUsecase entity.WorkerUsecase
}

const (
	ExitSuccess = "退出成功"
)

func (a SystemWorkerApi) Exit(c *gin.Context) {
	var req entity.ExitWorker
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequestWithMessage(err.Error(), c)
		return
	}
	a.WorkerUsecase.AddExit(req.UUID)
	response.OkWithMessage(ExitSuccess, c)
}

func (a SystemWorkerApi) FindList(c *gin.Context) {
	list, total, err := a.WorkerUsecase.FindList()
	if err != nil {
		response.ErrorWithMessage(FindFail, err, c)
	} else {
		response.Ok(entity.PageResult{
			List:  list,
			Total: total,
		}, FindSuccess, c)
	}
}
