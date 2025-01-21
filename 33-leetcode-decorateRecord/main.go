package main

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

func (q *Queue) ToValSlice() []int {
	s := make([]int, 0, 10)
	i := q.head
	for i != q.tail {
		s = append(s, q.val[i].Val)
		i = (i + 1) % len(q.val)
	}
	return s
}

func decorateRecord(root *TreeNode) [][]int {
	res := make([][]int, 0, 10)
	parentQueue := NewQueue()
	childQueue := NewQueue()
	if root != nil {
		parentQueue.Push(root)
		res = append(res, parentQueue.ToValSlice())
	}

	for !parentQueue.Empty() || !childQueue.Empty() {
		node, b := parentQueue.Pop()
		if b {
			if node.Left != nil {
				childQueue.Push(node.Left)
			}
			if node.Right != nil {
				childQueue.Push(node.Right)
			}
		}
		if parentQueue.Empty() && !childQueue.Empty() {
			parentQueue, childQueue = childQueue, parentQueue
			res = append(res, parentQueue.ToValSlice())
		}
	}

	return res
}

func main() {

}
