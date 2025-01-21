package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	val  []*TreeNode
	head int
	tail int
}

func NewQueue() *Queue {
	return &Queue{
		val:  make([]*TreeNode, 600),
		head: 0,
		tail: 0,
	}
}

func (q *Queue) Empty() bool {
	return q.head == q.tail
}

func (q *Queue) Push(node *TreeNode) bool {
	if (q.tail+1)%len(q.val) == q.head {
		return false
	}
	q.val[q.tail] = node
	q.tail = (q.tail + 1) % len(q.val)
	return true
}

func (q *Queue) Pop() (*TreeNode, bool) {
	if q.tail == q.head {
		return nil, false
	}
	popNode := q.val[q.head]
	q.head = (q.head + 1) % len(q.val)
	return popNode, true
}

func decorateRecord(root *TreeNode) []int {
	res := make([]int, 0, 100)
	q := NewQueue()
	if root != nil {
		q.Push(root)
	}

	for !q.Empty() {
		node, b := q.Pop()
		if b {
			res = append(res, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}
	return res
}

func main() {
}
