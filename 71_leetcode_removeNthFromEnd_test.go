package test

// https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn2925/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	var (
		p = head
		q = head
	)
	for ;n > 0; n-- {
		q = q.Next		
	}
	if q == nil && n > 0 {
		return nil
	}
	if q == nil && n == 0 {
		return head.Next
	}
	for q.Next != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return head
}