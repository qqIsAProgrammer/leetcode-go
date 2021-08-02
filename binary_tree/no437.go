package binary_tree

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	count := 0
	countPath(root, targetSum, &count)
	return count + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func countPath(root *TreeNode, targetSum int, count *int) {
	if root == nil {
		return
	}
	targetSum = targetSum - root.Val
	if targetSum == 0 {
		*count++
	}
	countPath(root.Left, targetSum, count)
	countPath(root.Right, targetSum, count)
}
