package ezgin

import (
	"github.com/go-errors/errors"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/logs"
)

// GoFunc 安全的协程调用
type GoFunc func(args ...interface{})

func (receiver GoFunc) SafeCall(args ...interface{}) {
	routineId := engine.MWTrace.GetCurRoutineId()
	go func(r GoFunc, preRoutineId string, gArgs ...interface{}) {
		defer func() {
			if err := recover(); err != nil {
				goErr := errors.Wrap(err, 3)
				reset := string([]byte{27, 91, 48, 109})
				logs.Enter.CError("MIDDLEWARE",
					"[Nice Recovery] panic recovered:\n\n{}{}\n\n{}{}",
					goErr.Error(), goErr.Stack(), reset)
			}
		}()
		engine.MWTrace.CopyPreAllToCurRoutine(preRoutineId)
		engine.MWXLang.CopyPreXLangToCurRoutine(preRoutineId)
		r(gArgs...)
	}(receiver, routineId, args...)
}
