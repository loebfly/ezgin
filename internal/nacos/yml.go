package nacos

type Yml struct {
	EZGinNacos EZGinNacos `koanf:"ezgin_nacos"`
}

type EZGinNacos struct {
	Server      string `koanf:"server"`       // 服务器ip地址 必填 多个用,分隔
	Port        string `koanf:"port"`         // 服务器端口 必填 多个用,分隔
	ClusterName string `koanf:"cluster_name"` // 集群名称 默认为DEFAULT
	GroupName   string `koanf:"group_name"`   // 组名称 默认为DEFAULT_GROUP
	Weight      int    `koanf:"weight"`       // 权重 默认1
	Lan         bool   `koanf:"lan"`          // 是否为内网 true:内网  false:外网
	LanNet      string `koanf:"lanNet"`       // 网段前缀 必填 例如：192.168.3.
	App         App
}

type App struct {
	Name    string // 应用名称
	Ip      string // 应用ip地址, 默认为本机ip
	Port    int    // http服务端口
	PortSsl int    // https服务端口
	Debug   bool   // 是否开启debug模式, 默认false, 如果开启, 则不会被其他服务调用
}
