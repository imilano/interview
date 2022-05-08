---
title: 28. Implement strStr
weight: 10
---

## Description

> Implement strStr().
> 
> Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.


## Solutions

太简单了在，直接看代码即可。
```go
func strStr(haystack string, needle string) int {
    var pos int
    size := len(haystack)
    targetSize := len(needle)
    if size == 0 || targetSize == 0 {
        return pos
    }
    
    var start int
    
    pos = -1
    for start < size {
        if haystack[start] != needle[0] {
            start++
            continue
        }
        
        if start + targetSize <= size && haystack[start: start+targetSize] == needle {
            pos = start
            break
        }
        start++
    }
    
    return pos
}
```
