package eztools

type threeInt int
type threeInt8 int8
type threeInt16 int16
type threeInt32 int32
type threeInt64 int64

type threeUint uint
type threeUint8 uint8
type threeUint16 uint16
type threeUint32 uint32
type threeUint64 uint64

type threeFloat32 float32
type threeFloat64 float64

type threeComplex64 complex64
type threeComplex128 complex128

type threeString string

const (
	TInt   threeInt   = 0 // "三木运算整数"
	TInt8  threeInt8  = 1 // "三木运算整数8位"
	TInt16 threeInt16 = 2 // "三木运算整数16位"
	TInt32 threeInt32 = 3 // "三木运算整数32位"
	TInt64 threeInt64 = 4 // "三木运算整数64位"

	TUint   threeUint   = 5 // "三木运算无符号整数"
	TUint8  threeUint8  = 6 // "三木运算无符号整数8位"
	TUint16 threeUint16 = 7 // "三木运算无符号整数16位"
	TUint32 threeUint32 = 8 // "三木运算无符号整数32位"
	TUint64 threeUint64 = 9 // "三木运算无符号整数64位"

	TFloat32 threeFloat32 = 10 // "三木运算浮点数32位"
	TFloat64 threeFloat64 = 11 // "三木运算浮点数64位"

	TComplex64  threeComplex64  = 12 // "三木运算复数32位"
	TComplex128 threeComplex128 = 13 // "三木运算复数64位"

	TString threeString = "~" // "三木运算字符串"
)

func (receiver threeInt) If(condition bool, trueVal int, falseVal int) int {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeInt8) If(condition bool, trueVal int8, falseVal int8) int8 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeInt16) If(condition bool, trueVal int16, falseVal int16) int16 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeInt32) If(condition bool, trueVal int32, falseVal int32) int32 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeInt64) If(condition bool, trueVal int64, falseVal int64) int64 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeUint) If(condition bool, trueVal uint, falseVal uint) uint {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeUint8) If(condition bool, trueVal uint8, falseVal uint8) uint8 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeUint16) If(condition bool, trueVal uint16, falseVal uint16) uint16 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeUint32) If(condition bool, trueVal uint32, falseVal uint32) uint32 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeUint64) If(condition bool, trueVal uint64, falseVal uint64) uint64 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeFloat32) If(condition bool, trueVal float32, falseVal float32) float32 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeFloat64) If(condition bool, trueVal float64, falseVal float64) float64 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeComplex64) If(condition bool, trueVal complex64, falseVal complex64) complex64 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeComplex128) If(condition bool, trueVal complex128, falseVal complex128) complex128 {
	if condition {
		return trueVal
	}
	return falseVal
}

func (receiver threeString) If(condition bool, trueVal string, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}
