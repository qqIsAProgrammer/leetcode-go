package binary_tree

import (
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
}

func TestSearch(t *testing.T) {
	var tn []*TreeNode
	tn2 := make([]*TreeNode, 0)
	log.Println(tn)
	log.Println(tn2)
}
