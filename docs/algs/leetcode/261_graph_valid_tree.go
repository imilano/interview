package leetcode

/*
Given n nodes labeled from 0 to n-1 and a list of undirected edges (each edge is a pair of nodes), write a function to check whether these edges make up a valid tree.
*/

// 这里的问题就是给定一个图，验证这个图是不是一棵树。如果一个图需要是一棵树，那么需要同时满足以下两个条件：
// 	1. 不能存在环
// 	2. 这个图需要是连通的，也就是说，不能存在一个节点跟任何其它节点都不连通。
// 基于上述分析，这里可以有如下解决方式：
// 	1. 使用 DFS。首先建立一个图，然后DFS 遍历这个图，建立一个 visited 数组表示某个点是否被遍历过。如果 DFS 发现了环，则直接返回 false。
// 		DFS 遍历结束后，遍历 visited 数组，如果 visited 中有节点没有被访问过，则返回 false；否则，返回 true。
// 	2. 使用并查集。首先假设每个节点的父节点都是-1，然后根据给出的边的关系， 查询两个节点是的根节点，如果根节点相等，则说明存在环，返回 false；
// 		否则将两个节点 union。因为二叉树的边数量等于 n-1，最后只需要检查一下边是否等于 n-1 即可，若不等，则说明存在孤立节点。

func validTree(n int, edges [][]int) bool {
	roots := make([]int, n)
	for index, _ := range roots {
		roots[index] = -1
	}

	for _, edge := range edges {
		p, q := find(&roots, edge[0]), find(&roots, edge[1])
		if p == q {
			return false
		}

		roots[p] = q
	}

	return len(edges) == n-1
}

func find(roots *[]int, node int) int {
	for (*roots)[node] != -1 {
		node = (*&*roots)[node]
	}

	return node
}
