package trace

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/internal/cache"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

const (
	HeaderXRequestId = "X-Request-Id"
	CacheTable       = "Middleware_Trace"
)

func (receiver enter) Middleware(c *gin.Context) {
	requestId := receiver.getRequestId(c)
	routineId := receiver.getRoutineId()
	cache.Enter.Table(CacheTable).Add(routineId, requestId, 5*time.Minute)
}

func (receiver enter) getRequestId(c *gin.Context) string {
	requestId := c.GetHeader(HeaderXRequestId)
	if requestId == "" {
		source := "0123456789abcdef"
		b := []byte(source)
		var result []byte
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < 16; i++ {
			result = append(result, b[r.Intn(len(b))])
		}
		return string(result)
	}
	return requestId
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

// GetCurReqId 获取当前请求Id
func (receiver enter) GetCurReqId() string {
	value, exist := cache.Enter.Table(CacheTable).Get(receiver.getRoutineId())
	if exist {
		return value.(string)
	}
	return ""
}
