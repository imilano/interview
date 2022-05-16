package leetcode

import (
	"math"
)

/*
There is an m x n rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges.

The island is partitioned into a grid of square cells. You are given an m x n integer matrix heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean.

Return a 2D list of grid coordinates result where result[i] = [ri, ci] denotes that rain water can flow from cell (ri, ci) to both the Pacific and Atlantic oceans.
*/

// 应该很容易就可以看出来，其实这是图搜索问题。
// 解决方法是，从对于图中的每一个点进行遍历，检查这个点是否能够有到达太平洋的路径以及到达大西洋的路径。如果两边都能够到达，则说明这个点就是我们要求的结果之一。
// 但是上述方法可能会导致很大的计算量，首先我们可以使用 DFS 的优化算法来进行遍历，其次的话我们还可以换一种思路来实现。
// 由于直接搜索会导致大量的计算，而我们所要的结果路径必然是要以边缘点为最后节点的，那么我们可以从边缘点为开始节点进行搜索。分别创建两个二维数组，表示太平洋和大西洋。
// 然后对边缘进行 DFS，对于符合要求的点，将该数组对应点设置为 true，最后的结果就是两个数组均为 true 的节点。
func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pacific, atlantic := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		pacificAtlanticDFS(heights, &pacific, math.MinInt, i, 0, m, n)
		pacificAtlanticDFS(heights, &atlantic, math.MinInt, i, n-1, m, n)
	}

	for i := 0; i < n; i++ {
		pacificAtlanticDFS(heights, &pacific, math.MinInt, 0, i, m, n)
		pacificAtlanticDFS(heights, &atlantic, math.MinInt, m-1, i, m, n)
	}

	var res [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func pacificAtlanticDFS(heights [][]int, visited *[][]bool, pre int, si int, sj int, m, n int) {
	if si < 0 || sj < 0 || si >= m || sj >= n || (*visited)[si][sj] || heights[si][sj] < pre {
		return
	}

	(*visited)[si][sj] = true
	pacificAtlanticDFS(heights, visited, (*&heights)[si][sj], si+1, sj, m, n)
	pacificAtlanticDFS(heights, visited, (*&heights)[si][sj], si-1, sj, m, n)
	pacificAtlanticDFS(heights, visited, (*&heights)[si][sj], si, sj+1, m, n)
	pacificAtlanticDFS(heights, visited, (*&heights)[si][sj], si, sj-1, m, n)
}
