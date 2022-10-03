package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/internal/config"
	"github.com/loebfly/ezgin/internal/logs"
	"github.com/loebfly/ezgin/internal/nacos"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	servers = make([]*http.Server, 0)
)

// Start 启动服务
func (receiver enter) Start(ymlPath string, ginEngine *gin.Engine) {

	receiver.initEZGin(ymlPath, ginEngine)
	ez := config.EZGin()

	logs.Enter.CDebug("APP", "|-----------------------------------|")
	logs.Enter.CDebug("APP", "|	   			{} {}				|", ez.App.Name, ez.App.Version)
	logs.Enter.CDebug("APP", "|-----------------------------------|")
	if ez.App.Port > 0 {
		logs.Enter.CDebug("APP", "|	启动成功!		HTTP端口: {}				|", ez.App.Port)
	}
	if ez.App.PortSsl > 0 {
		logs.Enter.CDebug("APP", "|	启动成功!		HTTPS端口: {}			|", ez.App.PortSsl)
	}
	logs.Enter.CDebug("APP", "|-----------------------------------|")
}

// ShutdownWhenExitSignal 服务异常退出时 优雅关闭服务
func (receiver enter) ShutdownWhenExitSignal(will func(os.Signal), did func(context.Context)) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan
	logs.Enter.Error("收到退出信号:{}", sig.String())
	logs.Enter.Error("关闭服务...")

	nacos.Enter.Unregister()

	will(sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, server := range servers {
		if server != nil {
			if err := server.Shutdown(ctx); err != nil {
				logs.Enter.Error("关闭服务失败:{}", err.Error())
				return
			}
		}
	}

	logs.Enter.Error("服务已关闭")
	did(ctx)
}
