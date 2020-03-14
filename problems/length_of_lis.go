/**
 * 300. 最长上升子序列
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/
 **/

package problems

func LengthOfLIS(nums []int) int {
	var d []int

	for i := 0; i < len(nums); i++ {
		switch {
		case len(d) == 0:
			fallthrough
		case nums[i] > d[len(d)-1]:
			d = append(d, nums[i])
		default:
			for j := 0; j < len(d); j++ {
				if nums[i] <= d[j] {
					d[j] = nums[i]
					break
				}

			}
		}
	}

	return len(d)
}
