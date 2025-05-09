package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isSymmetric(root *TreeNode) bool {
	queue := 	make([]*TreeNode, 1500)
	head, tail := 0, 0
	queue[tail] = root
	tail++
	for head != tail {
		for i := head; i < (head + tail)/2; i++ {
			if queue[i] == nil && queue[tail - 1 - i + head] != nil {
				return false
			}
			if queue[i] != nil && queue[tail - 1 - i + head] == nil {
				return false
			}
			if queue[i] != nil && queue[i].Val != queue[tail - 1 - i + head].Val {
				return false
			}
		}	
		end := tail 
		for head != end {
			if queue[head] == nil {
				head++
				continue
			}
			queue[tail] = queue[head].Left
			tail++
			queue[tail] = queue[head].Right
			tail++
			head++
		}
	}
	return true
}

func TestIsSymmetric(t *testing.T) {
	testCases := []struct{
		name string
		root *TreeNode
		expect bool
	}{
		{
			name: "对称",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expect: true,
		},
		{
			name: "不对称",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			expect: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			res := isSymmetric(tc.root)
			assert.Equal(t, tc.expect, res)
		})
	}
}


