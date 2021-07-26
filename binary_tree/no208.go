package binary_tree

type Trie struct {
	children [26]*Trie
	flag     bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	iter := t
	for _, ch := range word {
		if iter.children[ch-'a'] == nil {
			iter.children[ch-'a'] = &Trie{}
		}
		iter = iter.children[ch-'a']
	}
	iter.flag = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	child := t.searchPrefix(word)
	return child != nil && child.flag
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

func (t *Trie) searchPrefix(prefix string) *Trie {
	iter := t
	for _, ch := range prefix {
		if iter.children[ch-'a'] == nil {
			return nil
		}
		iter = iter.children[ch-'a']
	}
	return iter
}
