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
	HeaderXRequestIdKey = "X-Request-Id"
	HeaderXRealIPKey    = "X-Real-IP"
	HeaderXUserAgentKey = "X-User-Agent"
	HeaderXLangKey      = "X-Lang"

	XHeaderTable = "XHeaderTable"

	CacheDuration = 5 * time.Minute
)

func (receiver enter) Middleware(c *gin.Context) {
	routineId := receiver.GetCurRoutineId()
	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	if headers[HeaderXRequestIdKey] == "" {
		// 如果请求头中没有X-Request-Id，则生成一个
		headers[HeaderXRequestIdKey] = receiver.newRequestId()
	}

	if headers[HeaderXRealIPKey] == "" {
		// 如果请求头中没有X-Real-IP，则从X-Forwarded-For中获取,如果X-Forwarded-For也没有，则从RemoteAddr中获取
		headerXForwardedForKey := "X-Forwarded-For"
		if headers[headerXForwardedForKey] != "" {
			headers[HeaderXRealIPKey] = headers[headerXForwardedForKey]
		} else {
			headers[HeaderXRealIPKey] = c.ClientIP()
		}
	}

	if headers[HeaderXUserAgentKey] == "" {
		// 如果请求头中没有X-User-Agent，则从RemoteAddr中获取
		headers[HeaderXUserAgentKey] = c.Request.UserAgent()
	}

	if headers[HeaderXLangKey] == "" {
		// 如果请求头中没有X-Lang, 则默认为zh-cn
		headers[HeaderXLangKey] = "zh-cn"
	}

	// 将请求头信息存入缓存
	cache.Enter.Table(XHeaderTable).Add(routineId, headers, CacheDuration)
}

func (receiver enter) newRequestId() string {
	source := "0123456789abcdef"
	b := []byte(source)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
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
	return receiver.GetCurHeader()[HeaderXRequestIdKey]
}

// GetCurClientIP 获取当前客户端IP
func (receiver enter) GetCurClientIP() string {
	return receiver.GetCurHeader()[HeaderXRealIPKey]
}

// GetCurUserAgent 获取当前客户端UserAgent
func (receiver enter) GetCurUserAgent() string {
	return receiver.GetCurHeader()[HeaderXUserAgentKey]
}

// GetCurXLang 获取当前语言
func (receiver enter) GetCurXLang() string {
	return receiver.GetCurHeader()[HeaderXLangKey]
}

func (receiver enter) GetCurHeader() map[string]string {
	value, exist := cache.Enter.Table(XHeaderTable).Get(receiver.GetCurRoutineId())
	if exist {
		var headers = make(map[string]string)
		for k, v := range value.(map[string]string) {
			headers[k] = v
		}
		return headers
	}
	return map[string]string{}
}

// CopyPreHeaderToCurRoutine 复制上一个协程的所有信息到当前协程
func (receiver enter) CopyPreHeaderToCurRoutine(preRoutineId string) {
	value, exist := cache.Enter.Table(XHeaderTable).Get(preRoutineId)
	if exist {
		cache.Enter.Table(XHeaderTable).Add(receiver.GetCurRoutineId(), value, 5*time.Minute)
	}
}
