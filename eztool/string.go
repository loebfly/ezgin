package eztool

import (
	"strconv"
)

type Str string

// OriVal 获取原始值
func (receiver Str) OriVal() string {
	return string(receiver)
}

// ToBool 数字字符串转换为bool, 当转换失败时有默认值返回第一个默认值否则返回false
func (receiver Str) ToBool(defaultBool ...bool) bool {
	num, err := strconv.ParseBool(receiver.OriVal())
	if err != nil {
		if len(defaultBool) > 0 {
			return defaultBool[0]
		} else {
			return false
		}
	}
	return num
}

// ToUInt 数字字符串转换为uint, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToUInt(defaultUInt ...uint) uint {
	num, err := strconv.ParseUint(receiver.OriVal(), 10, 64)
	if err != nil {
		if len(defaultUInt) > 0 {
			return defaultUInt[0]
		} else {
			return 0
		}
	}
	return uint(num)
}

// ToUInt8 数字字符串转换为uint8, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToUInt8(defaultUInt8 ...uint8) uint8 {
	num, err := strconv.ParseUint(receiver.OriVal(), 10, 8)
	if err != nil {
		if len(defaultUInt8) > 0 {
			return defaultUInt8[0]
		} else {
			return 0
		}
	}
	return uint8(num)
}

// ToUInt16 数字字符串转换为uint16, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToUInt16(defaultUInt16 ...uint16) uint16 {
	num, err := strconv.ParseUint(receiver.OriVal(), 10, 16)
	if err != nil {
		if len(defaultUInt16) > 0 {
			return defaultUInt16[0]
		} else {
			return 0
		}
	}
	return uint16(num)
}

// ToUInt32 数字字符串转换为uint32, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToUInt32(defaultUInt32 ...uint32) uint32 {
	num, err := strconv.ParseUint(receiver.OriVal(), 10, 32)
	if err != nil {
		if len(defaultUInt32) > 0 {
			return defaultUInt32[0]
		} else {
			return 0
		}
	}
	return uint32(num)
}

// ToUInt64 数字字符串转换为uint64, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToUInt64(defaultUInt64 ...uint64) uint64 {
	num, err := strconv.ParseUint(receiver.OriVal(), 10, 64)
	if err != nil {
		if len(defaultUInt64) > 0 {
			return defaultUInt64[0]
		} else {
			return 0
		}
	}
	return num
}

// ToInt 数字字符串转换为int, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToInt(defaultInt ...int) int {
	num, err := strconv.Atoi(receiver.OriVal())
	if err != nil {
		if len(defaultInt) > 0 {
			return defaultInt[0]
		} else {
			return 0
		}
	}
	return num
}

// ToInt8 数字字符串转换为int8, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToInt8(defaultInt8 ...int8) int8 {
	num, err := strconv.ParseInt(receiver.OriVal(), 10, 8)
	if err != nil {
		if len(defaultInt8) > 0 {
			return defaultInt8[0]
		} else {
			return 0
		}
	}
	return int8(num)
}

// ToInt32 数字字符串转换为int32, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToInt32(defaultInt32 ...int32) int32 {
	num, err := strconv.ParseInt(receiver.OriVal(), 10, 32)
	if err != nil {
		if len(defaultInt32) > 0 {
			return defaultInt32[0]
		} else {
			return 0
		}
	}
	return int32(num)
}

// ToInt64 数字字符串转换为int64, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToInt64(defaultInt64 ...int64) int64 {
	num, err := strconv.ParseInt(receiver.OriVal(), 10, 64)
	if err != nil {
		if len(defaultInt64) > 0 {
			return defaultInt64[0]
		} else {
			return 0
		}
	}
	return num
}

// ToFloat32 数字字符串转换为float32, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToFloat32(defaultFloat32 ...float32) float32 {
	num, err := strconv.ParseFloat(receiver.OriVal(), 32)
	if err != nil {
		if len(defaultFloat32) > 0 {
			return defaultFloat32[0]
		} else {
			return 0
		}
	}
	return float32(num)
}

// ToFloat64 数字字符串转换为float64, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToFloat64(defaultFloat64 ...float64) float64 {
	num, err := strconv.ParseFloat(receiver.OriVal(), 64)
	if err != nil {
		if len(defaultFloat64) > 0 {
			return defaultFloat64[0]
		} else {
			return 0
		}
	}
	return num
}

// ToComplex64 数字字符串转换为complex64, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToComplex64(defaultComplex64 ...complex64) complex64 {
	num, err := strconv.ParseComplex(receiver.OriVal(), 64)
	if err != nil {
		if len(defaultComplex64) > 0 {
			return defaultComplex64[0]
		} else {
			return 0
		}
	}
	return complex64(num)
}

// ToComplex128 数字字符串转换为complex128, 当转换失败时有默认值返回第一个默认值否则返回0
func (receiver Str) ToComplex128(defaultComplex128 ...complex128) complex128 {
	num, err := strconv.ParseComplex(receiver.OriVal(), 128)
	if err != nil {
		if len(defaultComplex128) > 0 {
			return defaultComplex128[0]
		} else {
			return 0
		}
	}
	return num
}
