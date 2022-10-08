package engine

import "github.com/gin-gonic/gin"

type Yml struct {
	Mode       string      // gin模式 debug, release
	Middleware string      // 中间件, 用逗号分隔, 暂时支持cors, trace, logs, 不填则默认全部开启, - 表示不开启
	MwLogs     MwLogs      // 日志配置
	Engine     *gin.Engine // gin引擎
}

type MwLogs struct {
	MongoTag string // 需要与Nacos-Yml-Mongo中配置文件中的tag一致
	Table    string // 日志表名
}
