package app

import (
	"flag"
	"fmt"
	appDefine "github.com/loebfly/ezgin/app"
	engineDefine "github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/internal/config"
	"github.com/loebfly/ezgin/internal/dblite"
	"github.com/loebfly/ezgin/internal/dblite/mongo"
	"github.com/loebfly/ezgin/internal/dblite/mysql"
	"github.com/loebfly/ezgin/internal/dblite/redis"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/engine/middleware/reqlogs"
	"github.com/loebfly/ezgin/internal/i18n"
	"github.com/loebfly/ezgin/internal/logs"
	"github.com/loebfly/ezgin/internal/nacos"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	StartCfg *appDefine.Start
)

// getLocalYml 获取yml配置文件路径
func (receiver enter) getYml() string {
	var fileName string
	flag.StringVar(&fileName, "f", os.Args[0]+".yml", "yml配置文件名")
	flag.Parse()
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if strings.Contains(fileName, "/") || strings.Contains(fileName, "\\") {
		return fileName
	}
	return path + "/" + fileName
}

// initPath 初始化所有组件入口
func (receiver enter) initEZGin(start ...appDefine.Start) {
	if len(start) > 0 {
		StartCfg = &start[0]
	}
	receiver.initConfig()
}

func (receiver enter) initConfig() {
	ymlPath := ""
	if StartCfg != nil && StartCfg.YmlPath != "" {
		ymlPath = StartCfg.YmlPath
	} else {
		ymlPath = receiver.getYml()
	}
	config.InitPath(ymlPath)
	receiver.initLogs()
}

// initLogs 初始化日志模块
func (receiver enter) initLogs() {
	out := config.EZGin().Logs.Out
	file := config.EZGin().Logs.File
	if file == "" {
		file = "/opt/logs/" + config.EZGin().App.Name
	}
	yml := logs.Yml{
		Out:  out,
		File: file,
	}
	logs.InitObj(yml)
	receiver.initNacos()
}

// initNacos 初始化nacos
func (receiver enter) initNacos() {
	ez := config.EZGin()
	if ez.Nacos.Server != "" && ez.Nacos.Yml.Nacos != "" {
		nacosPrefix := ez.Nacos.Yml.Nacos
		if nacosPrefix != "" {
			nacosUrl := ez.GetNacosUrl(nacosPrefix)
			var yml nacos.Yml
			err := config.Enter.GetYmlObj(nacosUrl, &yml)
			if err != nil {
				panic(fmt.Errorf("nacos配置文件获取失败: %s", err.Error()))
			}
			yml.EZGinNacos.App = nacos.App{
				Name:    ez.App.Name,
				Ip:      ez.App.Ip,
				Port:    ez.App.Port,
				PortSsl: ez.App.PortSsl,
				Debug:   ez.App.Debug,
			}
			nacos.InitObj(yml.EZGinNacos)
		}
	}
	receiver.initDBLite()
}

func (receiver enter) initDBLite() {
	ez := config.EZGin()

	var mongoObjs []mongo.EZGinMongo
	if ez.Nacos.Yml.Mongo != "" {
		names := strings.Split(ez.Nacos.Yml.Mongo, ",")
		mongoObjs = make([]mongo.EZGinMongo, len(names))
		for i, name := range names {
			var yml mongo.Yml
			nacosUrl := ez.GetNacosUrl(name)
			err := config.Enter.GetYmlObj(nacosUrl, &yml)
			if err != nil {
				panic(fmt.Errorf("mysql配置文件获取失败: %s", err.Error()))
			}
			mongoObjs[i] = yml.EZGinMongo
		}
	}

	var mysqlObjs []mysql.EZGinMysql
	if ez.Nacos.Yml.Mysql != "" {
		names := strings.Split(ez.Nacos.Yml.Mysql, ",")
		mysqlObjs = make([]mysql.EZGinMysql, len(names))
		for i, name := range names {
			var yml mysql.Yml
			nacosUrl := ez.GetNacosUrl(name)
			err := config.Enter.GetYmlObj(nacosUrl, &yml)
			if err != nil {
				panic(fmt.Errorf("mysql配置文件获取失败: %s", err.Error()))
			}
			mysqlObjs[i] = yml.EZGinMysql
		}
	}
	var redisObjs []redis.EZGinRedis
	if ez.Nacos.Yml.Redis != "" {
		names := strings.Split(ez.Nacos.Yml.Redis, ",")
		redisObjs = make([]redis.EZGinRedis, len(names))
		for i, name := range names {
			var yml redis.Yml
			nacosUrl := ez.GetNacosUrl(name)
			err := config.Enter.GetYmlObj(nacosUrl, &yml)
			if err != nil {
				panic(fmt.Errorf("mysql配置文件获取失败: %s", err.Error()))
			}
			redisObjs[i] = yml.EZGinRedis
		}
	}
	dblite.InitDB(mongoObjs, mysqlObjs, redisObjs)
	receiver.initEngine()
}

// initEngine 初始化gin引擎
func (receiver enter) initEngine() {
	ez := config.EZGin()

	logMongoTag := ez.Gin.MwLogs.MongoTag
	if logMongoTag != "-" && logMongoTag != "" {
		if !dblite.IsExistMongoTag(logMongoTag) {
			panic(fmt.Errorf("mongo_tag:%s不存在", logMongoTag))
		}
	}

	logTable := ez.Gin.MwLogs.Table
	if logTable == "" {
		logTable = ez.App.Name + "APIRequestLogs"
	}
	var logChan chan reqlogs.ReqCtx
	if logMongoTag != "-" {
		logChan = make(chan reqlogs.ReqCtx, 1000)
		go func(tag, tableName string) {
			for ctx := range logChan {
				var db *mgo.Database
				var returnDB func(db *mgo.Database)
				var err error
				if tag != "" {
					db, returnDB, err = dblite.Enter.Mongo(tag)
				} else {
					db, returnDB, err = dblite.Enter.Mongo()
				}

				if err != nil {
					logs.Enter.CError("MIDDLEWARE", "写入日志失败, 获取数据库失败: %s", err.Error())
					return
				}
				ctx.Id = bson.NewObjectId()
				err = db.C(tableName).Insert(ctx)
				if err != nil {
					logs.Enter.CError("MIDDLEWARE", "写入日志失败: %s", err.Error())
					returnDB(db)
				}
				returnDB(db)
			}
		}(logMongoTag, logTable)
	}

	yml := engine.Yml{
		Mode:         ez.Gin.Mode,
		Middleware:   ez.Gin.Middleware,
		Engine:       StartCfg.GinCfg.Engine,
		LogChan:      logChan,
		RecoveryFunc: engineDefine.RecoveryFunc(StartCfg.GinCfg.RecoveryHandler),
	}
	engine.InitObj(yml)

	receiver.initServer()

	// 初始化404路由
	if StartCfg != nil && StartCfg.GinCfg.NoRouteHandler != nil {
		engine.Enter.GetOriGin().NoRoute(StartCfg.GinCfg.NoRouteHandler)
	}
	// 初始化文档路由
	if StartCfg != nil && StartCfg.GinCfg.SwaggerRelativePath != "" && StartCfg.GinCfg.SwaggerHandler != nil {
		engine.Enter.GetOriGin().GET(StartCfg.GinCfg.SwaggerRelativePath, StartCfg.GinCfg.SwaggerHandler)
	}
}

// initServer 初始化服务
func (receiver enter) initServer() {
	ez := config.EZGin()

	if ez.App.Port > 0 {
		// HTTP 端口
		servers = append(servers, &http.Server{
			Addr:    ":" + strconv.Itoa(ez.App.Port),
			Handler: engine.Enter.GetOriGin(),
		})
		go func() {
			listenErr := servers[0].ListenAndServe()
			logs.Enter.CWarn("APP", "ListenAndServe:{}:{}", ez.App.Port, listenErr.Error())
		}()
	}
	if ez.App.PortSsl > 0 {
		// HTTPS 端口
		servers = append(servers, &http.Server{
			Addr:    ":" + strconv.Itoa(ez.App.PortSsl),
			Handler: engine.Enter.GetOriGin(),
		})
		go func() {
			path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
			listenErr := servers[1].ListenAndServeTLS(path+"/"+ez.App.Cert, path+"/"+ez.App.Key)
			logs.Enter.CWarn("APP", "ListenAndServeTLS:{}:{}", ez.App.PortSsl, listenErr.Error())
		}()
	}
	receiver.initI18n()
}

func (receiver enter) initI18n() {
	ez := config.EZGin()
	if ez.I18n.AppName != "-" {
		appName := ez.I18n.AppName
		if appName == "" {
			appName = "default," + ez.App.Name
		}
		var yml = i18n.Yml{
			AppName:    strings.Split(appName, ","),
			ServerName: ez.I18n.ServerName,
			CheckUri:   ez.I18n.CheckUri,
			QueryUri:   ez.I18n.QueryUri,
			Duration:   ez.I18n.Duration,
		}
		i18n.InitObj(yml)
	}
}
