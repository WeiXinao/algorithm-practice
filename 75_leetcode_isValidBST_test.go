package test

func isValidBST(root *TreeNode) bool {
  if root == nil {
		return true
	}
	return isValidBST(root.Left) && isValidBST(root.Right) &&
		isGreaterThan(root.Right, root.Val) && isLessThan(root.Left, root.Val)
}

func isGreaterThan(root *TreeNode, target int) bool {
	if root == nil {
		return true
	}
	return root.Val > target && isGreaterThan(root.Left, target) && isGreaterThan(root.Right, target)
}

func isLessThan(root *TreeNode, target int) bool {
	if root == nil {
		return true
	}
	return root.Val < target && isLessThan(root.Left, target) && isLessThan(root.Right, target)
}