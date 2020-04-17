package utils

import (
	"math"
	"unicode/utf8"
)

func SmsCounter(content string) int {
	count := utf8.RuneCountInString(content)
	if count <= 70 {
		return 1
	}

	return int(math.Ceil(float64(count) / 67))
}
