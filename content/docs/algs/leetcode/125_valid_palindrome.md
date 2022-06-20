---
title: 0125. Valid Palindrome
weight: 10
tags: [
	"String",
	"Two Pointer"
]
---

## Description
> A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
> 
> Given a string s, return true if it is a palindrome, or false otherwise.


## Solutions
### Two Pointer
简单题，去除非数字字母的字符之后直接判断即可。
```go
func isPalindrome(s string) bool {
    var rs []rune
    for _, r := range s {
        // skip none-alphanumeric characters
        if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >='0' && r <= '9') {
            continue
        }
        
        // convert upper case to lower case
        if r >= 'A' && r <= 'Z' {
            r += 'a' - 'A'
        }
        
        // append to rs
        rs = append(rs, r)
    }
    
    
    // check if it is palindromic string
    left, right := 0, len(rs)-1
    for left < right {
        if rs[left] != rs[right] {
            return false
        }
        
        left++
        right--
    }
    
    return true
}
```
