---
title: 1062. Longest Repeating Substring
weight: 10
tags: [
	"DP",
	"String"
]
---
## Description
> Given a string S, find out the length of the longest repeating substring(s). Return 0 if no repeating substring exists.

## Solutions
### Dynamic Programming
这种什么最长、最大的题目，当然是使用 DP 来解啦，不过题主想了好一会也没想到该怎么定义状态方程，只好上网去看各路大神怎么搞的：这里定了 dp[i][j] 表示 s 中以第 i 个字符为结尾的子串为和以第 j 个字符为结尾的子串的最大公共**后缀**的长度。当s[i] == s[j] 时，dp[i][j] = dp[i-1][j-1] + 1；否则，dp[i][j] = 0。在 dp 更新的过程中，不断保存最大值即可。

这里 DP 的递推式跟「718. Maximum Length of Repeated Subarray」是一致的。
```go
func longestRepeatingSubstring(s string) int {
	res,size := math.MinInt, len(s)
	dp := make([]int, size + 1)
	for idx, _ := range dp {
		dp[idx] = make([]int, size + 1)
	}

	for i := 1; i <= size; i++ {
		for j := 1; j <i;j++ {
			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			res = max(res, dp[i][j])
		}
	}

	return res
}

func max(a,b int) int {
	if a < b {
		return b
	}

	return a
}
```
