package binary_tree

// AC
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		var tn *TreeNode
		for size != 0 {
			tn = queue[0]
			queue = queue[1:]
			if tn.Left != nil {
				queue = append(queue, tn.Left)
			}
			if tn.Right != nil {
				queue = append(queue, tn.Right)
			}
			size--
		}
		res = append(res, tn.Val)
	}

	return res
}
