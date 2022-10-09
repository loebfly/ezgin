package i18n

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/internal/cache"
	"github.com/loebfly/ezgin/internal/call"
	"github.com/loebfly/ezgin/internal/logs"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var ctl = new(control)

type control struct{}

const (
	CacheTableXLangVersion = "x-lang-version"
	CacheKeyXLangVersion   = "version"

	CacheTableXLang = "x-lang"
)

func (ctl *control) initXlang() {
	ctl.refreshCache()
	ticker := time.NewTicker(time.Duration(config.I18n.Duration) * time.Second)
	go func() {
		for range ticker.C {
			ctl.refreshCache()
		}
	}()
}

// refreshCache 刷新缓存
func (ctl *control) refreshCache() {
	for _, appName := range config.I18n.AppName {
		lastVersion, err := ctl.getAppXLangLastVersion(appName)
		if err != nil {
			logs.Enter.CError("I18N", "获取应用多语言版本号失败", err)
			continue
		}
		value, isExist := cache.Enter.Table(CacheTableXLangVersion).Get(appName + "-" + CacheKeyXLangVersion)
		if isExist {
			cacheVersion := value.(string)
			if lastVersion != cacheVersion {
				ctl.cacheAppXlangData(appName, lastVersion)
			}
		} else {
			ctl.cacheAppXlangData(appName, lastVersion)
		}

	}
}

func (ctl *control) cacheAppXlangData(appName, version string) {
	data, err := ctl.getLastAppXlangData(appName)
	if err != nil {
		logs.Enter.CError("I18N", "获取应用多语言数据失败", err)
		return
	}

	for key, value := range data {
		cache.Enter.Table(CacheTableXLang).Add(key, value, 0)
	}

	cache.Enter.Table(CacheTableXLangVersion).Add(appName+"-"+CacheKeyXLangVersion, version, 0)
}

// getAppXLangLastVersion 获取应用最新的多语言版本号
func (ctl *control) getAppXLangLastVersion(appName string) (string, error) {
	respStr, err := call.Enter.FormPost(config.I18n.ServerName, config.I18n.CheckUri, map[string]string{
		"appName": appName,
	})
	if err != nil {
		return "", err
	}
	var result engine.Result
	err = json.Unmarshal([]byte(respStr), &result)
	if err != nil {
		return "", err
	}
	if result.Status != 1 {
		return "", errors.New(result.Message)
	}
	data := make(map[string]string)
	respData, err := json.Marshal(result.Data)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(respData, &data)
	return data["version"], nil
}

// getLastAppXlangData 获取应用最新的多语言数据
func (ctl *control) getLastAppXlangData(appName string) (map[string]string, error) {
	respStr, err := call.Enter.FormPost(config.I18n.ServerName, config.I18n.QueryUri, map[string]string{
		"appName": appName,
	})
	if err != nil {
		return nil, err
	}
	var result engine.Result
	err = json.Unmarshal([]byte(respStr), &result)
	if err != nil {
		return nil, err
	}
	if result.Status != 1 {
		return nil, errors.New(result.Message)
	}
	data := make(map[string]string)
	respData, err := json.Marshal(result.Data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respData, &data)
	return data, nil
}

// getString 获取多语言字符串
func (ctl *control) getString(lang, strId string, args ...interface{}) string {
	value, isExist := cache.Enter.Table(CacheTableXLang).Get(fmt.Sprintf("%s:%s", strId, lang))
	if isExist {
		format := value.(string)
		// args 填充 format
		for _, arg := range args {
			str := ctl.argToString(arg)
			format = strings.Replace(format, "{}", str, 1)
		}
		return format
	}
	return strId
}

// ConvToString 任意类型转换为字符串
func (ctl *control) argToString(iFace interface{}) string {
	switch val := iFace.(type) {
	case []byte:
		return string(val)
	case string:
		return val
	}
	v := reflect.ValueOf(iFace)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Ptr, reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return ""
		}
		return string(b)
	}
	return fmt.Sprintf("%v", iFace)
}
