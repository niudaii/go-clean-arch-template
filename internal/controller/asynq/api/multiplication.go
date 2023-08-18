package api

import (
	"encoding/json"
	"go-clean-template/internal/entity"
	"strconv"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func runMultiplication(t *asynq.Task, inputs []string) (err error) {
	result := 1
	for _, input := range inputs {
		var num int
		num, err = strconv.Atoi(input)
		if err != nil {
			return
		}
		result *= num
	}
	// 整理结果
	zap.L().Info("[multiplication]", zap.Any("result", result))
	resp := entity.Result{
		Inputs: inputs,
		Result: result,
	}
	// 结果回传
	var data []byte
	data, err = json.Marshal(resp)
	_, err = t.ResultWriter().Write(data)
	return
}
