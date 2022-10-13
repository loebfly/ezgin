package engine

import (
	"github.com/gin-gonic/gin"
)

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
		cfg.Gin.Middleware = "cors,trace,logs,xlang,recover"
	}
	if cfg.Gin.Engine == nil {
		cfg.Gin.Engine = gin.Default()
	}
}
