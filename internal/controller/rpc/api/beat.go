package api

import (
	"context"
	"encoding/json"
	"fmt"
	"go-clean-template/internal/entity"
	"go-clean-template/pkg/worker"
	"time"

	amqprpc "github.com/0x4b53/amqp-rpc/v3"
	amqp "github.com/rabbitmq/amqp091-go"
)

type BeatApi struct {
	WorkerUsecase entity.WorkerUsecase
}

func (a BeatApi) Beat(c context.Context, rw *amqprpc.ResponseWriter, d amqp.Delivery) {
	var workerStatus worker.Status
	err := json.Unmarshal(d.Body, &workerStatus)
	if err != nil {
		_, _ = fmt.Fprintf(rw, err.Error())
		return
	}
	if a.WorkerUsecase.CheckExit(workerStatus.UUID) {
		_, _ = fmt.Fprintf(rw, "exit")
		return
	}
	workerStatus.UpdatedAt = time.Now()
	a.WorkerUsecase.Update(&workerStatus)
	_, _ = fmt.Fprintf(rw, "beat")
}
