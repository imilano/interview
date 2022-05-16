package leetcode

/*
Given n nodes labeled from 0 to n - 1 and a list of undirected edges (each edge is a pair of nodes), write a function to find the number of connected components in an undirected graph.
*/

// 其实这题很快就能看出可以用 DFS 来做，很快就应该可以看出来是要求联通图的个数。 其实主要难点还是在于如何表示一个无向图。
// 这里我们使用邻接表来表示无向图，然后对每个节点深度遍历，遍历时记录这个节点是否已经被遍历过，如果已经被遍历过，则直接返回即可。
// 遍历时存储 DFS 遍历的次数，即可得到联通图的个数，
func countComponents(n int, edges [][]int) int {
	var res int
	visited := make([]bool, n)
	adjacentTable := make([][]int, n)
	for _, edge := range edges {
		adjacentTable[edge[0]] = append(adjacentTable[edge[0]], edge[1])
		adjacentTable[edge[1]] = append(adjacentTable[edge[1]], edge[0])
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			countComponentsDFS(adjacentTable, &visited, i)
			res += 1
		}
	}

	return res
}

func countComponentsDFS(adjacentTable [][]int, visited *[]bool, point int) {
	if (*visited)[point] {
		return
	}

	(*visited)[point] = true
	peers := adjacentTable[point]
	for _, peer := range peers {
		countComponentsDFS(adjacentTable, visited, peer)
	}
}

// for test
func CountComponents(n int, edges [][]int) int {
	return countComponents(n, edges)
}
