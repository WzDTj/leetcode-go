/**
 * 876. 链表的中间结点
 * https://leetcode-cn.com/problems/middle-of-the-linked-list/
 **/

package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MiddleNode(head *ListNode) *ListNode {
	current, size := head, 0
	for current != nil {
		size++
		current = current.Next
	}

	target := size / 2

	current = head
	for i := 0; current != nil; i++ {
		if i == target {
			return current
		}

		current = current.Next
	}

	return nil
}
