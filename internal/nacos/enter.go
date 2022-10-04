package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

type enter int

const Enter = enter(0)

func InitObj(obj EZGinNacos) {
	Config.initObj(obj)
	ctl.register()
}

// GetClient 获取Nacos客户端
func (enter) GetClient() naming_client.INamingClient {
	return ctl.getClient()
}

// UnregisterIfNeed	如果需要则注销
func (enter) UnregisterIfNeed() {
	if ctl.client != nil {
		ctl.unregister()
	}
}

// GetService 获取服务
func (enter) GetService(name string) (url string, err error) {
	return ctl.getService(name)
}
