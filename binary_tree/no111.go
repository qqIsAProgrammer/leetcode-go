package binary_tree

func minInt(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// AC
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	return minInt(minDepth(root.Left), minDepth(root.Right)) + 1
}
