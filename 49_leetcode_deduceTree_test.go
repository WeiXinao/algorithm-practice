package test

import (
	"slices"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deduceTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	inIndexBoundary := slices.Index(inorder, preorder[0])
	preIndexBoundary := 1 + inIndexBoundary
	node := &TreeNode{
		Val:   preorder[0],
		Left:  deduceTree(preorder[1:preIndexBoundary], inorder[:inIndexBoundary]),
		Right: deduceTree(preorder[preIndexBoundary:], inorder[inIndexBoundary+1:]),
	}
	return node
}
