package test

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := make([]*TreeNode, 2000)
	head, tail := 0, 0
	queue[tail] = root
	tail++
	res := make([][]int, 0)
	for head != tail {
		valsPerlevel := make([]int, 0)
		for i := head; i < tail; i++ {
			valsPerlevel = append(valsPerlevel, queue[i].Val)
		}
		res = append(res, valsPerlevel)
		end := tail
		for head != end {
			if queue[head].Left != nil {
				queue[tail] = queue[head].Left
				tail++
			}
			if queue[head].Right != nil {
				queue[tail] = queue[head].Right
				tail++
			}
			head++
		}
	}
	return res
}