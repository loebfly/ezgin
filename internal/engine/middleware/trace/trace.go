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
	HeaderXRequestId    = "X-Request-Id"
	HeaderXRealIP       = "X-Real-IP"
	HeaderXForwardedFor = "X-Forwarded-For"
	HeaderXUserAgent    = "X-User-Agent"

	XRequestIdTable = "XRequestIdTable"
	XClientIPTable  = "XClientIPTable"
	XUserAgentTable = "XUserAgentTable"

	CacheDuration = 5 * time.Minute
)

func (receiver enter) Middleware(c *gin.Context) {
	requestId := receiver.newRequestId(c)
	routineId := receiver.GetCurRoutineId()
	cache.Enter.Table(XRequestIdTable).Add(routineId, requestId, CacheDuration)

	clientIP := c.ClientIP()
	if c.GetHeader(HeaderXRealIP) != "" {
		clientIP = c.GetHeader(HeaderXRealIP)
	}
	if c.GetHeader(HeaderXForwardedFor) != "" {
		clientIP = c.GetHeader(HeaderXForwardedFor)
	}
	cache.Enter.Table(XClientIPTable).Add(routineId, clientIP, CacheDuration)

	userAgent := c.GetHeader(HeaderXUserAgent)
	if userAgent == "" {
		userAgent = c.Request.UserAgent()
	}
	cache.Enter.Table(XUserAgentTable).Add(routineId, userAgent, CacheDuration)
}

func (receiver enter) newRequestId(c *gin.Context) string {
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

// GetCurRoutineId 获取当前协程Id
func (receiver enter) GetCurRoutineId() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return strconv.FormatUint(n, 10)
}

// GetCurReqId 获取当前请求Id
func (receiver enter) GetCurReqId() string {
	value, exist := cache.Enter.Table(XRequestIdTable).Get(receiver.GetCurRoutineId())
	if exist {
		return value.(string)
	}
	return ""
}

// CopyPreReqIdToCurRoutine 复制上一个协程的请求Id到当前协程
func (receiver enter) CopyPreReqIdToCurRoutine(preRoutineId string) {
	value, exist := cache.Enter.Table(XRequestIdTable).Get(preRoutineId)
	if exist {
		cache.Enter.Table(XRequestIdTable).Add(receiver.GetCurRoutineId(), value, CacheDuration)
	}
}

// GetCurClientIP 获取当前客户端IP
func (receiver enter) GetCurClientIP() string {
	value, exist := cache.Enter.Table(XClientIPTable).Get(receiver.GetCurRoutineId())
	if exist {
		return value.(string)
	}
	return ""
}

// CopyPreClientIPToCurRoutine 复制上一个协程的客户端IP到当前协程
func (receiver enter) CopyPreClientIPToCurRoutine(preRoutineId string) {
	value, exist := cache.Enter.Table(XClientIPTable).Get(preRoutineId)
	if exist {
		cache.Enter.Table(XClientIPTable).Add(receiver.GetCurRoutineId(), value, CacheDuration)
	}
}

// GetCurUserAgent 获取当前客户端UserAgent
func (receiver enter) GetCurUserAgent() string {
	value, exist := cache.Enter.Table(XUserAgentTable).Get(receiver.GetCurRoutineId())
	if exist {
		return value.(string)
	}
	return ""
}

// CopyPreUserAgentToCurRoutine 复制上一个协程的客户端UserAgent到当前协程
func (receiver enter) CopyPreUserAgentToCurRoutine(preRoutineId string) {
	value, exist := cache.Enter.Table(XUserAgentTable).Get(preRoutineId)
	if exist {
		cache.Enter.Table(XUserAgentTable).Add(receiver.GetCurRoutineId(), value, CacheDuration)
	}
}

// CopyPreAllToCurRoutine 复制上一个协程的所有信息到当前协程
func (receiver enter) CopyPreAllToCurRoutine(preRoutineId string) {
	receiver.CopyPreReqIdToCurRoutine(preRoutineId)
	receiver.CopyPreClientIPToCurRoutine(preRoutineId)
	receiver.CopyPreUserAgentToCurRoutine(preRoutineId)
}
