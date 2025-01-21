package main

import "fmt"

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

func (q *Queue) Length() int {
	i := 0
	p := q.head
	for {
		if p == q.tail {
			break
		}
		p = (p + 1) % len(q.val)
		i++
	}
	return i
}

func ReverseSlice(s []int) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func decorateRecord(root *TreeNode) [][]int {
	cnt := 0
	res := make([][]int, 0, 10)
	q := NewQueue()
	if root != nil {
		q.Push(root)
	}

	for !q.Empty() {
		length := q.Length()
		tmp := []int{}
		for i := 0; i < length; i++ {
			node, b := q.Pop()
			if b {
				tmp = append(tmp, node.Val)
				if node.Left != nil {
					q.Push(node.Left)
				}
				if node.Right != nil {
					q.Push(node.Right)
				}
			}
		}
		if cnt%2 == 1 {
			ReverseSlice(tmp)
		}
		res = append(res, tmp)
		cnt++
	}
	return res
}

func main() {
	//s := []int{1, 2, 3, 4, 5, 6}
	//ReverseSlice(s)
	//fmt.Println(s)

	q := NewQueue()
	q.Push(&TreeNode{
		Val: 1,
	})
	q.Push(&TreeNode{
		Val: 2,
	})
	fmt.Println(q.Length())

}
