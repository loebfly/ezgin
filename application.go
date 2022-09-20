package gokit

import (
	"context"
	"github.com/loebfly/klogs"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type application struct{}

// StopWhenException 服务异常退出时 优雅关闭服务
func (application) StopWhenException() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan
	klogs.Error("收到退出信号:{}", sig.String())
	klogs.Error("关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, server := range servers {
		if server != nil {
			if err := server.Shutdown(ctx); err != nil {
				klogs.Error("关闭服务失败:{}", err.Error())
				return
			}
		}
	}
	klogs.Error("服务已关闭")
}
