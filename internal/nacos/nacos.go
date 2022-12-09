package nacos

import (
	"errors"
	"github.com/loebfly/ezgin/ezcache"
	"github.com/loebfly/ezgin/ezlogs"
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

const (
	CacheTableService   = "nacos_service"
	CacheTableSubscribe = "nacos_subscribe"

	CacheDuration = 5 * time.Minute
)

// GetClient 获取Nacos客户端
func (c *control) getClient() naming_client.INamingClient {
	return c.client
}

// Register 注册Nacos客户端
func (c *control) register() bool {
	c.createCacheDirIfNeed()

	var serverConfigs = make([]constant.ServerConfig, 0)

	servers := strings.Split(Config.Nacos.Server, ",")
	ports := strings.Split(Config.Nacos.Port, ",")
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

	ezlogs.CDebug("NACOS", "服务器配置:{}", serverConfigs)

	clientConfig := constant.ClientConfig{
		UpdateCacheWhenEmpty: true,
		LogLevel:             "error",
	}
	ezlogs.CDebug("NACOS", "客户端配置:{}", clientConfig)

	naming, err := clients.CreateNamingClient(map[string]any{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		ezlogs.CError("NACOS", "客户端创建失败:{}", err)
		return false
	}
	c.client = naming

	ezlogs.CInfo("NACOS", "客户端创建成功")

	appIp := ""
	if Config.Nacos.App.Ip != "" {
		appIp = Config.Nacos.App.Ip
	} else {
		ips := c.getLocalIPV4()
		if len(ips) > 0 {
			appIp = ips[0]
		}
	}

	metadata := make(map[string]string)

	port := uint64(Config.Nacos.App.Port)
	if Config.Nacos.App.PortSsl > 0 &&
		Config.Nacos.App.Port == 0 {
		metadata["ssl"] = "true"
		port = uint64(Config.Nacos.App.PortSsl)
	}
	if Config.Nacos.App.Debug {
		metadata["debug"] = "true"
	}

	isSuccess, regErr := c.client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          appIp,
		Port:        port,
		Weight:      float64(Config.Nacos.Weight),
		Enable:      true,
		Healthy:     true,
		Metadata:    metadata,
		ClusterName: Config.Nacos.ClusterName,
		ServiceName: Config.Nacos.App.Name,
		GroupName:   Config.Nacos.GroupName,
		Ephemeral:   true,
	})
	if !isSuccess {
		ezlogs.CError("NACOS", "客户端注册失败:{}", regErr)
		return false
	}

	subErr := c.client.Subscribe(&vo.SubscribeParam{
		ServiceName: Config.Nacos.App.Name,
		Clusters:    []string{Config.Nacos.ClusterName},
		GroupName:   Config.Nacos.GroupName,
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			if err != nil {
				ezlogs.CError("NACOS", "客户端订阅错误:{}", err)
				return
			}
			ezlogs.CInfo("NACOS", "客户端订阅成功:{}", services)
		},
	})
	if subErr != nil {
		ezlogs.CError("NACOS", "客户端订阅失败:{}", subErr)
		return false
	}
	return true
}

// Unregister 注销Nacos客户端
func (c *control) unregister() {
	if c.client == nil {
		return
	}
	ezlogs.CWarn("NACOS", "正在注销客户端")
	subErr := c.client.Unsubscribe(&vo.SubscribeParam{
		ServiceName: Config.Nacos.App.Name,
		Clusters:    []string{Config.Nacos.ClusterName},
		GroupName:   Config.Nacos.GroupName,
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			if err != nil {
				ezlogs.CError("NACOS", "客户端取消订阅错误:{}", err)
				return
			}
			ezlogs.CDebug("NACOS", "客户端取消订阅成功:{}", services)
		},
	})
	if subErr != nil {
		ezlogs.CError("NACOS", "客户端取消订阅失败:{}", subErr)
	}

	appIp := ""
	if Config.Nacos.App.Ip != "" {
		appIp = Config.Nacos.App.Ip
	} else {
		ips := c.getLocalIPV4()
		if len(ips) > 0 {
			appIp = ips[0]
		}
	}

	port := uint64(Config.Nacos.App.Port)
	if Config.Nacos.App.PortSsl > 0 &&
		Config.Nacos.App.Port == 0 {
		port = uint64(Config.Nacos.App.PortSsl)
	}

	isSuccess, regErr := c.client.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          appIp,
		Port:        port,
		Cluster:     Config.Nacos.ClusterName,
		ServiceName: Config.Nacos.App.Name,
		GroupName:   Config.Nacos.GroupName,
		Ephemeral:   true,
	})
	if !isSuccess {
		ezlogs.CError("NACOS", "客户端注销失败:{}", regErr)
		return
	}
	ezlogs.CWarn("NACOS", "客户端注销成功")
}

// getService 获取服务
func (c *control) getService(name string) (url string, err error) {
	// 从缓存中获取, 有则随机返回缓存中的一个
	if cacheUrls, isExist := ezcache.Table(CacheTableService).Get(name); cacheUrls != nil && isExist {
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
		group = Config.Nacos.GroupName
		ezlogs.CDebug("NACOS", "尝试从{}组中获取到服务:{}", group, name)
		instances, err := c.client.SelectInstances(vo.SelectInstancesParam{
			Clusters:    []string{Config.Nacos.ClusterName, "DEFAULT"},
			ServiceName: name,
			GroupName:   group,
			HealthyOnly: true,
		})
		if instances == nil || len(instances) == 0 || err != nil {
			ezlogs.CWarn("NACOS", "未从{}组中获取到服务:{}", group, name)
			group = "DEFAULT_GROUP"
			ezlogs.CDebug("NACOS", "尝试从{}组中获取服务:{}", group, name)
			instances, err = c.client.SelectInstances(vo.SelectInstancesParam{
				Clusters:    []string{Config.Nacos.ClusterName, "DEFAULT"},
				ServiceName: name,
				GroupName:   group,
				HealthyOnly: true,
			})
			if instances == nil || len(instances) == 0 || err != nil {
				ezlogs.CWarn("NACOS", "未从从{}组中获取服务:{}", group, name)
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
		ezlogs.CError("NACOS", "未获取到服务:{}", name)
		return "", errors.New("未获取到服务:" + name)
	}
	err = nil
	ezlogs.CInfo("NACOS", "获取到服务:{}", targetInstance)
	url = targetInstance.Ip + ":" + strconv.Itoa(int(targetInstance.Port))
	if targetInstance.Metadata != nil &&
		targetInstance.Metadata["ssl"] == "true" {
		url = "https:" + "//" + url
	} else {
		url = "http:" + "//" + url
	}

	// 订阅服务，回调中更新缓存
	if !ezcache.Table(CacheTableSubscribe).IsExist(name) {
		subErr := c.subscribeService(name, group)
		if subErr != nil {
			ezlogs.CError("NACOS", "客户端订阅服务:{}失败:{}", name, subErr.Error())
		}
		ezcache.Table(CacheTableSubscribe).Add(name, true, CacheDuration)
	}
	return url, err
}

func (c *control) subscribeService(serviceName, groupName string) error {
	return c.client.Subscribe(&vo.SubscribeParam{
		ServiceName: serviceName,
		Clusters:    []string{"DEFAULT"},
		GroupName:   groupName,
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			ezlogs.CDebug("NACOS", "处理订阅服务回调:{}", services)
			if err != nil {
				ezlogs.CError("NACOS", "订阅回调错误:{}", err.Error())
				return
			}
			if services == nil || len(services) == 0 {
				ezlogs.CError("NACOS", "订阅回调服务列表为空")
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
				if _, ok := servicesMap[serviceName]; !ok {
					servicesMap[serviceName] = []string{host}
				} else {
					servicesMap[serviceName] = append(servicesMap[serviceName], host)
				}
			}
			for sName, hosts := range servicesMap {
				if ezcache.Table(CacheTableService).IsExist(sName) {
					ezcache.Table(CacheTableService).Delete(sName)
				}
				ezlogs.CInfo("NACOS", "添加{}服务缓存,列表:{}", sName, hosts)
				ezcache.Table(CacheTableService).Add(sName, hosts, CacheDuration)
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
		if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.IsGlobalUnicast() && ipNet.IP.To4() != nil {
			if ipNet.IP.IsPrivate() {
				lanIps = append(lanIps, ipNet.IP.String())
				if Config.Nacos.Lan &&
					strings.HasSuffix(ipNet.IP.String(), Config.Nacos.LanNet) {
					ips = append(ips, ipNet.IP.String())
				}
			} else {
				wanIps = append(wanIps, ipNet.IP.String())
				if Config.Nacos.Lan == false {
					ips = append(ips, ipNet.IP.String())
				}
			}
		}
	}
	if len(ips) == 0 {
		if Config.Nacos.Lan {
			ips = append(ips, wanIps...)
		} else {
			ips = append(ips, lanIps...)
		}
	}
	return ips
}
