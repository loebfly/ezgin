package eztools

import (
	"encoding/json"
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
	err := json.Unmarshal([]byte(receiver.ToJson()), &obj)
	if err != nil {
		return false
	}
	return true
}

// ToJson 转换为json字符串
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
