package main

import (
	"fmt"
	"math"
)

/**
Definition for a binary tree node.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, b := isBalancedWithHeight(root)
	return b
}

func isBalancedWithHeight(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftHeight, leftIsBalance := isBalancedWithHeight(root.Left)
	rightHeight, rightIsBalance := isBalancedWithHeight(root.Right)
	float64LeftHeight, float64RightHeight := float64(leftHeight), float64(rightHeight)
	height := 1 + int(math.Max(float64LeftHeight, float64RightHeight))
	if !leftIsBalance || !rightIsBalance {
		return height, false
	}
	return height, math.Abs(float64LeftHeight-float64RightHeight) <= 1.0
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(isBalanced(root))

	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isBalanced(root))
}
