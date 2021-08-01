package binary_tree

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	l := lowestCommonAncestor2(root.Left, p, q)
	r := lowestCommonAncestor2(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return l
	} else {
		return r
	}
}
