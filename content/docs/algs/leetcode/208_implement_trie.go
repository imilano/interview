package leetcode

// import "strings"

/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
*/

/*
// 方法 1， 直接使用一个 map 来记录一个字符是否出现过。时间复杂度比较大
type Trie struct {
    m map[string]bool
}


func Constructor() Trie {
    m := make(map[string]bool)
	return Trie{m}
}


func (this *Trie) Insert(word string)  {
    this.m[word] = true
}


func (this *Trie) Search(word string) bool {
    if _, ok := this.m[word]; ok {
		return true
	}

	return false
}


func (this *Trie) StartsWith(prefix string) bool {
    for str, _ := range this.m {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}

	return false
}
*/

// 使用数组来表示节点
type Trie struct {
	isWord   bool
	children [26]*Trie
}

// comment this to avoid conflict
// func Constructor() Trie {
// 	return Trie{}
// }

func (this *Trie) Insert(word string) {
	cur := this
	size := len(word)
	for i, r := range word {
		loc := r - 'a'
		if cur.children[loc] == nil {
			cur.children[loc] = &Trie{}
		}

		cur = cur.children[loc]
		if i == size-1 {
			cur.isWord = true
		}
	}
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, r := range word {
		loc := r - 'a'
		if cur.children[loc] == nil {
			return false
		}

		cur = cur.children[loc]
	}

	return cur.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, r := range prefix {
		loc := r - 'a'
		if cur.children[loc] == nil {
			return false
		}

		cur = cur.children[loc]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
