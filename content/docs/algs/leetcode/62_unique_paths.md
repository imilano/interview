---
title: 0062. Unique Paths
weight: 10
tags: [
    "DP",
    "Math"
]
---
## Description
> There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.
> 
> Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.
> 
> The test cases are generated so that the answer will be less than or equal to 2 * 109.

## Solutions

### DP

假设 dp[i][j] 表示到达[i,j] 这个位置的路径数，则很明显，dp[i][j] 的值只会受到 dp[i-1][j] 和 dp[i][j-1] 影响，进而可以进一步得到递推公式 dp[i][j] = dp[i-1][j] + dp[i][j-1]。 base case 的话，第一行第一列的路径应该为 1， 即 dp[i][0] = dp[0][j] = 1。最后 dp[m-1][n-1]即为所求。

```go
func uniquePaths(m int, n int) int {
	// corner case
    if m == 1 || n == 1 {
        return 1
    }
    
    dp := make([][]int, m)
    for idx, _ :=range dp {
        dp[idx] = make([]int, n)
    }
    
    //corner case
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    
    for i := 0; i < n; i++ {
        dp[0][i] = 1
    }
    
    for i := 1; i < m ; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    
    return dp[m-1][n-1]
}
```


### Math
这里也可使用排列组合的方法来做。对于一个 m 行 n 类的矩阵，要想从左上角到达右下角，那么必须需要向下走 m-1 步，向右走 n-1 步， 那么也就是说，我们需要从 m-1+n-2 的步数里面，跳出 m -1 个向下的步数和 n-1 个向右的步数。既然想到这里，那么这个问题就很好解决了，直接使用排列组合公式即可。

```go
func uniquePaths(m int, n int) int {
	// corner case
    if m == 1 || n == 1 {
        return 1
    }
    
    // 确保 n 是较小的那一个，不然下面的算法会导致溢出，进而导致错误的结算结果
    if n > m {
        n,m = m, n
    }
    
    // 计算排列组合 C_n_(m-1+n-1)
    res := 1
    for i := 0; i < n-1; i++ {
        res *= (n-1+m-1-i)
        res /= (i+1)
    }
    
    return res
}
```


