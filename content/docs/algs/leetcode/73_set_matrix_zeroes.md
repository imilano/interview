---
title: 0073. Set Matrix Zeroes
weight: 10
tags: [
    "Memorial"
]
---

## Descriptioin

> Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
> 
> You must do it in place.


## Solutions

这里需要注意的是，因为要求不使用额外的内存空间，所以我们就要有这种充分利用原有空间的觉悟。这里的思想是，先检查第一行第一列是否有 0，如果有，则将相应的标志设置为 true，然后使用第一行和第一列来做标记。从第二行第二列开始遍历数组，当 matrix[i][j] = 0 时，则将该行对应的第一个数设置为 0，将该列对应的第一个数设置为 0。遍历结束之后再遍历一次数组，也是从第一行和第一列开始，如果改行或者该列的第一个元素为 0， 则将相应的相应的格子设置为 0.遍历结束后，再根据刚开始设置的行 flag 和列 flag 判断是否应该将第一行和第一列设置为 0.

```go
func setZeroes(matrix [][]int)  {
    m, n := len(matrix), len(matrix[0])
    var rowZero, colZero bool
    
    // 检查第一列是否包含 0
    for i := 0; i < m ;i++ {
        if matrix[i][0] == 0 {
            colZero = true
            break
        }
    }
    
    // 检查第一行是否包含 0
    for i := 0; i < n; i++ {
        if matrix[0][i] == 0 {
            rowZero = true
            break
        }
    }
    
    
    // 检查剩余行和列是否包含 0，如果包含，则将相应的第一行位置设为 0，第一列位置设为 0
    for i := 1; i < m ;i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    
    // 根据第一行第一列的检测结果，将相应的行和列设置为 0
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

    // 如果第一行有 0， 则将第一行设置为 0
    if rowZero {
        for i := 0; i < n; i++ {
            matrix[0][i] = 0
        }
    }
    
    // 如果第一列有 0， 则将第一列设置为 0
    if colZero {
        for i := 0; i < m ;i++ {
            matrix[i][0] = 0
        }
    }
}
```
