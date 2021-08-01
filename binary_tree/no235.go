package binary_tree

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor1(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor1(root.Right, p, q)
	}
	return root
}
