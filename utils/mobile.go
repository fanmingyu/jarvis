package utils

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strings"
)

// MobileString将手机号自动脱敏的类型
type MobileString string

// String 自动脱敏的方法
func (m MobileString) String() string {
	mobiles := strings.Split(string(m), ",")
	result := []string{}
	for _, mobile := range mobiles {
		mobileByte := []byte(mobile)
		sum := md5.Sum(mobileByte)
		for i := 1; i < 7 && i < len(mobileByte); i++ {
			mobileByte[i] = '*'
		}
		result = append(result, fmt.Sprintf("%s(%x)", mobileByte, sum))
	}

	return strings.Join(result, ",")
}

func MobileFormat(mobile string) string {
	r, _ := regexp.Compile("(\\d{3})\\d{4}(\\d{4})")
	result := r.ReplaceAllString(mobile, "$1****$2")

	return result
}
