package binary_tree

// AC
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	var idx int
	for i := 0; i < len(preorder); i++ {
		if inorder[i] == preorder[0] {
			idx = i
			break
		}
	}

	root.Left = buildTree(preorder[1:idx+1], inorder[0:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
