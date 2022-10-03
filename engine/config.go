package engine

import "github.com/gin-gonic/gin"

type Yml struct {
	Mode       string      `yaml:"mode"`       // gin模式 debug, release
	Middleware string      `yaml:"middleware"` // 中间件, 用逗号分隔, 暂时支持cors, trace, logs, 不填则默认全部开启, - 表示不开启
	engine     *gin.Engine // gin引擎
}

type ymlConfig struct {
	Gin Yml `yml:"gin"` // gin配置
}

var config = new(ymlConfig)

func (cfg *ymlConfig) initObj(obj Yml) {
	cfg.Gin = obj
	cfg.fillNull()
}

func (cfg *ymlConfig) fillNull() {
	if cfg.Gin.Mode == "" {
		cfg.Gin.Mode = gin.ReleaseMode
	}
	if cfg.Gin.Middleware == "" {
		cfg.Gin.Middleware = "cors,trace,logs"
	}
	if cfg.Gin.engine == nil {
		cfg.Gin.engine = gin.Default()
	}
}
