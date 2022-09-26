package nacos

import (
	"errors"
	"github.com/loebfly/ezgin/cache"
	"github.com/loebfly/klogs"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var ctl = new(control)

type control struct {
	client naming_client.INamingClient
}

// GetClient 获取Nacos客户端
func (c *control) getClient() naming_client.INamingClient {
	return c.client
}

// Register 注册Nacos客户端
func (c *control) register() bool {
	c.createCacheDirIfNeed()

	var serverConfigs = make([]constant.ServerConfig, 0)

	servers := strings.Split(config.Nacos.Server, ",")
	ports := strings.Split(config.Nacos.Port, ",")
	for i, server := range servers {
		portIdx := 0
		if i < len(ports) {
			portIdx = i
		}
		port, _ := strconv.Atoi(ports[portIdx])
		serverConfigs = append(serverConfigs, constant.ServerConfig{
			IpAddr: server,
			Port:   uint64(port),
		})
	}

	klogs.C("NACOS").Debug("Nacos服务器配置:{}", serverConfigs)

	clientConfig := constant.ClientConfig{
		UpdateCacheWhenEmpty: true,
		LogLevel:             "error",
	}
	klogs.C("NACOS").Debug("Nacos客户端配置:{}", clientConfig)

	naming, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		klogs.C("NACOS").Error("Nacos客户端创建失败:{}", err)
		return false
	}
	c.client = naming

	klogs.C("NACOS").Debug("Nacos客户端创建成功")

	appIp := ""
	if config.Nacos.App.Ip != "" {
		appIp = config.Nacos.App.Ip
	} else {
		ips := c.getLocalIPV4()
		if len(ips) > 0 {
			appIp = ips[0]
		}
	}

	metadata := make(map[string]string)

	port := uint64(config.Nacos.App.Port)
	if config.Nacos.App.PortSsl > 0 &&
		config.Nacos.App.Port == 0 {
		metadata["ssl"] = "true"
		port = uint64(config.Nacos.App.PortSsl)
	}
	if config.Nacos.App.Debug {
		metadata["debug"] = "true"
	}

	isSuccess, regErr := c.client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          appIp,
		Port:        port,
		Weight:      float64(config.Nacos.Weight),
		Enable:      true,
		Healthy:     true,
		Metadata:    metadata,
		ClusterName: config.Nacos.ClusterName,
		ServiceName: config.Nacos.App.Name,
		GroupName:   config.Nacos.GroupName,
		Ephemeral:   true,
	})
	if !isSuccess {
		klogs.C("NACOS").Error("Nacos客户端注册失败:{}", regErr)
		return false
	}

	subErr := c.client.Subscribe(&vo.SubscribeParam{
		ServiceName: config.Nacos.App.Name,
		Clusters:    []string{config.Nacos.ClusterName},
		GroupName:   config.Nacos.GroupName,
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			if err != nil {
				klogs.C("NACOS").Error("Nacos客户端订阅错误:{}", err)
				return
			}
			klogs.C("NACOS").Debug("Nacos客户端订阅成功:{}", services)
		},
	})
	if subErr != nil {
		klogs.C("NACOS").Error("Nacos客户端订阅失败:{}", subErr)
		return false
	}
	return true
}

// Unregister 注销Nacos客户端
func (c *control) unregister() {
	if c.client == nil {
		klogs.C("NACOS").Error("Nacos客户端未注册")
		return
	}
	subErr := c.client.Unsubscribe(&vo.SubscribeParam{
		ServiceName: config.Nacos.App.Name,
		Clusters:    []string{config.Nacos.ClusterName},
		GroupName:   config.Nacos.GroupName,
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			if err != nil {
				klogs.C("NACOS").Error("Nacos客户端取消订阅错误:{}", err)
				return
			}
			klogs.C("NACOS").Debug("Nacos客户端取消订阅成功:{}", services)
		},
	})
	if subErr != nil {
		klogs.C("NACOS").Error("Nacos客户端取消订阅失败:{}", subErr)
	}

	appIp := ""
	if config.Nacos.App.Ip != "" {
		appIp = config.Nacos.App.Ip
	} else {
		ips := c.getLocalIPV4()
		if len(ips) > 0 {
			appIp = ips[0]
		}
	}

	port := uint64(config.Nacos.App.Port)
	if config.Nacos.App.PortSsl > 0 &&
		config.Nacos.App.Port == 0 {
		port = uint64(config.Nacos.App.PortSsl)
	}

	isSuccess, regErr := c.client.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          appIp,
		Port:        port,
		Cluster:     config.Nacos.ClusterName,
		ServiceName: config.Nacos.App.Name,
		GroupName:   config.Nacos.GroupName,
		Ephemeral:   true,
	})
	if !isSuccess {
		klogs.C("NACOS").Error("Nacos客户端注销失败:{}", regErr)
		return
	}
}

// getService 获取服务
func (c *control) getService(name string) (url string, err error) {
	// 从缓存中获取, 有则随机返回缓存中的一个
	if cacheUrls, isExist := cache.Enter.Table("nacos").Get(name); cacheUrls != nil && isExist {
		if urls, ok := cacheUrls.([]string); ok {
			if len(urls) > 0 {
				url = urls[rand.Intn(len(urls))]
				return
			}
		}
	}

	// 从Nacos中获取, 有则随机返回Nacos中的一个
	var targetInstance model.Instance
	var group string
	// 最多尝试3次
	for i := 0; i < 3; i++ {
		group = config.Nacos.GroupName
		klogs.C("NACOS").Debug("尝试从{}组中获取到服务:{}", group, name)
		instances, err := c.client.SelectInstances(vo.SelectInstancesParam{
			Clusters:    []string{config.Nacos.ClusterName, "DEFAULT"},
			ServiceName: name,
			GroupName:   group,
			HealthyOnly: true,
		})
		if instances == nil || len(instances) == 0 || err != nil {
			klogs.C("NACOS").Warn("未从{}组中获取到服务:{}", group, name)
			group = "DEFAULT_GROUP"
			klogs.C("NACOS").Debug("尝试从{}组中获取服务:{}", group, name)
			instances, err = c.client.SelectInstances(vo.SelectInstancesParam{
				Clusters:    []string{config.Nacos.ClusterName, "DEFAULT"},
				ServiceName: name,
				GroupName:   group,
				HealthyOnly: true,
			})
			if instances == nil || len(instances) == 0 || err != nil {
				klogs.C("NACOS").Warn("未从从{}组中获取服务:{}", group, name)
				continue
			}
		}

		// 删除debug实例
		for j := 0; j < len(instances); j++ {
			if instances[j].Metadata["debug"] == "true" {
				instances = append(instances[:j], instances[j+1:]...)
				j--
			}
		}
		if len(instances) > 0 {
			targetInstance = instances[rand.Intn(len(instances))]
			break
		}
	}

	if targetInstance.InstanceId == "" {
		klogs.C("NACOS").Error("未获取到服务:{}", name)
		return "", errors.New("未获取到服务:" + name)
	}
	err = nil
	klogs.C("NACOS").Debug("获取到服务:{}", targetInstance)
	url = targetInstance.Ip + ":" + strconv.Itoa(int(targetInstance.Port))
	if targetInstance.Metadata != nil &&
		targetInstance.Metadata["ssl"] == "true" {
		url = "https:" + "//" + url
	} else {
		url = "http:" + "//" + url
	}

	// 订阅服务，回调中更新缓存
	subErr := c.subscribeService(name, group)
	if subErr != nil {
		klogs.C("NACOS").Error("客户端订阅服务:{}失败:{}", name, subErr.Error())
	}

	return url, err
}

func (c *control) subscribeService(serviceName, groupName string) error {
	return c.client.Subscribe(&vo.SubscribeParam{
		ServiceName: serviceName,
		Clusters:    []string{"DEFAULT"},
		GroupName:   groupName,
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			klogs.C("NACOS").Debug("处理订阅服务回调:{}", services)
			if err != nil {
				klogs.Error("Nacos订阅回调错误:{}", err.Error())
				return
			}
			if services == nil || len(services) == 0 {
				klogs.C("NACOS").Error("订阅回调服务列表为空")
				return
			}
			servicesMap := make(map[string][]string)
			for _, s := range services {
				protocol := "http"
				if s.Metadata != nil {
					if s.Metadata["debug"] == "true" {
						continue
					}
					if s.Metadata["ssl"] == "true" {
						protocol = "https"
					}
				}

				host := protocol + "://" + s.Ip + ":" + strconv.Itoa(int(s.Port))
				if _, ok := servicesMap[s.ServiceName]; !ok {
					servicesMap[s.ServiceName] = []string{host}
				} else {
					servicesMap[s.ServiceName] = append(servicesMap[s.ServiceName], host)
				}
			}
			for sName, hosts := range servicesMap {
				if cache.Enter.Table("NACOS").IsExist(sName) {
					cache.Enter.Table("NACOS").Delete(sName)
				}
				klogs.C("NACOS").Debug("添加{}服务缓存,列表:{}", sName, hosts)
				cache.Enter.Table("NACOS").Add(sName, hosts, time.Minute*5)
			}
		},
	})
}

// createCacheDirIfNeed 创建缓存目录
func (c *control) createCacheDirIfNeed() {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	path += "/cache"
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		_ = os.Mkdir(path, 0777)
		path += "/naming"
		_ = os.Mkdir(path, 0777)
	}
}

// 获取本地IPV4地址列表
func (c *control) getLocalIPV4() []string {
	var ips, lanIps, wanIps = make([]string, 0), make([]string, 0), make([]string, 0)
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		return ips
	}
	for _, addr := range addrList {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.IsPrivate() {
				lanIps = append(lanIps, ipNet.IP.String())
				if config.Nacos.Lan &&
					strings.HasSuffix(ipNet.IP.String(), config.Nacos.LanNet) {
					ips = append(ips, ipNet.IP.String())
				}
			} else {
				wanIps = append(wanIps, ipNet.IP.String())
				if config.Nacos.Lan == false {
					ips = append(ips, ipNet.IP.String())
				}
			}
		}
	}
	if len(ips) == 0 {
		if config.Nacos.Lan {
			ips = append(ips, wanIps...)
		} else {
			ips = append(ips, lanIps...)
		}
	}
	return ips
}
