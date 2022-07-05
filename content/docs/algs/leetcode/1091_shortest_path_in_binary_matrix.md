---
title: 1091. Shortest Path in Binary Matrix
weight: 10
tags: [
	"Matrix",
	"DFS",
	"BFS"
]
---
## Description
> Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix. If there is no clear path, return -1.
> 
> A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right cell (i.e., (n - 1, n - 1)) such that:
> 
> - All the visited cells of the path are 0.
> - All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they share an edge or a corner).
> The length of a clear path is the number of visited cells of this path.

## Solutions
### BFS
 这个题很容易可以想到是使用 BFS，是迷宫遍历的变体。但是这里跟一般的迷宫遍历还是稍微有点区别：
 - 迷宫遍历只能往四个方向，而这里可以往八个方向。也就是说，千万别忘了在 dirs 中加上 [1,-1] 和 [-1, 1]
 - 对 res 进行加 1 操作的时机不同。仔细看代码，这里是跟一般的迷宫遍历很不一样的。这里可以这么理解，普通的迷宫每移动一步都需要加 1，而这一题相当于你在当前位置所能够到的所有位置都只算做一个 move。
 - 在 string 和 数字之间进行转换的时候，一定要使用标准转换函数，而不是使用 string 方法将数字转换为字符串，这样的转换结果是不对的。

下面就是题主初始写出来的答案，结果总是不能 AC：
```go
type Cords struct {
    X int
    Y int
}
func shortestPathBinaryMatrix(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    if m == 0  || n == 0 || grid[0][0] != 0 {
        return -1
    }
    
    var queue []Cords = []Cords{{0,0}}
    var dirs [][]int = [][]int{{1,0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}}
    res := 1
    grid[0][0] = 1
    for len(queue) != 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            for _, dir := range dirs {
                x, y := queue[i].X + dir[0], queue[i].Y + dir[1]
                if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != 0 {
                    continue
                }
                
                // res 不应该在这里加，是否到达右下角也不应该在这里进行判断
                res++
                if x == m -1 && y == n - 1 {
                    return res
                }
                
                // 标记该点已经被访问过
                grid[x][y] = 1
                queue = append(queue, Cords{x,y})
            }
        }
    }
    
    return -1
}
```
```go
func shortestPathBinaryMatrix(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    if m == 0  || n == 0 || grid[0][0] != 0 {
        return -1
    }
    
    var queue [][]int = [][]int{{0,0}}
    // 注意点 1，需要 8 个方向
    var dirs [][]int = [][]int{{1,0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
    visited := map[string]bool{"00": true}
    var res int
    for len(queue) != 0 {
        // 注意点 2，注意 res 自增的位置
        res++
        size := len(queue)
        for i := 0; i < size; i++ {
            // 注意点 3， 注意什么时候才需要对坐标进行判断
            if queue[i][0] == m -1 && queue[i][1] == n - 1 {
                return res
            }
            for _, dir := range dirs {
                x, y := queue[i][0] + dir[0], queue[i][1] + dir[1]
                // 注意点 4， 不能使用 string 来对数字进行转换
                key := strconv.Itoa(x) + strconv.Itoa(y)
                if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 1 || visited[key] {
                    continue
                }
                
                visited[key] = true
                queue = append(queue, []int{x, y})
            }
        }
        
        queue = queue[size:]
    }
    
    return -1
}
```
