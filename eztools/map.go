package eztools

import (
	"fmt"
	"reflect"
)

func Map(v interface{}) *mapT {
	return &mapT{v}
}

type mapT struct {
	obj interface{}
}

func (receiver *mapT) OriVal() interface{} {
	return receiver.obj
}

// IsExist 判断map是否存在某个key
func (receiver *mapT) IsExist(key interface{}) bool {
	v := reflect.ValueOf(receiver.OriVal())
	switch v.Kind() {
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if k.Interface() == key {
				return true
			}
		}
		return false
	default:
		panic(fmt.Errorf("不是map, 无法判断是否存在: %v", receiver.obj))
	}
}

// IsNil 判断map的值是否为空
func (receiver *mapT) IsNil(key interface{}) bool {
	v := reflect.ValueOf(receiver.OriVal())
	switch v.Kind() {
	case reflect.Map:
		return v.MapIndex(reflect.ValueOf(key)).IsNil()
	default:
		panic(fmt.Errorf("不是map, 无法判断值是否为空: %v", receiver.obj))
	}
}

// Len 获取map的长度
func (receiver *mapT) Len() int {
	v := reflect.ValueOf(receiver.OriVal())
	switch v.Kind() {
	case reflect.Map:
		return v.Len()
	default:
		panic(fmt.Errorf("不是map, 无法获取长度: %v", receiver.obj))
	}
}

// ForEach 遍历map
func (receiver *mapT) ForEach(fn func(key, value interface{})) {
	v := reflect.ValueOf(receiver.OriVal())
	switch v.Kind() {
	case reflect.Map:
		for _, k := range v.MapKeys() {
			fn(k.Interface(), v.MapIndex(k).Interface())
		}
	default:
		panic(fmt.Errorf("不是map, 无法遍历: %v", receiver.obj))
	}
}
