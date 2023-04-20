# EZGin 微服务快速开发脚手架

## 1. 安装

版本要求: go 1.19+

```bash
go get -u github.com/loebfly/ezgin
```

## 2. 配置
参考ezgin/template目录下的配置文件

## 3. 概览

- ezgin (服务启动、优化退出、gin封装)
- ezlogs (日志模块)
- ezdb (数据库模块)
- ezcfg (配置模块)
- ezi18n (国际化模块)
- ezgo (安全携程模块)
- ezcache (内存缓存模块)
- ezcall (微服务调用模块)

## 4. 示例

### 4.1 ezgin

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin"
	"github.com/loebfly/ezgin/app"
	"github.com/loebfly/ezgin/engine"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func main() {
	// 服务启动
	ezgin.Start(app.Start{
		YmlPath: "ezgin.yml", // 默认程序同级目录下的同名yml文件
		GinCfg: app.GinCfg{
			RecoveryHandler: func(c *gin.Context, err interface{}) {
				// 自定义程序崩溃错误处理
				c.JSON(http.StatusOK, i18n.SystemError.ErrorRes())
			},
			NoRouteHandler: func(c *gin.Context) {
				// 自定义404处理
				c.JSON(http.StatusOK, i18n.UrlNotFound.ErrorRes())
			},
			SwaggerRelativePath: "/docs/*any",
			SwaggerHandler:      ginSwagger.WrapHandler(swaggerFiles.Handler),
		},
	})
    /* 业务区域 */
	// 如路由注册
	ezgin.GET("/ping", func(c *gin.Context) engine.Result[any] {
		return i18n.Success.Result("pong")
	})
	
	/* 业务区域 */

	// 服务异常退出时 优雅关闭服务
	ezgin.ShutdownWhenExitSignal()
}
```

### 4.2 ezlogs

```go
package main

import (
    "github.com/loebfly/ezgin/ezlogs"
)

func main() {
	// {} 为占位符，后面的参数为占位符对应的值
	ezlogs.Debug("debug:{}", "hello, world")
	ezlogs.Info("info:{}", "hello, world")
	ezlogs.Warn("warn:{}", "hello, world")
	ezlogs.Error("error:{}", "hello, world")
}

```

### 4.3 ezdb

```go
package main

import (
	"github.com/loebfly/ezgin/ezdb"
	"github.com/loebfly/ezgin/ezlogs"
	"github.com/loebfly/ezgin/internal/dblite/kafka"
)

func main() {
	mysqlDb, err := ezdb.Mysql()
	if err != nil {
		ezlogs.Error("mysql 连接错误 error:{}", err)
		return
	}
	ezlogs.Info("mysql 连接获取成功")

	mongoDb, returnDB, err := ezdb.Mongo()
	if err != nil {
		ezlogs.Error("mongo 连接错误 error:{}", err)
		return
	}
	defer returnDB(mongoDb)
	ezlogs.Info("mongo 连接获取成功")

	redisDb, err := ezdb.Redis()
	if err != nil {
		ezlogs.Error("redis db 连接错误 error:{}", err)
		return
	}
	ezlogs.Info("redis db 连接获取成功")

	kafkaDb, err := ezdb.Kafka()
	if err != nil {
        ezlogs.Error("kafka 连接错误 error:{}", err)
        return
    }
	ezlogs.Info("kafka 连接获取成功")

}
```

### 4.4 ezcfg

```go
package main

import (
    "github.com/loebfly/ezgin/ezcfg"
    "github.com/loebfly/ezgin/ezlogs"
)

func main() {
    // 获取配置
    cfg := ezcfg.GetString("ezgin.app.name")
    ezlogs.Info("ezgin.app.name:{}", cfg)
}
```

### 4.5 ezi18n

```go
package main

import (
    "github.com/loebfly/ezgin/ezi18n"
    "github.com/loebfly/ezgin/ezlogs"
)

func main() {
	translate := ezi18n.StringId("messageId").GetTranslate()
    ezlogs.Info("translate:{}", translate)
}
```

### 4.6 ezgo

```go
package main

import (
	"github.com/loebfly/ezgin"
	"github.com/loebfly/ezgin/ezgo"
	"github.com/loebfly/ezgin/ezlogs"
)

func main() {
	// 安全携程
	safeGo := ezgo.New[string](func(args ...any) {
		ezlogs.Debug("args:{}", args)
	})
	safeGo.SetGoBeforeHandler(func() string {
		return ezgin.MWTrace.GetCurRoutineId()
	})
	safeGo.SetGoAfterHandler(func(params string) {
		ezgin.MWTrace.CopyPreHeaderToCurRoutine(params)
	})
	safeGo.Run("hello", "world")
}
```

### 4.7 ezcache

```go
package main

import (
    "github.com/loebfly/ezgin/ezcache"
    "github.com/loebfly/ezgin/ezlogs"
)

func main() {
    // 设置缓存
	ezcache.Table("ezgin").Add("key", "value", 0)
	// 获取缓存
	ezcache.Table("ezgin").Get("key")
}
```

### 4.8 ezcall

```go
package main

import (
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/ezcall"
	"github.com/loebfly/ezgin/ezlogs"
)

func main() {
	// 调用微服务
	//	1. form: 适用于form表单提交的接口
	res := ezcall.FormGetToResult[string]("微服务名称", "接口地址", map[string]string{
		"参数key": "参数值",
	})
	ezlogs.Info("res:{}", res)

	// 2. json: 适用于json格式的接口
	res := ezcall.JsonToResult[string](engine.Get, "微服务名称", "接口地址", map[string]string{
		"参数key": "参数值",
	}, nil)
	ezlogs.Info("res:{}", res)
	// 3. restful: 适用于restful风格的接口
	res := ezcall.RestfulToResult[string](engine.Get, "微服务名称", "接口地址", map[string]string{
		"路径key": "路径值",
        }, map[string]string{"参数key": "参数值"}, nil)
	ezlogs.Info("res:{}", res)
	// 4. 通用调用
	opt := ezcall.FormOptions{}
	//opt := ezcall.JsonOptions{}
	//opt := ezcall.RestfulOptions{}
	res := ezcall.RequestToResult[string](opt)
	ezlogs.Info("res:{}", res)
}

```

