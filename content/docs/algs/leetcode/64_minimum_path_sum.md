---
title: 0064. Minimum Path Sum
weight: 10
tags: [
	"DP",
	"Matrix"
]
---

## Description
> Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
> 
> Note: You can only move either down or right at any point in time.

## Solutions

### DP
这题跟第 `120. Triangle` 解法几乎一样。相比之下，这题要简单一些。这里很明显需要使用 DP，并且状态转移方程也是很明确的 `dp[i][j] += min(dp[i-1][j], dp[i][j-1])`，也就是说，当前位置的值只能从左边移动过来或者上边移动过来。特例情况就是第一行跟第一列，这个地方只有一种移动方向，所以需要特别处理一下。
```go
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    // corner case
    for i := 1; i < m ;i++ {
        grid[i][0] += grid[i-1][0]
    }
    for i := 1; i < n; i++ {
        grid[0][i] += grid[0][i-1]
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            grid[i][j] += min(grid[i-1][j], grid[i][j-1])
        }
    }
    
    return grid[m-1][n-1]
}

func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```

当然，如果你不想污染原数组，那么也可以创建一个新的二维数组来做 DP 数组。
