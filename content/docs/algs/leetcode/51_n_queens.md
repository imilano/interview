---
title: 0051. N Queens
weight: 10
tags: [
    "Brute Force",
	"Backtracing"
]
---

## Description
> The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
> 
> Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.
> 
> Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.
## Solutions

### Backtrace
N-皇后问题，就是一个棋盘上放置 N 个皇后，然后要保证当前行、当前列、主对角线和副对角线上都没有皇后。很惭愧，本科和硕士阶段都学过问题，但是现在都忘了。

N-皇后问题一般使用回溯来解决。
```go
func solveNQueens(n int) [][]string {
    // 将棋盘全部初始化为点
    var res [][]string
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

func backtrace(curRow, rows int, res *[][]string, board *[][]string) {
    // 递归终止条件
    if curRow == rows {
        // 注意这里不能使用 make 来做预创建
        var tmp []string
        for i := 0; i < curRow; i++ {
            t := strings.Join((*board)[i],"")
            tmp = append(tmp, t)
        }
        
        *res = append(*res, tmp)
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
