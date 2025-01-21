package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSymmetricTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	a := root.Left
	b := mirrorTree(root.Right)
	return isSame(a, b) && isSame(b, a)
}

func isSame(major *TreeNode, minor *TreeNode) bool {
	if major == nil {
		return true
	}
	if minor == nil {
		return false
	}
	return major.Val == minor.Val && isSame(major.Left, minor.Left) && isSame(major.Right, minor.Right)
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}

func main() {
}
