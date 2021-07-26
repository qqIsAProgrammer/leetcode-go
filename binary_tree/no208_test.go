package binary_tree

import "testing"

func TestInsert(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
}

func TestSearch(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	res := trie.StartsWith("app")
	t.Log(res)
}
