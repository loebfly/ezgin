package ezgin

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/loebfly/ezgin/internal/engine"
)

// DynamicArgsFunc 动态参数函数
/*
	示例:
	DynamicArgsFunc(func(args ...interface{}) {
		// do something
	}).SafeGoExec(func(preRoutineId string) {
		// do something in will call
	}, "hello", "world")
*/
type DynamicArgsFunc func(args ...interface{})

// SafeGoExec 安全的协程调用
func (receiver DynamicArgsFunc) SafeGoExec(will func(preRoutineId string), args ...interface{}) {
	routineId := engine.MWTrace.GetCurRoutineId()
	go func(r DynamicArgsFunc, preRoutineId string, gArgs ...interface{}) {
		defer func() {
			if err := recover(); err != nil {
				goErr := errors.Wrap(err, 2)
				reset := string([]byte{27, 91, 48, 109})
				fmt.Printf("[SafeGo] panic recovered:\n\n%s%s\n\n%s",
					goErr.Error(), goErr.Stack(), reset)
			}
		}()
		if will != nil {
			will(preRoutineId)
		}
		r(gArgs...)
	}(receiver, routineId, args...)
}
