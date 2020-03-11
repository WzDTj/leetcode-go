/**
 * 1013. 将数组分成和相等的三个部分
 * https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/
 **/

package problems

func CanThreePartsEqualSum(A []int) bool {
	var sum int

	for _, v := range A {
		sum += v
	}

	if sum%3 != 0 {
		return false
	}

	target := sum / 3

	sums, part := [2]int{0, 0}, 0

	for i := 0; i < len(A)-1; i++ {
		sums[part] += A[i]

		if sums[part] == target {
			part++
		}
		if part == 2 {
			return true
		}
	}

	return false
}
