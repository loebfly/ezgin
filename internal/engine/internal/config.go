package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/internal/engine"
)

type ymlConfig struct {
	Gin engine.Yml `yml:"gin"` // gin配置
}

var Config = new(ymlConfig)

func (cfg *ymlConfig) initObj(obj engine.Yml) {
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
	if cfg.Gin.Engine == nil {
		cfg.Gin.Engine = gin.Default()
	}
}
