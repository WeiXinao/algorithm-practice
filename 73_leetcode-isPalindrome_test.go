package test

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	new := &ListNode{
		Val: head.Val,
		Next: nil,
	}
	p := head.Next
	q := new
	for p != nil {
		node := &ListNode{
			Val: p.Val,
			Next: q,
		}
		q = node
		p = p.Next
	}
	p = head
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	if p != nil || q != nil {
		return false
	}
	return true
}

