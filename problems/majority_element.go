/**
 * 169. 多数元素
 * https://leetcode-cn.com/problems/majority-element/
 **/

package problems

func MajorityElement(nums []int) int {
	maps, target := map[int]int{}, len(nums)>>1

	for _, v := range nums {
		maps[v]++
		if maps[v] > target {
			return v
		}
	}

	return 0
}
