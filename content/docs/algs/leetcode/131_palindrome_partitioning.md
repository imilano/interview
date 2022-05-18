---
title: 131. Palindrome Partitioning
weight: 10
---
## Description
> Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.
> 
> A palindrome string is a string that reads the same backward as forward.


## Solutions

简单递归回溯即可。
```go
func partition(s string) [][]string {
    var res [][]string
    size := len(s)
    if size == 0 {
        return res
    }
    
    var cur []string
    helper(s,0,size, cur, &res)
    return res
}

func helper(s string, start, size int, cur []string, res *[][]string) {
    if start >= size {
        tmp := make([]string, len(cur))
        copy(tmp, cur)
        *res = append(*res, tmp)
        return
    }
    
    
    for i := start; i < size; i++ {
        if isPalindrome(s[start: i+1]) {
            cur = append(cur, s[start:i+1])
            helper(s, i+1, size, cur, res)
            cur = cur[:len(cur)-1]
        }
    }
}

func isPalindrome(s string) bool {
    size := len(s)
    left, right := 0, size -1
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    
    return true
}
```
