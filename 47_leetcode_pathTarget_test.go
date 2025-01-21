package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * Definition for a binary tree node.
 */
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func TestPathTarget(t *testing.T) {
	testCases := []struct {
		name      string
		root      *TreeNode
		targetSum int
		expect    [][]int
	}{
		{
			name: "路径存在",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val:   13,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			targetSum: 22,
			expect: [][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
		{
			name: "路径不存在，目标和不为0",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			targetSum: 5,
			expect:    [][]int{},
		},
		{
			name: "路径不存在，目标和为0",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			targetSum: 0,
			expect:    [][]int{},
		},
		{
			name: "存在节点为负数",
			root: &TreeNode{
				Val:  -2,
				Left: nil,
				Right: &TreeNode{
					Val:   -3,
					Left:  nil,
					Right: nil,
				},
			},
			targetSum: -5,
			expect: [][]int{
				{-2, -3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := pathTarget(tc.root, tc.targetSum)
			assert.Equal(t, tc.expect, res)
		})
	}
}

func pathTarget(root *TreeNode, target int) [][]int {
	paths := make([][]int, 0, 5)
	path := make([]int, 0, 5)
	pathTargetSum(root, target, 0, &path, &paths)
	return paths
}

func pathTargetSum(root *TreeNode, target int, sum int, path *[]int, paths *[][]int) {
	if root == nil {
		return
	}
	newPath := make([]int, len(*path))
	copy(newPath, *path)
	sum += root.Val
	newPath = append(newPath, root.Val)
	if sum == target && root.Left == nil && root.Right == nil {
		*paths = append(*paths, newPath)
	}
	pathTargetSum(root.Left, target, sum, &newPath, paths)
	pathTargetSum(root.Right, target, sum, &newPath, paths)
}
