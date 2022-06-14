---
title: 0583. Delete Operations for Two Strings
weight: 10
tags: [
	"String",
	"DP"
]
---

## Description
> Given two strings word1 and word2, return the minimum number of steps required to make word1 and word2 the same.
> 
> In one step, you can delete exactly one character in either string.

## Solutions

### DP
 求最少的删除步数来使两个字符串一致，其实就是求两个字符串的最长公共子序列（注意，子序列意味着可以不连续），最少的步数就是二者的长度之和减去最长公共子序列的两倍。所以问题就是如何求两个字符串的最长公共子序列。定义 `dp[i][j]` 表示 `str1`（假设长度为`m`）的前`i`个字符和`str2`（假设长度为`n`）的前`j`个字符的最长公共子序列的长度，根据这个定义，我们最求所求的值就是`dp[m][n]`，那么我们应该申请的数组就是 `(m+1) * (n+1)`的二维数组。
 
 当 `str1[i]` 和  `str2[j]` 相等时，`dp[i][j]` 的值取决于`str1`的前`i-1`个`i`个字符和`str2`的前`j`个字符的最长公共子序列，也就是说， `dp[i][j] = dp[i-1][j-1] + 1`;
 当 `str1[i]` 和 `str2[j]` 不相等时， 由于所求的时最大长度，`dp[i][j]` 的值就继承于 `dp[i-1][j]` 和 `dp[i][j-1]`中的较大者。
 而对于初始条件，可以轻松得到， `dp[0][j] = dp[i][0] = 0`。
 ```go
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    if m == 0 || n == 0 {
        return 0
    }
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int, n+1)
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1 
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    
    return m + n - 2 * dp[m][n]
}

func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}

 ```
