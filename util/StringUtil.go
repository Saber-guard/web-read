package util

import (
	"strconv"
)

type StringUtil struct {
}

// 按指定长度将字符串截取为数组
func (s StringUtil) SplitByLen(str string, unitLen int) []string {
	var (
		arr        []string
		strRune    = []rune(str)
		strruneLen = len(strRune)
	)

	// 循环截取
	offset := 0
	for offset < strruneLen {
		//  避免超出
		end := offset + unitLen
		if end > strruneLen {
			end = strruneLen
		}

		// 截取
		arr = append(arr, string(strRune[offset:end]))
		offset += unitLen
	}
	return arr
}

// 把输入的类型转换为字符串
func (s StringUtil) AllToStr(input interface{}) (str string, err error) {
	if conv, ok := input.(string); ok {
		str = conv
		return
	}
	if conv, ok := input.(int); ok {
		str = strconv.Itoa(conv)
		return
	}
	if conv, ok := input.(float32); ok {
		str = strconv.FormatFloat(float64(conv), 'f', 2, 32)
		return
	}
	if conv, ok := input.(float64); ok {
		str = strconv.FormatFloat(conv, 'f', 2, 64)
		return
	}

	return
}
