package test

// https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnarn7/

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteNode1(node *ListNode) {
	p := node
	q := (*ListNode)(nil)
	for p.Next != nil {
		q = p
		p = p.Next
		q.Val = p.Val
	}
	q.Next = nil
	p = nil
}

func deleteNode(node *ListNode) {
	node.Val, node.Next.Val = node.Next.Val, node.Val
	node.Next = node.Next.Next
}

