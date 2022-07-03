---
title: 0130. Surrounded Regions
weight: 10
tags: [
	"DFS",
	"BFS",
	"Matrix",
	"Graph"
]
---

## Description
> Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.
> 
> A region is captured by flipping all 'O's into 'X's in that surrounded region.

## Solutions
### DFS
这个题的难点在于，对于从边上的'O'为起点的所有可达的'O'，在最后的结果中要保留。那么也就是所，在遍历中，需要对从边上的'O'可达到的所有'O'做一个特殊处理。这里首先遍历四条边，如果当前点是'O'，那么从当前点开始进行DFS，对于每个可达的'O'点，都将其设置为'Y'，表示这个点不需要处理。上面的处理结束之后，剩下的'O'都是需要处理的'O'，那么只需要再遍历一遍矩阵，将每个‘O’设置为'X',同事将那么设置为'Y'的点还原为'O'即可。
```go
var dirs [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func solve(board [][]byte)  {
    rows, cols := len(board), len(board[0])
    // 先从边上的每个 'o'开始遍历，将其可达的所有 'o'都标记为 'y'
    for i := 0; i < rows; i++ {
        if board[i][cols-1] == 'O' {
            dfs(board, i, cols-1, rows, cols)
        }
        if board[i][0] == 'O' {
            dfs(board, i, 0, rows, cols)
        }
    }
    
    for i := 0; i < cols; i++ {
        if board[0][i] == 'O' {
            dfs(board, 0, i, rows, cols)
        }
        
        if board[rows-1][i] == 'O' {
            dfs(board, rows-1, i, rows, cols)
        }
    }
    

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            // 对于剩余的每个 'o'，将其变为 'x'
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
            // 将每个'y'还原为'o'
            if board[i][j] == 'Y' {
                board[i][j] = 'O'
            }
        }
    }
    
}


func dfs(board [][]byte, row, col, rows, cols int) {
    if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] == 'X' || board[row][col] == 'Y' {
        return 
    }
    
    if board[row][col] == 'O' {
        board[row][col] = 'Y'
    }
    
    
    for _, dir := range dirs {
        dfs(board, row+dir[0], col + dir[1], rows, cols)
    }   
}
```
