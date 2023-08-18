package router

import (
	"go-clean-template/internal/controller/rpc/api"
	"go-clean-template/internal/usecase"

	amqprpc "github.com/0x4b53/amqp-rpc/v3"
)

func NewBeatRouter(s *amqprpc.Server) {
	beatApi := &api.BeatApi{
		WorkerUsecase: usecase.NewWorkerUsecase(),
	}
	s.Bind(amqprpc.DirectBinding("beat", beatApi.Beat))
}
