// 211. Design Add and Search Words Data Structure
//
// https://leetcode.com/problems/design-add-and-search-words-data-structure/
//	Dictionary can be easily solved with data structure map.
//	However, '.' replace is not that easy to be solved with just map.
//	'.' is just a character, trie data structure chains a character to character to make the specific word.
//	Such feature can ignore the '.' character, and make the comparison easier than just solving by map.

// WordDictionary is the node of trie tree.
// This node has several children nodes and each child node is also a node of trie tree.
type WordDictionary struct {
	m map[byte]*WordDictionary
}

// Constructor function returns the root node of trie tree which is not the pointer variable.
func Constructor() WordDictionary {
	return WordDictionary{m: make(map[byte]*WordDictionary)}
}

// AddWord function adds the word into the dictionary by spliting each character.
// The each splitted character will be a node of trie tree.
func (this *WordDictionary) AddWord(word string) {
	current := this
	for i := 0; i < len(word); i++ {
		w := word[i]
		if v, ok := current.m[w]; ok {
			current = v
		} else {
			current.m[w] = &WordDictionary{m: make(map[byte]*WordDictionary)}
			current = current.m[w]
		}
	}
	current.m['0'] = &WordDictionary{m: make(map[byte]*WordDictionary)}
}

// Search function checks whether word is exist in the dictionary or not.
// This is done by recursive way.
func (this *WordDictionary) Search(word string) bool {
	var recursion func(word string, current *WordDictionary) bool
	recursion = func(word string, current *WordDictionary) bool {
		if len(word) == 0 {
			return false
		}
		if word[0] == '.' {
			for _, v := range current.m {
				if v.m['0'] != nil && len(word) == 1 {
					return true
				}
				if recursion(word[1:], v) {
					return true
				}
			}
		} else {
			if v, ok := current.m[word[0]]; ok {
				if v.m['0'] != nil && len(word) == 1 {
					return true
				} else {
					return recursion(word[1:], v)
				}
			}
		}
		return false
	}
	return recursion(word, this)
}

// Your WordDictionary object will be instantiated and called as such:
// obj := Constructor();
// obj.AddWord(word);
// param_2 := obj.Search(word);

