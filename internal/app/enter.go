package app

import (
	appDefine "github.com/loebfly/ezgin/app"
)

type enter int

const Enter = enter(0)

func Start(start ...appDefine.Start) {
	Enter.Start(start...)
}

// ShutdownWhenExitSignal 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignal(shutdown ...appDefine.Shutdown) {
	Enter.ShutdownWhenExitSignal(shutdown...)
}
