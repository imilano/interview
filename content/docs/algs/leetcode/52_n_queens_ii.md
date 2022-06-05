---
title: 0052. N Queens II
weight: 10
tags: [
    "Brute Force",
	"Backtracing"
]
---

## Description
> The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
> 
> Given an integer n, return the number of distinct solutions to the n-queens puzzle.

## Solutions
这题跟 51 题的解法其实是一致的，区别只是这里求的是一个计数值。
```go
func totalNQueens(n int) int  {
    // 将棋盘全部初始化为点
    var res int
    board := make([][]string, n)
    for idx, _ := range board {
        board[idx] = make([]string, n)
    }
    for i := 0; i < n ;i++ {
        for j := 0; j < n; j++ {
            board[i][j] = "."
        }
    }
    
    // 从第一行开始回溯
    backtrace(0, n, &res, &board)
    return res
}

func backtrace(curRow, rows int, res *int, board *[][]string) {
    // 递归终止条件
    if curRow == rows {
        *res += 1
        return
    }
    
    // 遍历这一行的每一列，检查当前棋盘是否合法，如果合法，则将当前格子置为 Queen
    for i := 0; i < rows; i++ {
        if isValid(*board, curRow, i, rows) {
            (*board)[curRow][i] = "Q"
            backtrace(curRow+1, rows, res, board)
            (*board)[curRow][i] = "."
        }
    }
}


func isValid(board [][]string, curRow int, curCol int, rows int) bool {
    // 查看这个列合不合法
    for i := 0; i < curRow; i++ {
        if board[i][curCol] == "Q" {
            return false
        }
    }
    
    // 查看主对角线合不合法
    for i, j := curRow - 1, curCol - 1; i >= 0 && j >= 0; i,j = i -1, j -1 {
        if board[i][j] == "Q" {
            return false
        }
    }
    
    
    // 查看副对角线合不合法
    for i, j := curRow - 1 , curCol + 1; i >= 0 && j < rows; i,j = i-1, j+1 {
        if board[i][j] == "Q" {
            return false
        }
    }
    
    return true
}
```
