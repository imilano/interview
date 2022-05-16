package leetcode

/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
*/

// 类似于图的区域染色,其实就是求整个图能经过几次 DFS。举例来说，每次 DFS 都赋予一个唯一值，然后使用 DFS 遍历每一个值为 1 的点，如果该点能够被遍历到，则将该点设为该唯一值。
// 最后计算唯一值的个数，如果可能的话，还要减去 0 的个数，所得数即为所求。
// 爆内存了（哭
// func numIslands(grid [][]byte) int {
//     m,n := len(grid), len(grid[0])

// 	uniqueID := byte('0' - 1)
// 	for i := 0; i < m;i++ {
// 		for j := 0; j < n;j++ {
//             if grid[i][j]  <= byte('0') {
// 				continue
// 			}
// 			numIslandsDFS(&grid, uniqueID,i,j,m,n)
// 			uniqueID -= 1
// 		}
// 	}

// 	return int(byte('0'-1) - uniqueID)
// }

// func numIslandsDFS(grid *[][]byte, id byte, si, sj,m,n int) {
//     if si < 0 || sj < 0 || si >= m || sj >= n || (*grid)[si][sj] <= byte('0') {
// 		return
// 	}
// 	(*grid)[si][sj] = id
// 	numIslandsDFS(grid, id, si+1, sj, m ,n)
// 	numIslandsDFS(grid, id, si-1, sj, m ,n)
// 	numIslandsDFS(grid, id, si, sj+1, m ,n)
// 	numIslandsDFS(grid, id, si, sj-1, m ,n)
// }

// 然而这样就不会爆内存，其实这里的实现跟上面的实现几乎一样，那为何爆内存呢？
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == byte('0') || visited[i][j] {
				continue
			}
			numIslandsDFS(&grid, &visited, i, j, m, n)
			res += 1

		}
	}

	return res
}

func numIslandsDFS(grid *[][]byte, visited *[][]bool, si, sj, m, n int) {
	if si < 0 || sj < 0 || si >= m || sj >= n || (*grid)[si][sj] == byte('0') || (*visited)[si][sj] {
		return
	}
	(*visited)[si][sj] = true
	numIslandsDFS(grid, visited, si+1, sj, m, n)
	numIslandsDFS(grid, visited, si-1, sj, m, n)
	numIslandsDFS(grid, visited, si, sj+1, m, n)
	numIslandsDFS(grid, visited, si, sj-1, m, n)
}
