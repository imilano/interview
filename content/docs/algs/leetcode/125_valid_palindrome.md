---
title: 125. Valid Palindrome
weight: 10
---

## Description
> A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
> 
> Given a string s, return true if it is a palindrome, or false otherwise.


## Solutions
简单题，去除字符之后直接判断即可。
```go
func isPalindrome(s string) bool {
    if len(s) <= 1 {
		return true
	}
	s = strings.ToLower(s)
	var str []rune
	for _, r := range s {
		if r  >= rune('a') && r <= rune('z') || r >= rune('0') && r <= rune('9'){
			str = append(str, r)
		}
	}

	size := len(str)
	if size <= 1{
		return true
	}

	left, right := 0, size -1
	for left < right {
		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}

	return true
}
```
