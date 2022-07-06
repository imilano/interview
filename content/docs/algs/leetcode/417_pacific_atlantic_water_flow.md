---
title: 0417. Pacific Atlantic Water Flow
weight: 10
tags: [
	"Matrix",
	"DFS"
]
---
## Description
> There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.
> 
> The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).
> 
> The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.
> 
> Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.

## Solutions
应该很容易就可以看出来，其实这是图搜索问题。

解决方法是，对于图中的每一个点进行遍历，检查这个点是否能够有到达太平洋的路径以及到达大西洋的路径。如果两边都能够到达，则说明这个点就是我们要求的结果之一。
但是上述方法可能会导致很大的计算量，首先我们可以使用 DFS 的优化算法来进行遍历，其次的话我们还可以换一种思路来实现。

由于直接搜索会导致大量的计算，而我们所要的结果路径必然是要以边缘点为最后节点的，那么我们可以从边缘点为开始节点进行搜索。分别创建两个二维数组，表示太平洋和大西洋。
然后对边缘进行 DFS，对于符合要求的点，将该数组对应点设置为 true，最后的结果就是两个数组均为 true 的节点。
```go
func pacificAtlantic(heights [][]int) [][]int {
	// 创建两个数组，分别代表从太平洋边缘出发能到达的点和从大西洋边缘出发能到达的点
    rows, cols := len(heights), len(heights[0])
    pacific, atlantic := make([][]bool, rows), make([][]bool, rows)
    for i := 0; i < rows; i++ {
        pacific[i] = make([]bool, cols)
        atlantic[i] = make([]bool, cols)
    }
    
	// 从边缘开始遍历，将所有能从太平洋边缘到达的点设置为 true，将所有能从大西洋边缘到达的点设置为 true
	// 到达意味着当前点的值要比前一个点要大。注意初始的 pre 要设置为最小值
    for i := 0; i < rows; i++ {
        dfs(heights, pacific, math.MinInt, i, 0, rows, cols)
        dfs(heights, atlantic, math.MinInt, i, cols - 1, rows, cols)
    }
    
    for i := 0; i < cols; i++ {
        dfs(heights, pacific, math.MinInt, 0, i, rows, cols)
        dfs(heights, atlantic, math.MinInt, rows-1, i, rows, cols)
    }
    
	// 从两大洋都能到达的点，才是我们要的结果
    var res [][]int
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if pacific[i][j] && atlantic[i][j] {
                res = append(res, []int{i, j})
            }
        }
    }
    
    return res
}

func dfs(heights [][]int, visited [][]bool, pre, curRow, curCol, rows, cols int) {
	// 对于越界的点，对于已经遍历过的点，对于当前值比前一个值要小的点，都需要再进行遍历
    if curRow < 0 || curRow >= rows || curCol < 0 || curCol >= cols || visited[curRow][curCol] || heights[curRow][curCol] < pre {
        return 
    }
    
	// 标记当前点已经遍历过，并开始遍历当前节点四周
    visited[curRow][curCol] = true
    dfs(heights, visited, heights[curRow][curCol], curRow + 1, curCol, rows, cols)
    dfs(heights, visited, heights[curRow][curCol], curRow - 1, curCol, rows, cols)
    dfs(heights, visited, heights[curRow][curCol], curRow, curCol + 1, rows, cols)
    dfs(heights, visited, heights[curRow][curCol], curRow, curCol - 1, rows, cols)
}
```
