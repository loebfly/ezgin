package nacos

import (
	"github.com/loebfly/ezgin/internal/nacos/internal"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

type enter int

const Enter = enter(0)

func InitObj(obj Yml) {
	internal.InitObj(obj)
	ctl.register()
}

// GetClient 获取Nacos客户端
func (enter) GetClient() naming_client.INamingClient {
	return ctl.getClient()
}

// Unregister 注销Nacos客户端
func (enter) Unregister() {
	ctl.unregister()
}

// GetService 获取服务
func (enter) GetService(name string) (url string, err error) {
	return ctl.getService(name)
}
