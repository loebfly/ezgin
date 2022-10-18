package eztools

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Str 字符串
/*
	使用说明:
		1. 该类型可便于字符串转化为各种数字类型
		2. 该类型提供了一些常见的正则验证方法
		3. 该类型提供了快速截取字符串的方法
*/
type Str string

// OriVal 获取原始值
func (receiver Str) OriVal() string {
	return string(receiver)
}

/******* - Convert - *******/

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

/****************************/

/******* - Verify - *******/

// IsCNMobile 判断是否是中国大陆手机号
func (receiver Str) IsCNMobile() bool {
	reg, _ := regexp.Compile(`^1[3456789]\d{9}$`)
	return reg.MatchString(receiver.OriVal())
}

// IsEmail 判断是否是邮箱
func (receiver Str) IsEmail() bool {
	reg, _ := regexp.Compile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	return reg.MatchString(receiver.OriVal())
}

// IsIDCard 判断是否是身份证号
func (receiver Str) IsIDCard() bool {
	reg, _ := regexp.Compile(`^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	return reg.MatchString(receiver.OriVal())
}

// IsURL 判断是否是URL
func (receiver Str) IsURL() bool {
	reg, _ := regexp.Compile(`^http[s]?://[\w.]+[\w/]$`)
	return reg.MatchString(receiver.OriVal())
}

// IsNumber 判断是否是数字
func (receiver Str) IsNumber() bool {
	reg, _ := regexp.Compile(`^[0-9]+$`)
	return reg.MatchString(receiver.OriVal())
}

// IsEnglish 判断是否是英文
func (receiver Str) IsEnglish() bool {
	reg, _ := regexp.Compile(`^[a-zA-Z]+$`)
	return reg.MatchString(receiver.OriVal())
}

// IsChinese 判断是否是中文
func (receiver Str) IsChinese() bool {
	for _, r := range receiver.OriVal() {
		if !unicode.Is(unicode.Han, r) {
			return false
		}
	}
	return true
}

// IsLowerCase 判断是否是小写
func (receiver Str) IsLowerCase() bool {
	reg, _ := regexp.Compile(`^[a-z]+$`)
	return reg.MatchString(receiver.OriVal())
}

// IsUpperCase 判断是否是大写
func (receiver Str) IsUpperCase() bool {
	reg, _ := regexp.Compile(`^[A-Z]+$`)
	return reg.MatchString(receiver.OriVal())
}

/****************************/

/******* - Other - *******/

// CharSub 以char的方式进行字符串截取
// start:开始位置，strLength:截取长度
// 如果 start < 0, 则从字符串末尾开始计算
// 如果 strLength <= 0, 则截取到字符串末尾
func (receiver Str) CharSub(start int, length ...int) Str {
	charList := []rune(receiver.OriVal())
	l := len(charList)
	step := 0
	end := 0

	if len(length) == 0 {
		step = l
	} else {
		step = length[0]
	}

	if start < 0 {
		start = l + start
	}
	end = start + step

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}

	if start > l {
		start = l
	}

	if end < 0 {
		end = 0
	}

	if end > l {
		end = l
	}
	return Str(charList[start:end])
}

// CharLen 返回char字符串长度
func (receiver Str) CharLen() int {
	return len([]rune(receiver.OriVal()))
}

// Before 获取某个字符串第一个出现位置之前的字符串，如果不存在返回源字符串
func (receiver Str) Before(target string) Str {
	if target == "" {
		return receiver
	}
	i := strings.Index(receiver.OriVal(), target)
	if i != -1 {
		return Str(receiver.OriVal()[:i])
	}
	return receiver
}

// BeforeLast 获取某个字符串最后出现位置之前的字符串，如果不存在返回源字符串
func (receiver Str) BeforeLast(target string) Str {
	if target == "" {
		return receiver
	}
	i := strings.LastIndex(receiver.OriVal(), target)
	if i != -1 {
		return Str(receiver.OriVal()[:i])
	}
	return receiver
}

// After 获取某个字符串第一个出现位置之后的字符串，如果不存在返回源字符串
func (receiver Str) After(target string) Str {
	if target == "" {
		return receiver
	}
	i := strings.Index(receiver.OriVal(), target)
	if i != -1 {
		return Str(receiver.OriVal()[i+len(target):])
	}
	return receiver
}

// AfterLast 获取某个字符串最后出现位置之后的字符串，如果不存在返回源字符串
func (receiver Str) AfterLast(target string) Str {
	if target == "" {
		return receiver
	}
	i := strings.LastIndex(receiver.OriVal(), target)
	if i != -1 {
		return Str(receiver.OriVal()[i+len(target):])
	}
	return receiver
}
