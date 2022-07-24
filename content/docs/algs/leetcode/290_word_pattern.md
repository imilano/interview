---
title: 0290. Word Pattern
weight: 10
tags: [
	"Hash Table",
]
---
## Description
> Given a pattern and a string s, find if s follows the same pattern.
> 
> Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

## Solutions
### Hash Table
简单题，只需要使用哈希表即可。需要注意的是，因为讲的 pattern 相同，所以意味着这是一个双射，只写单射是不行的。
```go
func wordPattern(pattern string, s string) bool {
    sps := strings.Split(s, " ")
    m,n := len(pattern), len(sps)
    if m == 0 && n == 0 {
        return true
    }
    
    if m == 0 || n == 0 || m != n {
        return false
    }
    
	// 双射，不仅意味着需要 d1 匹配 d2，还需要 d2 匹配 d1，所以只记录 pattern 中某个字符跟 s 中某个字符串匹配是不行的，
	// 所以还需要加上 d2 来做一个反向匹配。只有的 d1 没有 d2 的话，对于 pattern="abba", s="dog dog dog dog" 的情况就会出错
    d1, d2 := make(map[byte]string), make(map[string]bool)
    for i := 0; i < m; i++ {
        if str, ok := d1[pattern[i]]; ok && str != sps[i] {
            return false
        }
        
        d1[pattern[i]] = sps[i]
        d2[sps[i]] = true
        if len(d1) != len(d2) {
            return false
        }
    }
    
    
    return true
}
```
