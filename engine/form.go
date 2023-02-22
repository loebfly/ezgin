package engine

import (
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	ContentTypeFormUrlEncode = "application/x-www-form-urlencoded"
	ContentTypeFormMultipart = "multipart/form-data"
)

func GetFormParams(c *gin.Context) map[string]string {
	params := make(map[string]string)
	if c.Request.Method == "GET" {
		for k, v := range c.Request.URL.Query() {
			params[k] = v[0]
		}
		return params
	} else if c.Request.Method == "POST" {
		if strings.Contains(c.ContentType(), ContentTypeFormUrlEncode) {
			err := c.Request.ParseForm()
			if err != nil {
				return params
			}
			for k, v := range c.Request.PostForm {
				params[k] = v[0]
			}
			for k, v := range c.Request.URL.Query() {
				params[k] = v[0]
			}
		} else if strings.Contains(c.ContentType(), ContentTypeFormMultipart) {
			err := c.Request.ParseMultipartForm(100 * 1024 * 1024)
			if err != nil {
				return params
			}
			for k, v := range c.Request.MultipartForm.Value {
				params[k] = v[0]
			}
			for k, v := range c.Request.URL.Query() {
				params[k] = v[0]
			}
		}
	}
	return params
}
