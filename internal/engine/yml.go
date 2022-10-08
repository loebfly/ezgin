package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/internal/engine/middleware/reqlogs"
)

type Yml struct {
	Mode       string              // gin模式 debug, release
	Middleware string              // 中间件, 用逗号分隔, 暂时支持cors, trace, logs, 不填则默认全部开启, - 表示不开启
	LogChan    chan reqlogs.ReqCtx // 日志通道
	Engine     *gin.Engine         // gin引擎
}
