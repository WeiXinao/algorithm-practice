package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	lenA, lenB, cnt := 0, 0, 0
	for p != nil {
		lenA++
		p = p.Next
	}
	for q != nil {
		lenB++
		q = q.Next
	}
	p, q = headA, headB
	if lenA > lenB {
		cnt = lenA - lenB
		for cnt > 0 {
			p = p.Next
			cnt--
		}
	} else if lenA < lenB {
		cnt = lenB - lenA
		for cnt > 0 {
			q = q.Next
			cnt--
		}
	}
	var r *ListNode
	for p != nil && q != nil {
		if p == q {
			r = p
			break
		}
		p = p.Next
		q = q.Next
	}
	return r
}

func outputLinkList(head *ListNode) {
	if head == nil {
		fmt.Println("[]")
		return
	}
	p := head
	fmt.Print("[")
	for p != nil {
		fmt.Print(p.Val)
		if p.Next != nil {
			fmt.Print(", ")
		} else {
			fmt.Print("]")
		}
		p = p.Next
	}
	fmt.Println()
}

func inputLinkList() *ListNode {
	head := &ListNode{}
	input := 0

	p := head
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			break
		}
		newNode := &ListNode{}
		newNode.Val = input
		p.Next = newNode
		p = p.Next
	}
	return head.Next
}

func main() {
	headA := inputLinkList()
	var in int
	fmt.Scanln(&in)
	outputLinkList(headA)
	headB := inputLinkList()
	outputLinkList(headB)
}
