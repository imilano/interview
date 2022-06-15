---
title: 0240. Search a 2D Matrix II
weight: 10
tags: [
	"Matrix",
	"Search",
	"Binary Search"
]
---
## Description
> Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:
> 
> Integers in each row are sorted in ascending from left to right.
> Integers in each column are sorted in ascending from top to bottom.


## Solutions

### Binary Search
这个题使用跟 74 题一样的解法完全可以解决。
> 这题的二分解法真的比较巧妙。这里首先具备了两个特征：每一行从左向右递增，每一列从上到下递增。那么我们可以从第一列最后一个元素开始进行搜索，如果`target` 比当前元素要小，那么如果`target` 要存在，就只能在当前行上边，于是`row-1`；如果`target`比当前元素要大，那么如果该元素存在的话，智能在当前列右侧，于是`col+1`。按照以上规则进行遍历即可。
```go
func searchMatrix(matrix [][]int, target int) bool {
    m,n := len(matrix), len(matrix[0])   
    row, col := m - 1, 0
    for row >= 0 && col < n {
        if matrix[row][col] == target {
            return true
        }
        
        if matrix[row][col] < target {
            col++
        }
        if col < n && matrix[row][col] > target {
            row--
        }
    }
    
    return false
}
```
