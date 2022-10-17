package app

import (
	"context"
	appDefine "github.com/loebfly/ezgin/app"
	"github.com/loebfly/ezgin/internal/config"
	"github.com/loebfly/ezgin/internal/dblite"
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
func (receiver enter) Start(start ...appDefine.Start) {

	receiver.initEZGin(start...)
	ez := config.EZGin()

	logs.Enter.CInfo("APP", "|-----------------------------------|")
	logs.Enter.CInfo("APP", "| 服务名: {}", ez.App.Name)
	logs.Enter.CInfo("APP", "| 版本号: {}", ez.App.Version)
	logs.Enter.CInfo("APP", "|-----------------------------------|")
	if ez.App.Port > 0 {
		logs.Enter.CInfo("APP", "| HTTP端口: {}", ez.App.Port)
	}
	if ez.App.PortSsl > 0 {
		logs.Enter.CInfo("APP", "| HTTPS端口: {}", ez.App.PortSsl)
	}
	logs.Enter.CInfo("APP", "|-----------------------------------|")
}

// ShutdownWhenExitSignal 服务异常退出时 优雅关闭服务
func (receiver enter) ShutdownWhenExitSignal(shutdown ...appDefine.Shutdown) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan
	logs.Enter.CError("APP", "收到退出信号:{}", sig.String())
	nacos.DeInit()
	dblite.DeInit()

	if len(shutdown) > 0 && shutdown[0].WillHandler != nil {
		shutdown[0].WillHandler(sig)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, server := range servers {
		if server != nil {
			if err := server.Shutdown(ctx); err != nil {
				logs.Enter.CWarn("APP", "关闭{}端口失败:{}", server.Addr, err.Error())
				continue
			}
		}
	}

	if len(shutdown) > 0 && shutdown[0].DidHandler != nil {
		shutdown[0].DidHandler(ctx)
	}
}
