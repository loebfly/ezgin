package reqlogs

type enter int

const Enter = enter(0)

var reqCtxChan chan ReqCtx

var logWriter *mongoWriter

func (enter) OpenMongoWriter(mongoTag, table string) {
	if mongoTag == "-" {
		return
	}
	logWriter = &mongoWriter{
		MongoTag: mongoTag,
		Table:    table,
	}
	reqCtxChan = make(chan ReqCtx, 100)
	go Enter.HandleReqCtxChan()
}

func (enter) HandleReqCtxChan() {
	for ctx := range reqCtxChan {
		logWriter.Write(ctx)
	}
}
