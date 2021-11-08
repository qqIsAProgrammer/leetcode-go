package binary_tree

func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	iter := root
	for iter != nil || len(stack) > 0 {
		for iter != nil {
			ans = append(ans, iter.Val)
			stack = append(stack, iter)
			iter = iter.Left
		}
		iter = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return ans
}

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	iter := root
	for iter != nil || len(stack) > 0 {
		for iter != nil {
			stack = append(stack, iter)
			iter = iter.Left
		}
		ans = append(ans, stack[len(stack)-1].Val)
		iter = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return ans
}

// 对前序遍历顺序进行修改
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	iter := root
	for iter != nil || len(stack) > 0 {
		for iter != nil {
			ans = append([]int{iter.Val}, ans...)
			stack = append(stack, iter)
			iter = iter.Right
		}
		iter = stack[len(stack)-1].Left
		stack = stack[:len(stack)-1]
	}
	return ans
}
