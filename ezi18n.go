package ezgin

import (
	"errors"
	"fmt"
	"github.com/loebfly/ezgin/engine"
)

type I18nStringId string

func (receiver I18nStringId) string() string {
	return string(receiver)
}

// Error 返回自身翻译后字符串的error
func (receiver I18nStringId) Error() error {
	return errors.New(I18n.String(receiver.string()))
}

// ErrorWithMsg 返回自身翻译后字符串+:+消息的error
func (receiver I18nStringId) ErrorWithMsg(msg string) error {
	return errors.New(fmt.Sprintf("%s: %s", I18n.String(receiver.string()), msg))
}

// ErrorWithArgs 用于自身翻译后的字符串替换{}参数的error
func (receiver I18nStringId) ErrorWithArgs(args ...interface{}) error {
	return errors.New(I18n.StringFormat(receiver.string(), args...))
}

// ErrorJoinStrId 返回自身翻译后字符串与I18nStringId翻译后的字符串拼接的error
// split 分隔符
// strId 错误ID
func (receiver I18nStringId) ErrorJoinStrId(split string, strId ...I18nStringId) error {
	message := ""
	for _, id := range strId {
		if message == "" {
			message += id.GetTranslate()
			continue
		}
		message += split + id.GetTranslate()
	}
	return receiver.ErrorWithMsg(message)
}

// Result 返回 status 为 1 的 engine.Result, 用于返回成功的结果
// data 数据
func (receiver I18nStringId) Result(data interface{}) engine.Result {
	return engine.Result{
		Status:  1,
		Message: I18n.String(receiver.string()),
		Data:    data,
	}
}

// ResultWithPage 返回 status 为 1 的 engine.Result, 用于返回带分页的成功结果
// data 数据
// page 分页信息
func (receiver I18nStringId) ResultWithPage(data interface{}, page engine.Page) engine.Result {
	return engine.Result{
		Status:  1,
		Message: I18n.String(receiver.string()),
		Data:    data,
		Page:    &page,
	}
}

// ErrorRes 返回自身翻译后字符串用于Message的engine.Result
// status 状态码, 默认为-1
func (receiver I18nStringId) ErrorRes(status ...int) engine.Result {
	targetStatus := -1
	if len(status) > 0 {
		targetStatus = status[0]
	}
	return engine.Result{
		Status:  targetStatus,
		Message: I18n.String(receiver.string()),
	}
}

// ErrorResWithMsg 返回自身翻译后字符串与msg字符串拼接作为message的engine.Result
// msg 分隔符
// status 状态码, 默认为-1
func (receiver I18nStringId) ErrorResWithMsg(msg string, status ...int) engine.Result {
	targetStatus := -1
	if len(status) > 0 {
		targetStatus = status[0]
	}
	return engine.Result{
		Status:  targetStatus,
		Message: I18n.String(receiver.string()) + ": " + msg,
	}
}

// ErrorResWithArgs 返回自身翻译后字符串并替换{}参数用于Message, status=-1的engine.Result
func (receiver I18nStringId) ErrorResWithArgs(args ...interface{}) engine.Result {
	return engine.Result{
		Status:  -1,
		Message: I18n.StringFormat(receiver.string(), args...),
	}
}

// ErrorResWithStatusAndArgs 返回自身翻译后字符串并替换{}参数用于Message的engine.Result
// status 状态码, 默认为-1
// args 参数
func (receiver I18nStringId) ErrorResWithStatusAndArgs(status int, args ...interface{}) engine.Result {
	return engine.Result{
		Status:  status,
		Message: I18n.StringFormat(receiver.string(), args...),
	}
}

// GetTranslate 获取翻译后的字符串
func (receiver I18nStringId) GetTranslate() string {
	return I18n.String(receiver.string())
}

// GetTranslateWithArgs 获取翻译后的字符串, 并替换{}参数
func (receiver I18nStringId) GetTranslateWithArgs(args ...interface{}) string {
	return I18n.StringFormat(receiver.string(), args...)
}
