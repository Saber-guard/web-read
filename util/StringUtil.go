package util

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
