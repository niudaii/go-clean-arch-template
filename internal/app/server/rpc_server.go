package server

import (
	"go-clean-template/internal/controller/rpc/middleware"
	"go-clean-template/internal/controller/rpc/router"

	amqprpc "github.com/0x4b53/amqp-rpc/v3"
	amqprpcmw "github.com/0x4b53/amqp-rpc/v3/middleware"
	"go.uber.org/zap"
)

func RunRPCServer(url string) {
	s := amqprpc.NewServer(url)
	rpcHandler(s)
	s.ListenAndServe()
}

func rpcHandler(s *amqprpc.Server) {
	logger := zap.L().Named("[amqp-rpc]")
	s.AddMiddleware(amqprpcmw.PanicRecoveryLogging(middleware.ErrorLogger(logger)))
	s.WithErrorLogger(middleware.ErrorLogger(logger))
	s.WithDebugLogger(middleware.DebugLogger(logger))

	router.NewBeatRouter(s)
}
