---
title: 1047. Remove All Adjacent Duplicates in String
weight: 10
tags: [
	"Stack"
]
---
## Description
> You are given a string s consisting of lowercase English letters. A duplicate removal consists of choosing two adjacent and equal letters and removing them.
> 
> We repeatedly make duplicate removals on s until we no longer can.
> 
> Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.



## Solutions
很容易想到用栈。这里需要注意的是，对于"abbaca"这样的字符，移除之后的结果是 “ca”，而不是“abaca”喔，这里的删除是指一旦连续两个字符相同，那么把这两个字符都删去，而不是只保留一个。
```go
func removeDuplicates(s string) string {
    size := len(s)
    if size <= 1 {
        return s
    }
    
    var stack []rune
    for _, r := range s {
        size = len(stack)
        if size == 0 || stack[size-1] != r {
            stack = append(stack, r)
            continue
        }
        
        stack = stack[:size-1]
    }
    
    return string(stack)
}
```
