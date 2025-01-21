package main

/**
 * TreeNode Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	if A == nil {
		return false
	}
	if A.Val == B.Val && isSame(A, B) {
		return true
	} else if isSubStructure(A.Left, B) {
		return true
	} else {
		return isSubStructure(A.Right, B)
	}
}

func isSame(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return A.Val == B.Val && isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
}
