---
title: 0867. Transpose Matrix
weight: 10
tags: [
	"Matrix",
	"Array"
]
---
## Description
> Given a 2D integer array matrix, return the transpose of matrix.
> 
> The transpose of a matrix is the matrix flipped over its main diagonal, switching the matrix's row and column indices.

## Solutions
简言之就是需要行列倒置。
```go
func transpose(matrix [][]int) [][]int {
    m,n := len(matrix),len(matrix[0])
    res := make([][]int, n)
    for idx, _ := range res {
        res[idx] = make([]int, m)
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            res[j][i] = matrix[i][j]
        }
    }
    
    return res
}
```
