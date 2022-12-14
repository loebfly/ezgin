package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

type enter int

const Enter = enter(0)

// InitObj 初始化Nacos客户端
func InitObj(obj EZGinNacos) {
	Config.initObj(obj)
	ctl.register()
}

// DeInit 注销Nacos客户端
func DeInit() {
	if ctl.client != nil {
		ctl.unregister()
	}
}

// GetClient 获取Nacos客户端
func (enter) GetClient() naming_client.INamingClient {
	return ctl.getClient()
}

// GetService 获取服务
func (enter) GetService(name string) (url string, err error) {
	return ctl.getService(name)
}

// CleanServiceCache 清理缓存
func (enter) CleanServiceCache(name string) {
	ctl.cleanServiceCache(name)
}
