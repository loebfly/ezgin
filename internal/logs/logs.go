package logs

import (
	"encoding/json"
	"fmt"
	"github.com/loebfly/ezgin/internal/logs/color"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"

	OutConsole = "console"
	OutFile    = "file"
)

type Logger struct {
	category string
	level    string
}

func (logger Logger) Debug(format string, args ...interface{}) {
	logger.level = LevelDebug
	logger.outPut(format, args...)
}

func (logger Logger) Info(format string, args ...interface{}) {
	logger.level = LevelInfo
	logger.outPut(format, args...)
}

func (logger Logger) Warn(format string, args ...interface{}) {
	logger.level = LevelWarn
	logger.outPut(format, args...)
}

func (logger Logger) Error(format string, args ...interface{}) {
	logger.level = LevelError
	logger.outPut(format, args...)
}

func (logger Logger) outPut(format string, args ...interface{}) {
	// Determine caller func
	createdAt := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	skip := 3
	pc, _, line, ok := runtime.Caller(skip)
	src := "(UNKNOWN)"
	if ok {
		src = runtime.FuncForPC(pc).Name()
		callFunc := src
		if strings.Contains(callFunc, "github.com/loebfly/ezgin") {
			// 取最后一个/后面的字符串
			callFunc = callFunc[strings.LastIndex(callFunc, "/")+1:]
			// 取第一个.前面的字符串
			callFunc = callFunc[:strings.Index(callFunc, ".")]
			// 取倒数第二个.后面的字符串
			callFunc += src[strings.LastIndex(src, "."):]
			callFunc = "ezgin." + callFunc
		}
		src = fmt.Sprintf("[%s:%d]", callFunc, line)
	}

	// args 填充 format
	for _, arg := range args {
		str := logger.argToString(arg)
		format = strings.Replace(format, "{}", str, 1)
	}

	logStack := make([]string, 0)
	logStack = append(logStack, createdAt)
	if logger.category != "" {
		category := "[" + logger.category + "]"
		logStack = append(logStack, category)
	}
	level := "[" + logger.level + "]"
	logStack = append(logStack, level, src, format)

	if len(WillOutputHandlers) > 0 {
		for _, handler := range WillOutputHandlers {
			extraInfo := handler(logger.category, logger.level)
			if extraInfo != nil {
				for k, v := range extraInfo {
					if v > len(logStack)-1 {
						logStack = append(logStack, k)
					} else if v <= 0 {
						logStack = append([]string{k}, logStack...)
					} else {
						logStack = append(logStack[:v], append([]string{k}, logStack[v:]...)...)
					}
				}
			}
		}
	}

	log := strings.Join(logStack, " ")

	if strings.Contains(Config.Logs.Out, OutConsole) {
		logger.outputToConsole(log)
	}
	if strings.Contains(Config.Logs.Out, OutFile) {
		logger.outputToFile(log)
	}
}

func (logger Logger) outputToConsole(content string) {
	logger.getColor().Set()
	fmt.Println(content)
	color.Unset()
}

func (logger Logger) outputToFile(content string) {
	// 写入log文件
	filePath := Config.Logs.File + "."
	filePath += time.Now().Format("2006-01-02")
	filePath += ".log"
	fileDir := path.Dir(filePath)
	err := os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		fmt.Println("log write to file failure:", err)
		return
	}

	var file *os.File
	// 判断文件是否存在
	if _, err = os.Stat(filePath); err == nil {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			fmt.Println("log write to file failure:", err)
			return
		}
	} else {
		// 文件不存在则创建文件, 并写入权限
		file, err = os.Create(filePath)
		if err != nil {
			fmt.Println("log write to file failure:", err)
			return
		}
		// 写入权限
		err = file.Chmod(os.ModePerm)
		if err != nil {
			fmt.Println("log write to file failure:", err)
			return
		}
	}

	defer func(file *os.File) {
		if file != nil {
			_ = file.Close()
		}
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("log write to file failure:", err)
		return
	}
}

// ConvToString 任意类型转换为字符串
func (logger Logger) argToString(iFace interface{}) string {
	switch val := iFace.(type) {
	case []byte:
		return string(val)
	case string:
		return val
	}
	v := reflect.ValueOf(iFace)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Ptr, reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return ""
		}
		return string(b)
	}
	return fmt.Sprintf("%v", iFace)
}

// getColor 根据日志级别获取对应的颜色
func (logger Logger) getColor() *color.Color {
	switch logger.level {
	case LevelDebug:
		return color.New(color.FgYellow)
	case LevelInfo:
		return color.New(color.FgGreen)
	case LevelWarn:
		return color.New(color.FgMagenta)
	case LevelError:
		return color.New(color.FgRed)
	default:
		return color.New(color.Reset)
	}
}
