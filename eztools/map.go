package eztools

import "reflect"

func Map(v interface{}) *mapT {
	return &mapT{v}
}

type mapT struct {
	mapObj interface{}
}

// OriVal 获取原始值
func (receiver *mapT) OriVal() interface{} {
	return receiver.mapObj
}

func (receiver *mapT) ToStr() map[string]string {
	result := make(map[string]string)
	for k, v := range receiver.OriVal().(map[string]interface{}) {
		result[k] = v.(string)
	}
	return result
}

// IsExist 判断key是否存在
func (receiver *mapT) IsExist(k string) bool {
	v := reflect.ValueOf(receiver.mapObj)
	switch v.Kind() {
	case reflect.Map:
		val := v.MapIndex(reflect.ValueOf(k)).Interface()
		return val != nil
	default:
		return false
	}
}

// IsEmptyV 判断值是否为空
func (receiver *mapT) IsEmptyV(k string) bool {
	v := reflect.ValueOf(receiver.mapObj)
	switch v.Kind() {
	case reflect.Map:
		val := v.MapIndex(reflect.ValueOf(k)).Interface()
		return val == "" || val == nil
	default:
		return false
	}
}
