package logs

type enter int

const Enter = enter(0)

func InitObj(obj Yml) {
	Config.initObj(obj)
}

func (enter) Debug(format string, args ...any) {
	Logger{category: ""}.Debug(format, args...)
}

func (enter) Info(format string, args ...any) {
	Logger{category: ""}.Info(format, args...)
}

func (enter) Warn(format string, args ...any) {
	Logger{category: ""}.Warn(format, args...)
}

func (enter) Error(format string, args ...any) {
	Logger{category: ""}.Error(format, args...)
}

func (enter) CDebug(category, format string, args ...any) {
	Logger{category: category}.Debug(format, args...)
}

func (enter) CInfo(category, format string, args ...any) {
	Logger{category: category}.Info(format, args...)
}

func (enter) CWarn(category, format string, args ...any) {
	Logger{category: category}.Warn(format, args...)
}

func (enter) CError(category, format string, args ...any) {
	Logger{category: category}.Error(format, args...)
}
