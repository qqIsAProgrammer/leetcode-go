package binary_tree

import "math"

var res = math.MaxInt32
var pre *TreeNode

func minDiffInBST(root *TreeNode) int {
	inorder(root)
	return res
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	if pre != nil && root.Val-pre.Val < res {
		res = root.Val - pre.Val
	}
	pre = root
	inorder(root.Right)
}
