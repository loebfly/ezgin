package logs

type enter int

const Enter = enter(0)

func (enter) InitObj(obj Yml) error {
	return config.initObj(obj)
}

func (enter) Debug(format string, args ...interface{}) {
	Logger{category: ""}.Debug(format, args...)
}

func (enter) Info(format string, args ...interface{}) {
	Logger{category: ""}.Info(format, args...)
}

func (enter) Warn(format string, args ...interface{}) {
	Logger{category: ""}.Warn(format, args...)
}

func (enter) Error(format string, args ...interface{}) {
	Logger{category: ""}.Error(format, args...)
}

func (enter) CDebug(category, format string, args ...interface{}) {
	Logger{category: category}.Debug(format, args...)
}

func (enter) CInfo(category, format string, args ...interface{}) {
	Logger{category: category}.Info(format, args...)
}

func (enter) CWarn(category, format string, args ...interface{}) {
	Logger{category: category}.Warn(format, args...)
}

func (enter) CError(category, format string, args ...interface{}) {
	Logger{category: category}.Error(format, args...)
}
