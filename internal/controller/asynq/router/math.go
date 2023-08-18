package router

import (
	"go-clean-template/internal/controller/asynq/api"
	"go-clean-template/pkg/task"

	"github.com/hibiken/asynq"
)

func NewMathRouter(mux *asynq.ServeMux) {
	mux.HandleFunc(task.TypeMath, api.HandleMath)
}
