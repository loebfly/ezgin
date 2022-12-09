package ginrecover

import (
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/loebfly/ezgin/ezlogs"
	"net/http/httputil"
)

func (receiver enter) Middleware(f func(c *gin.Context, err any)) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				goErr := errors.Wrap(err, 3)
				reset := string([]byte{27, 91, 48, 109})
				ezlogs.CError("MIDDLEWARE",
					"[Nice Recovery] panic recovered:\n\n{}{}\n\n{}{}",
					httpRequest, goErr.Error(), goErr.Stack(), reset)
				if f != nil {
					f(c, err)
				}
			}
		}()
		c.Next() // execute all the handlers
	}
}
