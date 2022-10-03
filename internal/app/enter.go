package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"os"
)

type enter int

const Enter = enter(0)

func StartWithEngine(ymlPath string, engine *gin.Engine) {
	Enter.Start(ymlPath, engine)
}

// ShutdownWhenExitSignalWithCallBack 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignalWithCallBack(will func(os.Signal), did func(context.Context)) {
	Enter.ShutdownWhenExitSignal(will, did)
}
