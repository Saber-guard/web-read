package util

type StringUtil struct {
}

// 按指定长度将字符串截取为数组
func (s StringUtil) SplitByLen(str string, unitLen int) (arr []string) {
	arr = make([]string, 1)
	strRune := []rune(str)
	strruneLen := len(strRune)
	// 循环截取
	offset := 0
	for offset < strruneLen {
		if offset == 0 {
			arr[0] = string(strRune[offset : offset+unitLen])
		} else {
			arr = append(arr, string(strRune[offset:offset+unitLen]))
		}
		offset += unitLen
	}
	return
}
