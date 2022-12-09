package ezgo

import (
	"fmt"
	"github.com/go-errors/errors"
)

type safeGo[P any] struct {
	argsF     func(args ...any) // 动态参数函数
	goBeforeF func() P          // 开启协程前的处理函数
	goAfterF  func(params P)    // 开启协程后的处理函数
}

// New 创建一个安全的协程调用
func New[P any](argsF func(args ...any)) *safeGo[P] {
	return &safeGo[P]{
		argsF: argsF,
	}
}

// SetGoBeforeHandler 设置协程前的处理函数
func (receiver *safeGo[P]) SetGoBeforeHandler(goBeforeF func() P) *safeGo[P] {
	receiver.goBeforeF = goBeforeF
	return receiver
}

// SetGoAfterHandler 设置协程后的处理函数
func (receiver *safeGo[P]) SetGoAfterHandler(callBeforeF func(params P)) *safeGo[P] {
	receiver.goAfterF = callBeforeF
	return receiver
}

// Run 运行
func (receiver *safeGo[P]) Run(args ...any) {
	var goBeforeParams P
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
