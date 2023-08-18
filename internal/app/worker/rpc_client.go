package worker

import (
	"encoding/json"
	"go-clean-template/internal/controller/rpc/middleware"
	"go-clean-template/pkg/worker"
	"os"
	"time"

	amqprpc "github.com/0x4b53/amqp-rpc/v3"
	"go.uber.org/zap"
)

func RunRPCClient(url string) {
	c := amqprpc.NewClient(url)

	logger := zap.L().Named("[amqp-rpc]")
	c.WithErrorLogger(middleware.ErrorLogger(logger))
	c.WithDebugLogger(middleware.DebugLogger(logger))

	// 心跳
	for {
		req, _ := json.Marshal(worker.GetStatus())
		resp, err := c.Send(
			amqprpc.NewRequest().WithRoutingKey("beat").
				WithBody(string(req)).
				WithTimeout(5 * time.Second),
		)
		if err != nil {
			zap.L().Error("heartbeat fail", zap.Error(err))
		} else {
			body := string(resp.Body)
			logger.Debug("heartbeat success", zap.String("resp", body))
			if body == "exit" {
				logger.Info("收到 exit 信号, 程序退出")
				os.Exit(0)
			}
		}

		time.Sleep(60 * time.Second)
	}
}
