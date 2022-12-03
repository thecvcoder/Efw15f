package core

import (
	"fmt"
	"github.com/thecvcoder/cloud-api-go/global"
	"github.com/thecvcoder/cloud-api-go/initialize"
	"os"
	"os/signal"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	r := initialize.InitRouter()
	addr := fmt.Sprintf("0.0.0.0:%d", global.CONFIG.System.Port)
	s := initServer(addr, r)

	// 开启一个goroutine启动服务
	go func() {
		if err := s.ListenAndServe(); err != nil {
			global.LOG.Fatalf("启动服务失败: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.LOG.Info("shutdown server!")
}
