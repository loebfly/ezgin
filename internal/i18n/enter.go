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

func (enter) GetString(lang, messageId string, args ...interface{}) string {
	return ctl.getString(lang, messageId, args...)
}

func (enter) GetCurLangString(messageId string, args ...interface{}) string {
	lang := engine.Enter.GetCurXLang()
	return ctl.getString(lang, messageId, args...)
}
