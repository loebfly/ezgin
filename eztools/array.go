package eztools

import (
	"fmt"
	"reflect"
)

// Array 转换为可操作对象
/*
	使用说明:
		1. IsExist(value interface{}) 判断切片是否存在某个值
		2. IsNil(index int) 判断切片是否为空
		3. Len() 获取切片的长度
		4. ForEach(fn func(index int)) 遍历切片
*/
func Array(v interface{}) *array {
	return &array{v}
}

type array struct {
	obj interface{}
}

// OriVal 获取原始值
func (receiver *array) OriVal() interface{} {
	return receiver.obj
}

// IsExist 判断切片是否存在某个值
func (receiver *array) IsExist(value interface{}) bool {
	v := reflect.ValueOf(receiver.OriVal())
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() == value {
				return true
			}
		}
		return false
	default:
		panic(fmt.Errorf("不是切片或数组, 无法判断是否存在: %v", receiver.obj))
	}
}

// IsNil 判断切片是否为空
func (receiver *array) IsNil(index int) bool {
	v := reflect.ValueOf(receiver.OriVal())
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		return v.Index(index).IsNil()
	default:
		panic(fmt.Errorf("不是切片或数组, 无法判断是否为空: %v", receiver.obj))
	}

}

// Len 获取切片的长度
func (receiver *array) Len() int {
	v := reflect.ValueOf(receiver.OriVal())
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		return v.Len()
	default:
		panic(fmt.Errorf("不是切片或数组, 无法获取长度: %v", receiver.obj))
	}
}

// ForEach 遍历切片
func (receiver *array) ForEach(fn func(index int)) {
	v := reflect.ValueOf(receiver.OriVal())
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			fn(i)
		}
	default:
		panic(fmt.Errorf("不是切片或数组, 无法遍历: %v", receiver.obj))
	}
}
