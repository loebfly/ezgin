package reqlogs

type enter int

const Enter = enter(0)

var logChan chan ReqCtx

func (enter) SetLogChan(c chan ReqCtx) {
	logChan = c
}
