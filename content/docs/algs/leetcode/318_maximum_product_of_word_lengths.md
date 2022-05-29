---
title: 0318. Maximum Product of Word Lengths
weight: 10
tags: [
	"String",
	"Hash Table"
	"Bit Manipulation"
]
---

## Description
> Given a string array words, return the maximum value of length(word[i]) * length(word[j]) where the two words do not share common letters. If no such two words exist, return 0.

## Solutions
{{< katex >}} \Omicron (n^2) {{< /katex >}} 的解法，相当于暴力求解了，因为限定了字符串中只会出现小写字符，而小写字符只有 26 个，那么就可以一个 32 位的 int 值 mask 来记录每个字符是否出现过，然后不同字符的 mask 之间进行相与，如果与的结果为 0，说明这两个字符没有重叠字符，那么就更新结果。
```go
func maxProduct(words []string) int {
    var res int
    size := len(words)
    dict := make(map[int]int, size)
    for i := 0; i < size; i++ {
        for _, r := range words[i] {
            dict[i] |= 1 << (int(r) - int('a'))
        }
        
        for j := 0; j < i; j++ {
            if (dict[i] & dict[j]) == 0 {
                res = max(res, len(words[i]) * len(words[j]))
            }
        }
    }
    
    return res
}


func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}
```
