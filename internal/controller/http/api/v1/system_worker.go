package v1

import (
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/response"

	"github.com/gin-gonic/gin"
)

type SystemWorkerApi struct {
	WorkerUsecase entity.WorkerUsecase
}

func (a SystemWorkerApi) Exit(c *gin.Context) {
	var req entity.ExitWorker
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequestWithMessage(err.Error(), c)
		return
	}
	a.WorkerUsecase.AddExit(req.UUID)
	response.OkWithMessage("退出成功", c)
}

func (a SystemWorkerApi) FindList(c *gin.Context) {
	list, total, err := a.WorkerUsecase.FindList()
	if err != nil {
		response.ErrorWithMessage("查询失败", err, c)
	} else {
		response.Ok(entity.PageResult{
			List:  list,
			Total: total,
		}, "查询成功", c)
	}
}
