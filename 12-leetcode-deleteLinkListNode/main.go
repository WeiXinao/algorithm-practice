package main

/*
	删除链表节点

	给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
	返回删除后的链表的头节点。

	示例 1:
	输入: head = [4,5,1,9], val = 5
	输出: [4,1,9]
	解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
	
	示例 2:
	输入: head = [4,5,1,9], val = 1
	输出: [4,5,9]
	解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
*/

/**
 * Definition for singly-linked list.
 */

import (
	"fmt"
)
type ListNode struct {
		Val int
		Next *ListNode
}
 
func deleteNode(head *ListNode, val int) *ListNode {
	var (
		p *ListNode
		q *ListNode
	)
	if head.Val == val {
		return head.Next
	}
	q = head
	p = head.Next
	for p != nil {
		if p.Val == val {
			q.Next = p.Next
		}	else {
			q = q.Next
		}
		p = p.Next
	}
	return head
}

func outputLinkList(head *ListNode) {
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
		_, err := fmt.Scanln(&input);
		if err != nil {
			break
		} 
		newNode := &ListNode{}
		newNode.Val = input
		p.Next = newNode
		p = p.Next
	}

	outputLinkList(head.Next)

	delVal := 0
	fmt.Scanln()
	fmt.Scanln(&delVal)
	fmt.Printf("delVal: %v\n", delVal)
	head = deleteNode(head.Next, delVal)
	
	outputLinkList(head)
}