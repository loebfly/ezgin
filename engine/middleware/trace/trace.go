package trace

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"runtime"
	"strconv"
)

const (
	HeaderXRequestId = "X-Request-Id"
)

var (
	requestIdCache = make(map[uint64]string) // key routineId, value requestId
)

func (receiver enter) Middleware(c *gin.Context) {
	requestId := c.GetHeader(HeaderXRequestId)
	routineId := receiver.getRoutineId()
	requestIdCache[routineId] = requestId
}

// GetRoutineId 获取当前协程Id
func (receiver enter) getRoutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// DelCurReqId 删除当前缓存的请求Id
func (receiver enter) DelCurReqId() {
	delete(requestIdCache, receiver.getRoutineId())
}

// GetCurReqId 获取当前请求Id
func (receiver enter) GetCurReqId() string {
	return requestIdCache[receiver.getRoutineId()]
}
