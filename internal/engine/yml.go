package engine

import "github.com/gin-gonic/gin"

type Yml struct {
	Mode       string      `yaml:"mode"`       // gin模式 debug, release
	Middleware string      `yaml:"middleware"` // 中间件, 用逗号分隔, 暂时支持cors, trace, logs, 不填则默认全部开启, - 表示不开启
	Engine     *gin.Engine // gin引擎
}
