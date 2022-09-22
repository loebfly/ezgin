package ezgin

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/app"
)

const (
	App = app.Enter(0)
)

func Start(ymlPath string, engine gin.Engine) {

}
