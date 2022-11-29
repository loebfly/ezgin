package xlang

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/internal/cache"
	"runtime"
	"strconv"
	"time"
)

const (
	CacheTable = "Middleware_XLang"
)

func (receiver enter) Middleware(c *gin.Context) {
	lang := c.GetHeader("X-Lang")
	if lang == "" {
		lang = "zh-cn"
	}
	routineId := receiver.GetCurRoutineId()
	cache.Enter.Table(CacheTable).Add(routineId, lang, 5*time.Minute)
}

// GetCurRoutineId 获取当前协程Id
func (receiver enter) GetCurRoutineId() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return strconv.FormatUint(n, 10)
}

// GetCurXLang 获取当前语言
func (receiver enter) GetCurXLang() string {
	value, exist := cache.Enter.Table(CacheTable).Get(receiver.GetCurRoutineId())
	if exist {
		return value.(string)
	}
	return ""
}

// CopyPreXLangToCurRoutine 复制上一个请求的语言
func (receiver enter) CopyPreXLangToCurRoutine(preRoutineId string) {
	value, exist := cache.Enter.Table(CacheTable).Get(preRoutineId)
	if exist {
		cache.Enter.Table(CacheTable).Add(receiver.GetCurRoutineId(), value, 5*time.Minute)
	}
}
