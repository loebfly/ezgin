package engine

import "github.com/gin-gonic/gin"

type Yml struct {
	Mode       string `yaml:"mode"`       // gin模式 debug, release
	Middleware string `yaml:"middleware"` // 中间件, 用逗号分隔, 暂时支持cors, trace, logs, 不填则默认全部开启, - 表示不开启
	Logs       struct {
		Mongo string `koanf:"mongo"`      // 需要与Nacos-Yml-Mongo中配置文件名对应
		Table string `koanf:"table_name"` // 日志表名
	} `koanf:"logs"` // 日志配置
	Engine *gin.Engine // gin引擎
}
