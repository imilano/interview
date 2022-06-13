---
title: 0120. Triangle
weight: 10
tags: [
	"DP"
]
---
## Description
> Given a triangle array, return the minimum path sum from top to bottom.
> 
> For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

## Solutions

### Brute Force
题主首先想出了贪心的解法，但是因为”局部最优并不代表全局最优“，所以没能通过，不过还是贴一下代码：
```go
func minimumTotal(triangle [][]int) int {
    size := len(triangle)
    arr := make([]int, size+1)
    
    var start int
    for i := 1; i <= size; i++ {
        if start + 1 < len(triangle[i-1]) && triangle[i-1][start+1] < triangle[i-1][start] {
            start = start + 1
        }
        arr[i] = dp[i-1] + triangle[i-1][start]
    }
    
    return arr[size]
}
```

既然上面的贪心不行，那么使用递归来试试呢？于是楼主又写出了下面的解法：
```go
func minimumTotal(triangle [][]int) int {
    size := len(triangle)
    res := math.MaxInt
    helper(triangle, 0, 0, 0, size, &res)
    return res
}


func helper(triangle [][]int,pathSum int, curRow int, curCol int, size int, res *int) {
    if curRow >= size {
        *res = min(*res, pathSum)
        return
    }
    
    pathSum += triangle[curRow][curCol]
    helper(triangle, pathSum, curRow+1, curCol, size, res)
    helper(triangle, pathSum, curRow+1, curCol+1, size, res)
}

func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```
很明显，这是暴力枚举啊，结果不出意外的超时了，不过解法应该是没问题的。

### Dynamic Programming
这里还是要使用 DP 的方法来解，还是因为”局部最优解不代表全局最优解“，所以我们这里需要计算每一行的每个数的自顶向下的路径和，然后从最后一行的结果中选出最小的那个数来作为我们的结果。既然思路是这样，那么状态转移方程大概就是 `dp[i][j] += min(dp[i-1][j-1], dp[i-1][j)`，这里我们可以直接复用 triangle 这个数组，直接在数组上进行操作。另外还需要注意的是，迭代只能从第二行开始，并且每一行中的第一个数和最后一个数需要进行特殊处理。综合以上，代码如下：
```go
func minimumTotal(triangle [][]int) int {
    size := len(triangle)
    for i := 1; i < size; i++ {
        rowSize := len(triangle[i])
        for j := 0; j < rowSize; j++ {
            if j == 0 {
                triangle[i][j] += triangle[i-1][j]
            } else if j == rowSize - 1 {
                triangle[i][j] += triangle[i-1][j-1]
            } else {
                triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
            }
        }
    }
    
    res := math.MaxInt
    rowLen := len(triangle[size-1])
    for i := 0; i < rowLen; i++ {
        res = min(res, triangle[size-1][i])
    }
    
    return res
}


func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```
