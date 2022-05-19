---
title: 0036. Valid Sudoku
weight: 10
---

## Description

> Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
> 
> 1. Each row must contain the digits 1-9 without repetition.
> 2. Each column must contain the digits 1-9 without repetition.
> 3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

## Solutions

一个是要注意如何快速创建三维数组，另一个是要注意如何进行坐标转换。
```go
func isValidSudoku(board [][]byte) bool {
    size := 9
    
    // 一次性将所有行和列创建完毕
    row, col := make([][]int, size), make([][]int, size)
    for i := 0; i < size; i++ {
        row[i], col[i] = make([]int, size), make([]int, size)
    }
    
    // 表示 9 个gird，么个 grid 有 9 个元素. 最后创建出来的是一个 grid[3][3][9] 的三维数组
    grid := make([][][]int, 3)
    for idx, _ := range grid {
        grid[idx] = make([][]int, 3)
        for i, _ := range grid[idx] {
            grid[idx][i] = make([]int, size)
        }
    }

	// 下面这种创建三维数组的方式太繁琐了
    // for i := 0; i < 3; i++ {
    //     grid[i] = make([][]int, 3)
    // }
    // for i := 0; i < 3; i++ {
    //     for j := 0; j < 3; j++ {
    //         grid[i][j] = make([]int, size)
    //     }
    // }
    
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            if board[i][j] == '.' {
                continue
            }
            
            // 之所以要减去 1，是为了将值控制在 0-8 之内，避免数组访问越界
            index := int(board[i][j] - '0' - 1)
            row[i][index]++
            col[j][index]++
            grid[i/3][j/3][index]++
            if row[i][index] > 1 || col[j][index] > 1 || grid[i/3][j/3][index] > 1 {
                return false
            }

        }
    }

    
    return true
}
```
