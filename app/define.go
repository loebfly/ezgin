package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"os"
)

type RecoveryFunc func(c *gin.Context, err interface{})

type Start struct {
	YmlPath string // yml配置文件路径, 为空时默认为当前程序所在目录的同名yml文件
	GinCfg  GinCfg // gin配置
}

type GinCfg struct {
	Engine              *gin.Engine     // gin引擎, 传nil则使用gin默认引擎
	RecoveryHandler     RecoveryFunc    // 异常恢复处理函数, 传nil则使用默认处理函数
	NoRouteHandler      gin.HandlerFunc // 404处理函数, 传nil不处理
	SwaggerRelativePath string          // swagger相对路径, 传空则不启用swagger
	SwaggerHandler      gin.HandlerFunc // swagger处理函数, 传nil则不启用swagger
}

type Shutdown struct {
	WillHandler func(os.Signal)       // 退出前回调
	DidHandler  func(context.Context) // 退出后回调
}
