package reqlogs

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type respWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w respWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w respWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
