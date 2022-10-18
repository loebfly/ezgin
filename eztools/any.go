package eztools

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Any 转换为操作对象
func Any(v interface{}) *object {
	return &object{v}
}

type object struct {
	obj interface{}
}

// OriVal 获取原始值
func (receiver *object) OriVal() interface{} {
	return receiver.obj
}

// ToObject 转换为对象
func (receiver *object) ToObject(obj interface{}) bool {
	if receiver.OriVal() == nil {
		return false
	}
	jsonStr := receiver.ToJson()
	if jsonStr == "" {
		return false
	}
	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		return false
	}
	return true
}

// ToJson 转换为json字符串
func (receiver *object) ToJson() string {
	switch val := receiver.OriVal().(type) {
	case []byte:
		return string(val)
	case string:
		return val
	}
	v := reflect.ValueOf(receiver.OriVal())
	switch v.Kind() {
	case reflect.Ptr, reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
		b, err := json.Marshal(receiver.OriVal())
		if err != nil {
			return ""
		} else {
			result := string(b)
			result = strings.Replace(result, "\\u003c", "<", -1)
			result = strings.Replace(result, "\\u003e", ">", -1)
			result = strings.Replace(result, "\\u0026", "&", -1)
			return result
		}
	default:
		return ""
	}

}

// ToString 转换为字符串
func (receiver *object) ToString() string {
	switch val := receiver.OriVal().(type) {
	case []byte:
		return string(val)
	case string:
		return val
	}
	v := reflect.ValueOf(receiver.OriVal())
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
	return fmt.Sprintf("%v", receiver.OriVal())
}
