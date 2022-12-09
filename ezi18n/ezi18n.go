package ezi18n

import (
	"errors"
	"fmt"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/internal/i18n"
)

type StringId string

func (receiver StringId) string() string {
	return string(receiver)
}

// Error 返回自身翻译后字符串的error
func (receiver StringId) Error() error {
	return errors.New(i18n.Enter.String(receiver.string()))
}

// ErrorWithMsg 返回自身翻译后字符串+:+消息的error
func (receiver StringId) ErrorWithMsg(msg string) error {
	return errors.New(fmt.Sprintf("%s: %s", i18n.Enter.String(receiver.string()), msg))
}

// ErrorWithArgs 用于自身翻译后的字符串替换{}参数的error
func (receiver StringId) ErrorWithArgs(args ...any) error {
	return errors.New(i18n.Enter.StringFormat(receiver.string(), args...))
}

// ErrorJoinStrId 返回自身翻译后字符串与I18nStringId翻译后的字符串拼接的error
// split 分隔符
// strId 错误ID
func (receiver StringId) ErrorJoinStrId(split string, strId ...StringId) error {
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
func (receiver StringId) Result(data any) engine.Result[any] {
	return engine.Result[any]{
		Status:  1,
		Message: i18n.Enter.String(receiver.string()),
		Data:    data,
	}
}

// ResultWithPage 返回 status 为 1 的 engine.Result, 用于返回带分页的成功结果
// data 数据
// page 分页信息
func (receiver StringId) ResultWithPage(data any, page engine.Page) engine.Result[any] {
	return engine.Result[any]{
		Status:  1,
		Message: i18n.Enter.String(receiver.string()),
		Data:    data,
		Page:    &page,
	}
}

// ErrorRes 返回自身翻译后字符串用于Message的engine.Result
// status 状态码, 默认为-1
func (receiver StringId) ErrorRes(status ...int) engine.Result[any] {
	targetStatus := -1
	if len(status) > 0 {
		targetStatus = status[0]
	}
	return engine.Result[any]{
		Status:  targetStatus,
		Message: i18n.Enter.String(receiver.string()),
	}
}

// ErrorResWithMsg 返回自身翻译后字符串与msg字符串拼接作为message的engine.Result
// msg 分隔符
// status 状态码, 默认为-1
func (receiver StringId) ErrorResWithMsg(msg string, status ...int) engine.Result[any] {
	targetStatus := -1
	if len(status) > 0 {
		targetStatus = status[0]
	}
	return engine.Result[any]{
		Status:  targetStatus,
		Message: i18n.Enter.String(receiver.string()) + ": " + msg,
	}
}

// ErrorResWithArgs 返回自身翻译后字符串并替换{}参数用于Message, status=-1的engine.Result
func (receiver StringId) ErrorResWithArgs(args ...any) engine.Result[any] {
	return engine.Result[any]{
		Status:  -1,
		Message: i18n.Enter.StringFormat(receiver.string(), args...),
	}
}

// ErrorResWithStatusAndArgs 返回自身翻译后字符串并替换{}参数用于Message的engine.Result
// status 状态码, 默认为-1
// args 参数
func (receiver StringId) ErrorResWithStatusAndArgs(status int, args ...any) engine.Result[any] {
	return engine.Result[any]{
		Status:  status,
		Message: i18n.Enter.StringFormat(receiver.string(), args...),
	}
}

// CheckRes 检查参数是否丢失, 丢失则返回 status为1003 错误, 否则返回status为1的engine.Result
/*
	使用说明:
		receiver 的翻译字符串需要带{},如果丢失用于key替换后作为Message，例如:参数{}丢失
		params 参数列表
		key 参数名
*/
func (receiver StringId) CheckRes(params map[string]string, key ...string) engine.Result[any] {
	for _, param := range key {
		val := params[param]
		if val == "" {
			return receiver.ErrorResWithArgs(param, 1003)
		}
	}
	return receiver.Result(nil)
}

// GetTranslate 获取翻译后的字符串
func (receiver StringId) GetTranslate() string {
	return i18n.Enter.String(receiver.string())
}

// GetTranslateWithArgs 获取翻译后的字符串, 并替换{}参数
func (receiver StringId) GetTranslateWithArgs(args ...any) string {
	return i18n.Enter.StringFormat(receiver.string(), args...)
}

// GetTranslateByLang 根据语言获取翻译后的字符串
func (receiver StringId) GetTranslateByLang(lang, messageId string, args ...any) string {
	return i18n.Enter.StringByLang(lang, messageId, args...)
}

// GetTranslateWithArgsByLang 根据语言获取翻译后的字符串, 并替换{}参数
func (receiver StringId) GetTranslateWithArgsByLang(lang, messageId string, args ...any) string {
	return i18n.Enter.StringFormatByLang(lang, messageId, args...)
}
