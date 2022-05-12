---
title: 54. Spiral Matrix
weight: 10
---

## Description

> Given an m x n matrix, return all elements of the matrix in spiral order.

## Solutions

其实就是顺时针旋转打印数组。
```go
func spiralOrder(matrix [][]int) []int {
    var res []int
    m, n := len(matrix), len(matrix[0])
    
    up, bottom, left, right := 0, m-1, 0, n - 1
    for up <= bottom && left <= right {
        for i := left; i <= right; i++ {
            res = append(res, matrix[up][i])
        }
        up++
        if up > bottom {
            break
        }
        
        for i := up; i <= bottom; i++ {
            res = append(res, matrix[i][right])
        }
        
        right--
        if right < left {
            break
        }
        
        for i := right; i >= left; i-- {
            res = append(res, matrix[bottom][i])
        }
        bottom--
        if bottom < up {
            break
        }
        
        for i := bottom; i >= up;i-- {
            res = append(res, matrix[i][left])
        }
        left++
        if left > right {
            break
        }
    }
    
    return res
}
```
