package binary_tree

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			pred := curr.Left
			for pred.Right != nil {
				pred = pred.Right
			}
			pred.Right = curr.Right
			curr.Right = next
			curr.Left = nil
		}
		curr = curr.Right
	}
}
