package main

import "fmt"

//Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainingPlan(head *ListNode, cnt int) *ListNode {
	p := head
	q := head
	for cnt > 0 {
		q = q.Next
		cnt--
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}
	return p
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
	var (
		input = 0
		cnt   = 0
		p     = head
	)

	for {
		_, err := fmt.Scanf("%d\n", &input)
		if err != nil {
			break
		}
		newNode := &ListNode{}
		newNode.Val = input
		p.Next = newNode
		p = p.Next
	}

	fmt.Print("cnt:")
	fmt.Scanf("%d\n", &input)

	fmt.Scanf("%d\n", &cnt)

	outputLinkList(head.Next)

	head = trainingPlan(head.Next, cnt)

	fmt.Println(head.Val)
}
