---
title: 0647. Palindromic Substrings
weight: 10
tags: [
	"Two Pointer",
	"String"
]
---

## Description
> Given a string s, return the number of palindromic substrings in it.
> 
> A string is a palindrome when it reads the same backward as forward.
> 
> A substring is a contiguous sequence of characters within the string.

## Solutions

### Two Pointer
回文串的解法都是相对比较单一的，这里还是沿用老套路：遍历字符串中的每个位置，检查以这个位置为中心，或者以这个位置和前一个位置为中心的字符是否能组成回文串，检查的时候把 res 传到子函数中去，这样就可以更新 res 。
```go
func countSubstrings(s string) int {
    var res int
    size := len(s)
    for i := 0; i < size; i++ {
        findPalindrome(s, i,i,size,&res)
        findPalindrome(s,i-1,i,size,&res)
    }
    
    return res
}

func findPalindrome(s string, left,right,size int, res *int) {
    for left >= 0 && right < size {
        if s[left] != s[right] {
            break
        }
        
        *res++
        left--
        right++
    }
}
```
