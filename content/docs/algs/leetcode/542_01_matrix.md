---
title: 0542. 01 Matrix
weight: 10
tags: [
	"Matrix",
	"BFS"
]
---

## Description
> Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
> 
> The distance between two adjacent cells is 1.

## Solutions
### BFS
这里很容易能够想出 BFS 的解法，然后代码不难写。但是在题主写好并提交之后，在最后一个 case 上超时了，题主一看这个 case 的结构，心想不超时才怪呢....

主要思路就是，如果一个格子是 0，那么该格子的距离肯定是 0。而如果该格子是 1，那么其距离就是对该格子进行 BFS 的距离。超时代码如下：
```go
var dirs [][]int = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
func updateMatrix(mat [][]int) [][]int {
    // 初始化结果数组，将每个位置的值都设置为最大值
    rows, cols := len(mat), len(mat[0])
    res := make([][]int, rows)
    for idx, _ := range res {
        arr := make([]int, cols)
        for i := 0; i < cols; i++ {
            arr[i] = math.MaxInt
        }
        res[idx] = arr
    }
    
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if mat[i][j] == 0 {
                res[i][j] = 0
            } else {
                res[i][j] = bfs(mat, i, j, rows, cols)
            }
        }
    }
    
    return res
}


func bfs(mat [][]int, curRow, curCol, rows, cols int) int {
    var queue [][]int
    var res int
    queue = append(queue, []int{curRow, curCol})
    visited := make([][]bool, rows)
    for idx, _ := range visited {
        visited[idx] = make([]bool, cols)
    }
    visited[curRow][curCol] = true
    for len(queue) != 0 {
        res++
        size := len(queue)
        for i := 0; i < size; i++ {
            for _, dir := range dirs {
                x, y := queue[i][0] + dir[0], queue[i][1] + dir[1]
                if x < 0 || x >= rows || y < 0 || y >= cols || visited[x][y] {
                    continue
                }
                if mat[x][y] == 0 {
                    return res
                }
                
                if mat[x][y] == 1 {
                    queue = append(queue, []int{x, y})
                }
                
                visited[x][y] = true
            }
            
        }
        
        queue = queue[size:]
    }
    
    return res
}
```
题主想到的优化方式是，不是从每个值为 1 的格子为起点进行遍历，而是从每个值为 0 的格子为起点进行遍历，但是最后觉得这个方式其实并没有优化多少，这是对最后一个 case 而言是优化，对其它 case 而言可能耗时更长了。

这是题主在网上看到的优化方法：首先遍历以此数组，将每个值为 0 的点入栈，将值为 1 的点的值都设置为 MaxInt。然后以队列中的点为起点进行遍历，每当遇到边界或者当前当前路径值+1 大于格子当前值的话，则跳过，没有更新该点的必要，否则将周围点的值更新为当前值加 1，然后将周围点的坐标加入队列，然后继续遍历。（题主感觉这里其实也没有优化多少，无非是减少了 BFS 的数量，但是这个解法对最后一个 case 管用...）。最后 AC 的代码如下：
```go
var dirs [][]int = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
func updateMatrix(mat [][]int) [][]int {
    rows, cols := len(mat), len(mat[0])
    var queue [][]int
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
			// 将每个值为 0 的点都加入到队列中
            if mat[i][j] == 0 {
                queue = append(queue, []int{i, j})
            }
   
			// 如果值为 1， 则将其设置为 MaxInt
            if mat[i][j] == 1 {
                mat[i][j] = math.MaxInt
            }
        }
    }
    
    var res int
    for len(queue) != 0 {
        res++
        size := len(queue)
        for i := 0; i < size; i++ {
            for _, dir := range dirs {
				// 如果越界或者格子值大于当前路径值 res，则跳过，无需更新
                x, y := queue[i][0] + dir[0], queue[i][1] + dir[1]
                if x >= rows || x < 0 || y >= cols || y < 0 || res > mat[x][y] {
                    continue
                }
                
				// 否则，更新格子值，并且将当前格子加入到队列中
                mat[x][y] = res
                queue = append(queue, []int{x, y})
            }
        }
        
		// 缩小队列
        queue = queue[size:]
    }
    
    return mat
}


```
