package eztools

import (
	"math/rand"
	"time"
)

// RandomStr 是一个随机字符串类型
/*
	使用说明:
		Get 随机获取指定长度字符串
*/
type RandomStr string

const (
	RandomStrHex    RandomStr = "0123456789abcdef"
	RandomStrNum    RandomStr = "0123456789"
	RandomStrEn     RandomStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	RandomStrEnU    RandomStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	RandomStrEnL    RandomStr = "abcdefghijklmnopqrstuvwxyz"
	RandomStrCase   RandomStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&_="
	RandomStrNumEnL RandomStr = "0123456789abcdefghijklmnopqrstuvwxyz"
	RandomStrNumEnU RandomStr = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Get 随机获取字符串, length: 长度
func (receiver RandomStr) Get(length int) string {
	bytes := []byte(receiver)
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}
