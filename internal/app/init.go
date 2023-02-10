package app

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	appDefine "github.com/loebfly/ezgin/app"
	engineDefine "github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/ezcfg"
	"github.com/loebfly/ezgin/ezdb"
	"github.com/loebfly/ezgin/ezlogs"
	"github.com/loebfly/ezgin/internal/config"
	"github.com/loebfly/ezgin/internal/dblite"
	"github.com/loebfly/ezgin/internal/dblite/kafka"
	"github.com/loebfly/ezgin/internal/dblite/mongo"
	"github.com/loebfly/ezgin/internal/dblite/mysql"
	"github.com/loebfly/ezgin/internal/dblite/redis"
	"github.com/loebfly/ezgin/internal/engine"
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
	level := config.EZGin().Logs.Level
	out := config.EZGin().Logs.Out
	file := config.EZGin().Logs.File
	if file == "" {
		file = "/opt/logs/" + config.EZGin().App.Name
	}
	yml := logs.Yml{
		Level: level,
		Out:   out,
		File:  file,
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
			nacosUrl := ez.GetYmlUrlOrPath(nacosPrefix)
			var yml nacos.Yml
			err := ezcfg.GetYmlObj(nacosUrl, &yml)
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

	var mongoObjs = make([]mongo.EZGinMongo, 0)
	if ez.Nacos.Yml.Mongo != "" {
		names := strings.Split(ez.Nacos.Yml.Mongo, ",")
		mongoObjs = make([]mongo.EZGinMongo, len(names))
		for _, name := range names {
			var yml mongo.Yml
			nacosUrl := ez.GetYmlUrlOrPath(name)
			err := ezcfg.GetYmlObj(nacosUrl, &yml)
			if err != nil {
				panic(fmt.Errorf("mysql配置文件获取失败: %s", err.Error()))
			}
			mongoObjs = append(mongoObjs, yml.EZGinMongo)
		}
	}

	if ez.Nacos.Yml.MongoSet != "" {
		var yml mongo.SetYml
		nacosUrl := ez.GetYmlUrlOrPath(ez.Nacos.Yml.MongoSet)
		err := ezcfg.GetYmlObj(nacosUrl, &yml)
		if err != nil {
			panic(fmt.Errorf("mysql配置文件获取失败: %s", err.Error()))
		}
		for tag, cfg := range yml.MongoSet {
			cfg.Tag = tag
			mongoObjs = append(mongoObjs, cfg)
		}
	}

	var mysqlObjs = make([]mysql.EZGinMysql, 0)
	if ez.Nacos.Yml.Mysql != "" {
		names := strings.Split(ez.Nacos.Yml.Mysql, ",")
		for _, name := range names {
			var yml mysql.Yml
			nacosUrl := ez.GetYmlUrlOrPath(name)
			err := ezcfg.GetYmlObj(nacosUrl, &yml)
			if err != nil {
				panic(fmt.Errorf("mysql配置文件获取失败: %s", err.Error()))
			}
			mysqlObjs = append(mysqlObjs, yml.EZGinMysql)
		}
	}

	if ez.Nacos.Yml.MysqlSet != "" {
		var yml mysql.SetYml
		nacosUrl := ez.GetYmlUrlOrPath(ez.Nacos.Yml.MysqlSet)
		err := ezcfg.GetYmlObj(nacosUrl, &yml)
		if err != nil {
			panic(fmt.Errorf("mysql配置文件获取失败: %s", err.Error()))
		}
		for tag, cfg := range yml.MysqlSet {
			cfg.Tag = tag
			mysqlObjs = append(mysqlObjs, cfg)
		}
	}

	var redisObjs = make([]redis.EZGinRedis, 0)
	if ez.Nacos.Yml.Redis != "" {
		names := strings.Split(ez.Nacos.Yml.Redis, ",")
		for _, name := range names {
			var yml redis.Yml
			nacosUrl := ez.GetYmlUrlOrPath(name)
			err := ezcfg.GetYmlObj(nacosUrl, &yml)
			if err != nil {
				panic(fmt.Errorf("mysql配置文件获取失败: %s", err.Error()))
			}
			redisObjs = append(redisObjs, yml.EZGinRedis)
		}
	}

	if ez.Nacos.Yml.RedisSet != "" {
		var yml redis.SetYml
		nacosUrl := ez.GetYmlUrlOrPath(ez.Nacos.Yml.RedisSet)
		err := ezcfg.GetYmlObj(nacosUrl, &yml)
		if err != nil {
			panic(fmt.Errorf("mysql配置文件获取失败: %s", err.Error()))
		}
		for tag, cfg := range yml.RedisSet {
			cfg.Tag = tag
			redisObjs = append(redisObjs, cfg)
		}
	}

	var kafkaObjs []kafka.EZGinKafka
	if ez.Nacos.Yml.Kafka != "" {
		name := ez.Nacos.Yml.Kafka
		kafkaObjs = make([]kafka.EZGinKafka, 1)
		var yml kafka.Yml
		nacosUrl := ez.GetYmlUrlOrPath(name)
		err := ezcfg.GetYmlObj(nacosUrl, &yml)
		if err != nil {
			panic(fmt.Errorf("mysql配置文件获取失败: %s", err.Error()))
		}
		kafkaObjs[0] = yml.EZGinKafka
	}
	dblite.InitDB(mongoObjs, mysqlObjs, redisObjs, kafkaObjs)
	receiver.initEngine()
}

// initEngine 初始化gin引擎
func (receiver enter) initEngine() {
	ez := config.EZGin()

	logMongoTag := ez.Gin.MwLogs.MongoTag
	logMongoTable := ez.Gin.MwLogs.MongoTable
	if logMongoTable == "" {
		logMongoTable = ez.App.Name + "APIRequestLogs"
	}

	logKafkaTopic := ez.Gin.MwLogs.KafkaTopic
	if logKafkaTopic == "" {
		logKafkaTopic = ez.App.Name
	}

	var logChan chan engineDefine.ReqCtx
	if logMongoTag != "-" || logKafkaTopic != "-" {
		logChan = make(chan engineDefine.ReqCtx, 1000)
		go func(mongoTag, mongoTable, kafkaTopic string) {
			for ctx := range logChan {
				if mongoTag != "-" {
					var db *mgo.Database
					var returnDB func(db *mgo.Database)
					var err error
					if mongoTag != "" {
						db, returnDB, err = ezdb.Mongo(mongoTag)
					} else {
						db, returnDB, err = ezdb.Mongo()
					}

					if err != nil {
						ezlogs.CError("MIDDLEWARE", "写入Mongo日志失败, 获取数据库失败: {}", err.Error())
						return
					}
					ctx.Id = bson.NewObjectId()
					err = db.C(mongoTable).Insert(ctx)
					if err != nil {
						ezlogs.CError("MIDDLEWARE", "写入Mongo日志失败: {}", err.Error())
						returnDB(db)
					}
					returnDB(db)
				}
				if kafkaTopic != "-" && ezdb.Kafka().GetClient() != nil {
					topicList := strings.Split(kafkaTopic, ",")
					for _, topic := range topicList {
						err := ezdb.Kafka().InputMsgForTopic(topic, ctx.ToJson())
						if err != nil {
							ezlogs.CError("MIDDLEWARE", "kafka输入失败: {}", err.Error())
						}
					}
				}
			}
		}(logMongoTag, logMongoTable, logKafkaTopic)
	}
	var ginEngine *gin.Engine
	var recoveryFunc engineDefine.RecoveryFunc
	if StartCfg != nil {
		ginEngine = StartCfg.GinCfg.Engine
		recoveryFunc = engineDefine.RecoveryFunc(StartCfg.GinCfg.RecoveryHandler)
	}
	yml := engine.Yml{
		Mode:         ez.Gin.Mode,
		Middleware:   ez.Gin.Middleware,
		Engine:       ginEngine,
		LogChan:      logChan,
		RecoveryFunc: recoveryFunc,
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
			ezlogs.CWarn("APP", "ListenAndServe:{}:{}", ez.App.Port, listenErr.Error())
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
			ezlogs.CWarn("APP", "ListenAndServeTLS:{}:{}", ez.App.PortSsl, listenErr.Error())
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
