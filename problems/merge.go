/**
 * 面试题 10.01. 合并排序的数组
 * https://leetcode-cn.com/problems/sorted-merge-lcci/
 **/

package problems

func Merge(A []int, m int, B []int, n int) {
	backup := make([]int, len(A))
	copy(backup, A)

	for i, j := 0, 0; i < m || j < n; {
		switch {
		case i == m:
			A[i+j] = B[j]
			j++
		case j == n:
			fallthrough
		case backup[i] <= B[j]:
			A[i+j] = backup[i]
			i++
		default:
			A[i+j] = B[j]
			j++
		}
	}
}
