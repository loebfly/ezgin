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
	routineId := receiver.getRoutineId()
	cache.Enter.Table(CacheTable).Add(routineId, lang, 5*time.Minute)
}

// GetRoutineId 获取当前协程Id
func (receiver enter) getRoutineId() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return strconv.FormatUint(n, 10)
}

// GetCurXLang 获取当前语言
func (receiver enter) GetCurXLang() string {
	value, exist := cache.Enter.Table(CacheTable).Get(receiver.getRoutineId())
	if exist {
		return value.(string)
	}
	return ""
}
