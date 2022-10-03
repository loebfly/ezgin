package internal

import "github.com/loebfly/ezgin/internal/nacos"

func InitObj(yml nacos.Yml) {
	Config.initObj(yml)
}
