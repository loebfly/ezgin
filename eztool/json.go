package eztool

import (
	"encoding/json"
	"strings"
)

// ToObj 将字符串转换为对象
func (receiver Str) ToObj(obj interface{}) bool {
	err := json.Unmarshal([]byte(receiver.OriVal()), &obj)
	if err != nil {
		return false
	}
	return true
}

// Obj 转换为操作对象
func Obj(v interface{}) *object {
	return &object{v}
}

type object struct {
	obj interface{}
}

// OriVal 获取原始值
func (receiver *object) OriVal() interface{} {
	return receiver.obj
}

// ToJson 将对象转换为json字符串
func (receiver *object) ToJson() string {
	b, err := json.Marshal(receiver.OriVal())
	if err != nil {
		return "{}"
	} else {
		result := string(b)
		result = strings.Replace(result, "\\u003c", "<", -1)
		result = strings.Replace(result, "\\u003e", ">", -1)
		result = strings.Replace(result, "\\u0026", "&", -1)
		return result
	}
}
