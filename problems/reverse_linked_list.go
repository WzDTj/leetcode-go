/**
 * 206. 反转链表
 * https://leetcode-cn.com/problems/reverse-linked-list/
 **/

package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	current := head
	var previous *ListNode

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}
