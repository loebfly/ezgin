package ezlogs

import "github.com/loebfly/ezgin/internal/logs"

func Debug(format string, args ...any) {
	logs.Logger{Category: ""}.Debug(format, args...)
}

func Info(format string, args ...any) {
	logs.Logger{Category: ""}.Info(format, args...)
}

func Warn(format string, args ...any) {
	logs.Logger{Category: ""}.Warn(format, args...)
}

func Error(format string, args ...any) {
	logs.Logger{Category: ""}.Error(format, args...)
}

func CDebug(category, format string, args ...any) {
	logs.Logger{Category: category}.Debug(format, args...)
}

func CInfo(category, format string, args ...any) {
	logs.Logger{Category: category}.Info(format, args...)
}

func CWarn(category, format string, args ...any) {
	logs.Logger{Category: category}.Warn(format, args...)
}

func CError(category, format string, args ...any) {
	logs.Logger{Category: category}.Error(format, args...)
}
