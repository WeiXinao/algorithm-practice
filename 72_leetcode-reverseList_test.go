package test

// https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnnhm6/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	first := new(ListNode)
	first.Next = head

	pre := first.Next
	cur := pre.Next
	for pre.Next != nil {
		pre.Next = cur.Next
		cur.Next = first.Next
		first.Next = cur
		cur = pre.Next
	}
	return first.Next
}