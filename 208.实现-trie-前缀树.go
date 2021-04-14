package leetcode_go

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	count  int
	childs []*TrieNode
}

/** Initialize your data structure here. */
// func Constructor() Trie {
// 	return Trie{
// 		root: &TrieNode{
// 			count:  0,
// 			childs: make([]*TrieNode, 26),
// 		},
// 	}
// }

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	chs := []byte(word)
	ptr := this.root
	for _, v := range chs {
		if ptr.childs[v-'a'] == nil {
			ptr.childs[v-'a'] = &TrieNode{
				count:  0,
				childs: make([]*TrieNode, 26),
			}
		}
		ptr = ptr.childs[v-'a']
	}
	ptr.count++
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	ptr := this.visit(word)
	if ptr == nil {
		return false
	}
	if ptr.count > 0 {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	ptr := this.visit(prefix)
	if ptr == nil {
		return false
	}
	return true
}

func (this *Trie) visit(word string) *TrieNode {
	chs := []byte(word)
	ptr := this.root
	for _, v := range chs {
		if ptr.childs[v-'a'] == nil {
			return nil
		}
		ptr = ptr.childs[v-'a']
	}
	return ptr
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
