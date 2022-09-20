package gokit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var App = new(application)

var (
	servers = make([]*http.Server, 0)
)

func InitWithYml(ymlPath string) {
	engine := gin.Default()

	servers = append(servers, &http.Server{
		Addr:    ":" + "port",
		Handler: engine,
	})
	servers = append(servers, &http.Server{
		Addr:    ":" + "port_ssl",
		Handler: engine,
	})
}
