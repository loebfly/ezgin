package ezgin

import (
	"github.com/go-errors/errors"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/logs"
)

// DynamicArgsFunc 动态参数函数
/*
	示例:
	// 1. 定义一个安全的协程调用
	var safeGo DynamicArgsFunc = func(args ...interface{}) {
		// do something
	}

	// 2. 调用
	safeGo.Exec(func(preRoutineId string) {
		// do something in will call
		engine.CopyPreAllMwDataToCurRoutine(preRoutineId)
	}, "hello", "world")
*/
type DynamicArgsFunc func(args ...interface{})

// SafeGoExec 安全的协程调用
func (receiver DynamicArgsFunc) SafeGoExec(will func(preRoutineId string), args ...interface{}) {
	routineId := engine.MWTrace.GetCurRoutineId()
	go func(r DynamicArgsFunc, preRoutineId string, gArgs ...interface{}) {
		defer func() {
			if err := recover(); err != nil {
				goErr := errors.Wrap(err, 3)
				reset := string([]byte{27, 91, 48, 109})
				logs.Enter.CError("MIDDLEWARE",
					"[Nice Recovery] panic recovered:\n\n{}{}\n\n{}{}",
					goErr.Error(), goErr.Stack(), reset)
			}
		}()
		if will != nil {
			will(preRoutineId)
		}
		r(gArgs...)
	}(receiver, routineId, args...)
}
