package engine

import (
	"github.com/gin-gonic/gin"
	engineDefine "github.com/loebfly/ezgin/engine"
)

type Yml struct {
	Mode         string                    // gin模式 debug, release
	Middleware   string                    // 中间件, 用逗号分隔, 暂时支持cors,trace,logs,recover,xlang, 不填则默认全部开启, - 表示不开启
	LogChan      chan engineDefine.ReqCtx  // 日志通道
	Engine       *gin.Engine               // gin引擎
	RecoveryFunc engineDefine.RecoveryFunc // 异常回调
}
