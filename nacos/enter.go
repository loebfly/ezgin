package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

type enter int

const Enter = enter(0)

func (enter) InitObj(obj Yml) error {
	err := config.InitObj(obj)
	if err != nil {
		return err
	}
	isSuccess := ctl.register()
	if !isSuccess {
		return errors.New("nacos register error")
	}
	return nil
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
