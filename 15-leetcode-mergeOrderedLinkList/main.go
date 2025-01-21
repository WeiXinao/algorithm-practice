package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainningPlan(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Next: nil}
	r := head
	p := l1
	q := l2
	var newNode *ListNode
	for p != nil && q != nil {
		if p.Val > q.Val {
			newNode = &ListNode{q.Val, nil}
			q = q.Next
		} else {
			newNode = &ListNode{p.Val, nil}
			p = p.Next
		}
		r.Next = newNode
		r = r.Next
	}
	if p != nil {
		r.Next = p
	}
	if q != nil {
		r.Next = q
	}
	return head.Next
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
	l1 := inputLinkList()
	var in int
	fmt.Scanln(&in)
	outputLinkList(l1)
	l2 := inputLinkList()
	outputLinkList(l2)

	res := trainningPlan(l1, l2)
	outputLinkList(res)
}
