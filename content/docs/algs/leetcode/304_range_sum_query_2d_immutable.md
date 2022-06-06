---
title: 0304. Range Sum Query 2D - Immutable
weight: 10
tags: [
	"DP",
	"Prefix Sum"
]
---

## Description
> Given a 2D matrix matrix, handle multiple queries of the following type:
> 
> Calculate the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
> Implement the NumMatrix class:
> 
> NumMatrix(int[][] matrix) Initializes the object with the integer matrix matrix.
> int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).

## Solutions

### Brute Force
这是一个会超时的解法。
```go
type NumMatrix struct {
    matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    rows, cols := len(matrix), len(matrix[0])
    return NumMatrix{matrix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    var res int
    for i := row1; i <= row2; i++ {
        for j := col1; j <= col2; j++ {
            res += (*this).matrix[i][j]
        }
    }
    
    return res
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
 ```

 ### Prefix Sum
 这是一个参考 303 题使用前缀和来求解的方法。使用了前缀和之后，如果我们想要快速求出`(r1, c1)`到`(r2, c2)`的矩形区间时，只需`prefix[r2][c2] - prefix[r2][c1 - 1] - prefix[r1 - 1][c2] + prefix[r1 - 1][c1 - 1]`即可，而又因为我们使用辅助行和辅助列，prefix 矩阵的位置相比原有位置每行每列都各新增了 1，所以最后的求解表达式就是`prefix[r2+1][c2+1] - prefix[r2+1][c1] - prefix[r1][c2+1] + prefix[r1][c1]`。
 ```go
type NumMatrix struct {
    prefix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m,n := len(matrix), len(matrix[0])
    prefix := make([][]int, m+1)
    for idx, _ := range prefix {
        prefix[idx] = make([]int, n+1)
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
                 prefix[i][j] = prefix[i][j-1] + prefix[i-1][j] - prefix[i-1][j-1] + matrix[i-1][j-1]   
        }
    }
    
    return NumMatrix{prefix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return (*this).prefix[row2+1][col2+1] - (*this).prefix[row2+1][col1] - (*this).prefix[row1][col2+1] + (*this).prefix[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

 ```
