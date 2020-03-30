/**
 * 面试题62. 圆圈中最后剩下的数字
 * https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
 **/

package problems

func LastRemaining(n int, m int) int {
	var ans int
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}

	return ans
}
