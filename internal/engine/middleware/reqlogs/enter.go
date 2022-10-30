package reqlogs

import engineDefine "github.com/loebfly/ezgin/engine"

type enter int

const Enter = enter(0)

var logChan chan engineDefine.ReqCtx

func (enter) SetLogChan(c chan engineDefine.ReqCtx) {
	logChan = c
}
