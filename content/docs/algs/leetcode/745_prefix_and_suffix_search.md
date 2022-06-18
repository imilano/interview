---
title: 0745. Prefix and Suffix Search
weight: 10
tags: [
	"Hash Table",
	"String"
]
---
## Description
> Design a special dictionary with some words that searchs the words in it by a prefix and a suffix.
> 
> Implement the WordFilter class:
> 
> - `WordFilter(string[] words)` Initializes the object with the words in the dictionary.
> - `f(string prefix, string suffix)` Returns the index of the word in the dictionary, which has the prefix prefix and the suffix suffix. If there is more than one valid index, return the largest of them. If there is no such word in the dictionary, return -1.

## Solutions
### Hash Table
一个长度为 n 的单词会有 n 个前缀和后缀，那么也就有 {{< katext  >}} \Omicron(n^2) {{< /katext >}} 种组合。这里可以使用一个哈希表，来记录每个单词的前缀和后缀的组合对应的下标，key 为前缀和后缀的组合，为了应对相同字符出现的情况，我们可以在前后缀之间加上一个特殊字符，value 为 该前后缀组合对应的字符的下标。因为题目要求当字符相同时，返回最大的下标，那么这里我们只需要按照数组顺序从左向右进行遍历，然后遇到重复的前后缀组合，直接对下标进行替换即可。
```go
type WordFilter struct {
    dict map[string]int
}


func Constructor(words []string) WordFilter {
    dict := make(map[string]int)
    size := len(words)
    for i := 0; i < size; i++ {
        s := len(words[i])
        for j := 0; j <= s; j++ {
            for k := 0; k < s; k++ {
                key := words[i][:j] + "#" + words[i][k:]
                dict[key] = i
            }
        }
    }
    return WordFilter{dict}
}


func (this *WordFilter) F(prefix string, suffix string) int {
    key := prefix + "#" + suffix
    if idx, ok := (*this).dict[key]; ok {
        return idx
    }
    
    return -1
}


/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
```
