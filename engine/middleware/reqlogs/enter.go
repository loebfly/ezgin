package reqlogs

import "github.com/loebfly/ezgin/logs"

type enter int

const Enter = enter(0)

var logChan = make(chan ReqCtx, 1000)

func (enter) HandleLogChan() {
	for ctx := range logChan {
		logs.Enter.CDebug("access", "{}", ctx)
	}
}
