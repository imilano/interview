---
title: 0490. The Maze
weight: 10
tags: [
	"Matrix",
	"DFS",
	"BFS"
]
---

## Description
> There is a ball in a maze with empty spaces and walls. The ball can go through empty spaces by rolling up, down, left or right, but it won't stop rolling until hitting a wall. When the ball stops, it could choose the next direction.
> 
> Given the ball's start position, the destination and the maze, determine whether the ball could stop at the destination.
> 
> The maze is represented by a binary 2D array. 1 means the wall and 0 means the empty space. You may assume that the borders of the maze are all walls. The start and destination coordinates are represented by row and column indexes.

## Solutions
### DFS
这个题与一般的 DFS 题稍微不同的地方在于，这里并不是每次走一步，而是选定方向之后就一直走，直到遇到墙才停下来。但其实也没有变化很多，DFS 模板还是可以套上来用。

这里需要注意两个点：
- 首先，可以直接复用原数组来做标记，表示当前格子是否已经被遍历过。
- 其次，不需要对每个格子都做标记，实际上，只需要对边界的格子来做标记即可。
```go
var dirs = [][]int{{0,1}, {0, -1}, {1, 0}, {-1, 0}}

func hasPath(maze [][]int, start, destination int[]) bool {
	rows, cols := len(maze), len(maze[0])
	return dfs(maze, start[0], start[1], destination[0], destination[1], rows, cols)
}

func dfs(maze [][]int, starti, startj, destinationi, destinationj, rows, cols int) bool {
	// 如果可达，直接返回 true
	if starti == destinationi && startj == destinationj {
		return true
	}

	var res bool
	// 直接复用原数组，表示当前点已经遍历过
	maze[starti][startj] = -1
	for _, dir := range dirs {
		// 一直往一个方向走，直到越界或者遇到墙壁
		x, y := starti, startj
		for x >= 0 && x < rows && y >= 0 && y < cols && maze[x][y] != -1 {
			x += dir[0]
			y += dir[1]
		}

		// 上面的循环会导致当前的[x,y]处于非法格子处，所以需要回退一步
		x -= dir[0]
		y -= dir[1]

		// 当前格子没有遍历过，则以这个格子为起点，又开始往四面八方遍历
		// 其实可以看出来，这里会有一次重复遍历。比如当前遍历中你一直往下走，遇到了墙壁，
		// 那么下一次遍历中你可能就会往上走，然后又回到上一次遍历的起点。
		// 但是因为该起点已经被遍历过了，所以已经被标记过，因而不会再走入下面这个递归中
		//
		// 上面说的这个问题，可以将 dirs 换成一个 map，key 为遍历的方向，
		// 然后在下面的 if 语句中再加一个条件，这样就可以避免往原方向再进行一次遍历
		if maze[x][y] != -1 {
			res |= dfs(maze, x, y, destinationi, destinationj, rows, cols, cols)
		}
	}

	// 放回本次遍历的结果
	return res
}
```
