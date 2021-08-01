package binary_tree

import (
	"fmt"
	"testing"
)

func TestTraversal(t *testing.T) {
	// [1,0,48,null,null,12,49]
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 0},
		Right: &TreeNode{
			Val:   48,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 49},
		},
	}
	fmt.Println(minDiffInBST(root))
}
