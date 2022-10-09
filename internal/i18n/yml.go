package i18n

type Yml struct {
	AppName    []string // i18n应用名称, 多个用逗号分隔, 默认为default,${App.Name}
	ServerName string   // i18n微服务名称
	CheckUri   string   // i18n服务检查uri, 返回是否有升级
	QueryUri   string   // i18n服务查询uri, 返回map[string]string
	Duration   int      //  i18n服务查询间隔
}
