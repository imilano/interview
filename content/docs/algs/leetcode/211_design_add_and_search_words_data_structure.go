package leetcode

/*
Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:

WordDictionary() Initializes the object.
void addWord(word) Adds word to the data structure, it can be matched later.
bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.
*/

// 其实这里也是前缀树的实现。不过简单实现的话，我们可以使用 map，不过由于此处有通配符 . ，所以使用 map 来实现前缀树可能有点复杂。
// 使用一个长度为 26 的数组来实现前缀树的节点
type WordDictionary struct {
	isWord   bool                // 判断一个节点是不是走到单词尾
	children [26]*WordDictionary // 子节点，长度为 26 的数组
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	size := len(word)
	for i, r := range word {
		// 如果当前节点的子节点不存在，则创建一个
		loc := r - 'a'
		if cur.children[loc] == nil {
			cur.children[loc] = &WordDictionary{}
		}

		// 移动到下一个子节点
		cur = cur.children[loc]

		// 如果已经遍历到单词末尾，则将该节点置为单词结尾
		if i == size-1 {
			cur.isWord = true
		}
	}
}

// 由于增加通配符，所以当遇到 . 时，就需要对每一个节点进行深度遍历，查找所有子树，只要一个节点返回 true，则整个 search 返回 true
func (this *WordDictionary) Search(word string) bool {
	return searchDFS(this, word)
}

func searchDFS(node *WordDictionary, word string) bool {
	size := len(word)
	if size == 0 {
		return node.isWord
	}
	// TODO recheck
	// 注意这里如何应对通配符
	if word[0] == '.' {
		for _, child := range node.children {
			if child != nil && searchDFS(child, word[1:]) {
				return true
			}
		}

		return false
	} else {
		loc := word[0] - 'a'
		if node.children[loc] == nil {
			return false
		}

		return searchDFS(node.children[loc], word[1:])
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
