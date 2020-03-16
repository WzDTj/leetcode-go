/**
 * 面试题 01.06. 字符串压缩
 * https://leetcode-cn.com/problems/compress-string-lcci/
 **/

package problems

import (
	"strconv"
)

func CompressString(S string) string {
	var ans, char string
	count, length := 0, len(S)

	for i := 0; i < length; i++ {
		currentChar := string(S[i])
		switch {
		case count == 0:
			char = currentChar
			count = 1
		case currentChar != char:
			ans += (char + strconv.Itoa(count))
			char = currentChar
			count = 1
		default:
			count++
		}

	}
	ans += (char + strconv.Itoa(count))

	if len(ans) >= length {
		return S
	}
	return ans
}
