/**
 * 1071. 字符串的最大公因子
 * https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/
 **/

package problems

import (
	"math"
	"sort"
	"strings"
)

func GcdOfStrings(str1 string, str2 string) string {
	shorter := getShorterString(str1, str2)
	commonDivisors := getCommonDivisors(len(str1), len(str2))

	for _, v := range commonDivisors {
		s := shorter[:v]

		if isGcdOfStrings([]string{str1, str2}, s) {
			return s
		}
	}

	return ""
}

func isGcdOfStrings(strs []string, divider string) bool {
	for _, v := range strs {
		remainder := strings.ReplaceAll(v, divider, "")
		if remainder != "" {
			return false
		}
	}

	return true
}

func getShorterString(str1, str2 string) string {
	if len(str1) < len(str2) {
		return str1
	}

	return str2
}

func getCommonDivisors(num1, num2 int) []int {
	min := int(math.Min(float64(num1), float64(num2)))

	var commonDivisors []int
	for i := 1; i <= min/2; i++ {
		if min%i != 0 {
			continue
		}
		if i*i == min {
			commonDivisors = append(commonDivisors, i)
		} else {
			commonDivisors = append(commonDivisors, i, min/i)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(commonDivisors)))

	return commonDivisors
}
