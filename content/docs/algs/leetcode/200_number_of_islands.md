---
title: 0200. Number of Islands
weight: 10
tags: [
	"DFS",
	"Matrix"
]
---

## Description
> Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
> 
> An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

## Solutions
### DFS
典型的 DFS 应用。这里跟求一个图中有多少个连通分量一样，因为每个方格只能是 '0' 或者 '1'，所以我们可以在遍历的时候直接把已经遍历过的地方都设为'0'。如果不想污染原数组，那么就用一个 map 来记录该位置是否被遍历过即可。
```go
func numIslands(grid [][]byte) int {
    row, col := len(grid), len(grid[0])
    var res int
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if grid[i][j] != '0' {
                res++
                dfs(grid, i, j, row, col)
            }
        }
    }
    
    return res
}

func dfs(grid [][]byte, row, col, rows, cols int) {
    if row >= rows || row < 0 || col >= cols || col < 0 || grid[row][col] == '0' {
        return
    }
    
    if grid[row][col] != '0' {
        grid[row][col] = '0'
        dfs(grid, row+1, col, rows, cols)
        dfs(grid, row-1, col, rows, cols)
        dfs(grid, row, col+1, rows, cols)
        dfs(grid, row, col-1, rows, cols)
    }
}
```
