---
title: 0474. Ones and Zeroes
weight: 10
tags: [
	"String",
	"DP"
]
---

## Description

> You are given an array of binary strings strs and two integers m and n.
> 
> Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.
> 
> A set x is a subset of a set y if all elements of x are also elements of y.


## Solutions
这里很难想到要用动态规划吧。假设 `dp[i][j]` 表示一个能包含 i 个 0 和 j 个 1 的子集中的字符串个数。则当扫描到字符串`str[i]`时，统计其中出现0 位 zeros，1 为 ones。则 {{< katex >}}dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1){{< /katex >}}。

```go
 func findMaxForm(strs []string, m int, n int) int {
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int, n+1)
    }
    
    
    for _, str := range strs {
        zeroes, ones := count(str)
		// 注意领会这里，也就子问题的方向，应该是从 dp[0][0] 到 dp[zeroes][ones]
        for i := m; i >= zeroes; i-- {
            for j := n; j >= ones; j-- {
                dp[i][j] = max(dp[i][j], dp[i-zeroes][j-ones]+1)    
            }
        }
    }
    
    return dp[m][n]
}

func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}

func count(str string) (int, int) {
    var ones, zeroes int
    for idx, _ := range str {
        if str[idx] == '0' {
            zeroes++
        } else {
            ones++
        }
    }
    
    return zeroes, ones
}
```
