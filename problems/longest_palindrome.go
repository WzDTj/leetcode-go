/**
 * 409. 最长回文串
 * https://leetcode-cn.com/problems/longest-palindrome/
 **/

package problems

func LongestPalindrome(s string) int {
	ans, countMap := 0, map[string]int{}
	for _, v := range s {
		countMap[string(v)]++
	}

	for _, v := range countMap {
		ans += v / 2 * 2
	}

	if ans < len(s) {
		ans++
	}

	return ans
}
