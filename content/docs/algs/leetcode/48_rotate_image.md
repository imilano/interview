---
title: 0048. Rotate Image
weight: 10
---

## Description

> You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

> You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.


## Solutions

补充几个矩阵变换公式：
	- 主对角线对折： arr[i][j] = arr[j][i]
	- 副对角线对折： arr[i][j] = arr[n-j-1][n-i-1]
		```go
		for i := 0; i < n; i++ {
			for j := 0; j < n-i; j++ {
				matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
			}
		}
		```
	- 横向对折： arr[i][j] = arr[n-i-1][j]
	- 纵向对折： arr[i][j] = arr[i][n-j-1]
  
```go
// 先左右对折，再根据副对角线对折
func rotate(matrix [][]int)  {
    m := len(matrix)
    
    // 左右对折
    for i := 0; i < m; i++ {
        left, right := 0, m-1
        for left < right {
            matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
            left++
            right--
        }
    }
    
    // 副对角线对折
    for i := 0; i < m ;i++ {
        for j := 0; j < m-i ;j++ {
            matrix[i][j], matrix[m-j-1][m-i-1] = matrix[m-j-1][m-i-1], matrix[i][j]
        }
    }

}
```
