package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func trainningPlan(head *ListNode) *ListNode {
	headOfhead := &ListNode{}
	headOfhead.Next = head
	if headOfhead.Next == nil {
		return head
	}
	p := headOfhead.Next
	q := p.Next
	for p.Next != nil {
		p.Next = q.Next
		q.Next = nil

		q.Next = headOfhead.Next
		headOfhead.Next = q

		q = p.Next
	}
	return headOfhead.Next
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

func main() {
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

	outputLinkList(head.Next)

	head = trainningPlan(head.Next)

	outputLinkList(head)
}
