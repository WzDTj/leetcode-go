/**
 * 912. 排序数组
 * https://leetcode-cn.com/problems/sort-an-array/
 **/

package problems

import "sort"

func SortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}
