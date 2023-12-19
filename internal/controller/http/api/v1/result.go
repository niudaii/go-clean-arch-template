package v1

import (
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/response"

	"github.com/gin-gonic/gin"
)

const (
	FindSuccess = "查询成功"
	FindFail    = "查询失败"
)

type ResultApi struct {
	ResultUsecase entity.ResultUsecase
}

func (a ResultApi) FindList(c *gin.Context) {
	var req entity.SearchResult
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequestWithMessage(err.Error(), c)
		return
	}
	list, total, err := a.ResultUsecase.FindList(&req)
	if err != nil {
		response.ErrorWithMessage(FindFail, err, c)
	} else {
		response.Ok(entity.PageResult{
			List:  list,
			Total: total,
		}, FindSuccess, c)
	}
}
