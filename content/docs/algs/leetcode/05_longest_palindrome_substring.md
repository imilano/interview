---
title: 0005. Longest Palindrome Substring
weight: 10
tags: [
	"String",
	"Longest Common Substring",
	"DP",
	"Two Pointer"
]
---
## Description
> Given a string s, return the longest palindromic substring in s.

## Solutions
### 中心扩散
题主首先想到的是中心扩散方法。核心思想就是遍历以此字符串，然后对于遍历到的每个位置，找出以这个位置为中心能找到的最长回文串。需要注意的是，一个回文串可能是偶数长度，也可能是奇数长度。如果是奇数长度，那么就需要以 i 为中心来对左右进行扩散；而如果是偶数长度，那么就需要以 i 和 i-1 为中心来进行扩散。
```go
func longestPalindrome(s string) string {
	size := len(s)
	if size <= 1 {
		return s
	}

	var res string
	for i := 0; i < size; i++ {
        // 最长回文串可能会出现在以 i 为中心对称的子串上，也可能出现在以 i 和 i - 1 为中心的子串上 
		r1 := longestPalindromeHelper(s, i, i, size)
		r2 := longestPalindromeHelper(s, i-1, i, size)
		res = getMaxString(r1, r2, res)
	}

	return res
}

func longestPalindromeHelper(s string, left, right, size int) string {
	var res string
	for left >= 0 && right < size {
		if s[left] == s[right] {
            // res 放在这里更新，这样的话，就不用写判断 left 和 right 是否有效的逻辑了
			res = s[left : right+1]
			left--
			right++
		} else {
			break
		}
	}

	return res
}

func getMaxString(a, b, s string) string {
	if len(a) < len(b) {
		if len(b) < len(s) {
			return s
		} else {
			return b
		}
	} else {
		if len(a) < len(s) {
			return s
		} else {
			return a
		}
	}
}
```

### Longest Common Stirng
将输入字符串逆转之后，原问题就可以转换为求最长公共子串(leetcode 第 718 题)的问题。
```go
```

### Dynamic Programming
定义 dp[j][i] 表示
```go
func longestPalindrome(s string) string {
	size := len(s)
    dp := make([][]bool, size)
    for idx, _ := range dp {
        dp[idx] = make([]bool, size)
    }
    
    left, step := 0, 1
    for i := 0; i < size; i++ {
        dp[i][i] = true
        for j := 0; j < size; j++ {
            dp[j][i] = s[i] == s[j] && (i-j <= 1 || dp[j+1][i-1])
            if dp[j][i] && step < i - j + 1 {
                left = j
                step = i - j + 1
            }
        }
    }
    
    return s[left: left+step]
}
```
