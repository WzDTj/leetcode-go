/**
 * 836. 矩形重叠
 * https://leetcode-cn.com/problems/rectangle-overlap/
 **/

package problems

func IsRectangleOverlap(rec1 []int, rec2 []int) bool {
	switch {
	case rec2[0] >= rec1[2]:
		fallthrough
	case rec2[1] >= rec1[3]:
		fallthrough
	case rec2[2] <= rec1[0]:
		fallthrough
	case rec2[3] <= rec1[1]:
		return false
	default:
		return true
	}
}
