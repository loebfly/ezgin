package reqlogs

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/internal/dblite"
	"github.com/loebfly/ezgin/internal/logs"
	"gopkg.in/mgo.v2/bson"
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

type mongoWriter struct {
	MongoTag string
	Table    string
}

func (w *mongoWriter) Write(ctx ReqCtx) {
	db, returnDB, err := dblite.Enter.Mongo(w.MongoTag)
	if err != nil {
		logs.Enter.CError("MIDDLEWARE", "写入日志失败, 获取数据库失败: %s", err.Error())
		return
	}
	ctx.Id = bson.NewObjectId()
	err = db.C(w.Table).Insert(ctx)
	if err != nil {
		logs.Enter.CError("MIDDLEWARE", "写入日志失败: %s", err.Error())
		returnDB(db)
	}
	returnDB(db)
}
