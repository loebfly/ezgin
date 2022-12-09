package i18n

import (
	"github.com/loebfly/ezgin/internal/engine"
)

type enter int

const Enter = enter(0)

func InitObj(obj Yml) {
	config.initObj(obj)
	ctl.initXlang()
}

func (enter) String(messageId string) string {
	lang := engine.MWTrace.GetCurXLang()
	return ctl.getString(lang, messageId)
}

func (enter) StringFormat(messageId string, args ...any) string {
	lang := engine.MWTrace.GetCurXLang()
	return ctl.getString(lang, messageId, args...)
}

func (enter) StringByLang(lang, messageId string, args ...any) string {
	return ctl.getString(lang, messageId, args...)
}

func (enter) StringFormatByLang(lang, messageId string, args ...any) string {
	return ctl.getString(lang, messageId, args...)
}
