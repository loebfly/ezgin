package ezgin

import (
	"fmt"
	"github.com/go-errors/errors"
)

type safeGo struct {
	argsF     func(args ...any)           // 动态参数函数
	goBeforeF func() map[string]any       // 开启协程前的处理函数
	goAfterF  func(params map[string]any) // 开启协程后的处理函数
}

// NewSafeGo 创建一个安全的协程调用
/*
	示例:
	safeGo := ezgin.NewSafeGo(func(args ...any) {
		fmt.Println(args)
	})
	safeGo.SetGoBeforeHandler(func() map[string]any {
		return map[string]any{"preRoutineId": ezgin.Engine.GetMWTraceCurRoutineId()}
	})
	safeGo.SetGoAfterHandler(func(params map[string]any) {
		ezgin.Engine.CopyMWTracePreHeaderToCurRoutine(params["preRoutineId"].(string))
	})
	safeGo.Run("hello", "world")
*/
func NewSafeGo(argsF func(args ...any)) *safeGo {
	return &safeGo{
		argsF: argsF,
	}
}

// SetGoBeforeHandler 设置协程前的处理函数
func (receiver *safeGo) SetGoBeforeHandler(goBeforeF func() map[string]any) *safeGo {
	receiver.goBeforeF = goBeforeF
	return receiver
}

// SetGoAfterHandler 设置协程后的处理函数
func (receiver *safeGo) SetGoAfterHandler(callBeforeF func(params map[string]any)) *safeGo {
	receiver.goAfterF = callBeforeF
	return receiver
}

// Run 运行
func (receiver *safeGo) Run(args ...any) {
	var goBeforeParams map[string]any
	if receiver.goBeforeF != nil {
		goBeforeParams = receiver.goBeforeF()
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				goErr := errors.Wrap(err, 2)
				reset := string([]byte{27, 91, 48, 109})
				fmt.Printf("[SafeGo] panic recovered:\n\n%s%s\n\n%s",
					goErr.Error(), goErr.Stack(), reset)
			}
		}()
		if receiver.goAfterF != nil {
			receiver.goAfterF(goBeforeParams)
		}
		if receiver.argsF != nil {
			receiver.argsF(args...)
		}
	}()
}
