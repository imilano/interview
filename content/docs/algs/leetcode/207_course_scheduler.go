package leetcode

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.
*/

// 其实应该和容易能想到，这题考的就是图中的环检测。
// 关于图中的环检测，可以使用： 1. 拓扑排序 2.BFS/DFS
// 拓扑排序的话，可以带环图去除正常边后入度不为 0 的角度来计算。
// 这里就新构造一个图，然后使用 DFS 来查询。
// type Node struct {
// 	Val int
// 	Neighbors []*Node
// }

func canFinish(numCourses int, prerequisites [][]int) bool {
	nodes := make(map[int]*Node, numCourses)
	for i := 0; i < numCourses; i++ {
		nodes[i] = new(Node)
		nodes[i].Val = i
	}
	// build tree
	for _, pre := range prerequisites {
		nodes[pre[0]].Neighbors = append(nodes[pre[0]].Neighbors, nodes[pre[1]])
	}

	visited := make(map[int]int, numCourses)
	// 这里会出现一种情况，就是我们的图并不是一个连通图，会有不与任何其它节点连接的节点存在，所以起始节点不能设为 0 节点，而是应该每个节点都遍历一次
	// 而且对于有向图的遍历来说，确实需要对每个节点都进行遍历
	for i := 0; i < numCourses; i++ {

		if !canFinishDFS(nodes[i], &visited) {
			return false
		}
	}
	return true
}

// 这里有个问题，为什么使用了三种状态，而不是两种状态呢？DFS 的话，不应该使用两种状态就可以了吗？
// 因为一次dfs对应一个连通分量（这里指的是有向图的连通分量），我们只需要使用1表示在这次dfs里面已经访问过，如果在这次dfs里面再次访问，就表示有环，
// dfs之后再把1变为0即可。这样做虽然是正确的，但是会超时。所以我们引入第三种状态，例如-1，表示这个节点已经在之前的dfs里面访问过而且这个节点的连通分量不会存在环，
// 所以我们可以直接返回true。这样做可以避免重复访问。这才是引入三种状态的根本原因。
func canFinishDFS(node *Node, m *map[int]int) bool {
	if (*m)[node.Val] == -1 {
		return false
	}
	if (*m)[node.Val] == 1 {
		return true
	}

	(*m)[node.Val] = -1
	neighbors := node.Neighbors
	for _, neighbor := range neighbors {
		if !canFinishDFS(neighbor, m) {
			return false
		}
	}

	(*m)[node.Val] = 1
	return true
}

// 使用队列的 BFS 来对有向图进行环检测
func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	size := len(prerequisites)
	if size == 0 {
		return true
	}
	graph := make([][]int, numCourses)
	in := make([]int, numCourses)
	for _, edge := range prerequisites {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		in[edge[0]] += 1
	}

	var queue []int
	for point, degree := range in {
		if degree == 0 {
			queue = append(queue, point)
		}
	}

	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]
		points := graph[front]

		for _, point := range points {
			in[point] -= 1
			// 为什么是在这里进行检测，而不是在下面注释掉的部分?
			// 因为如果是在下面进行检测的话，我们会将哪些原本已经入队的入度为0的节点再次入队，这显然是不对的，相当于进入了死循环，会超时
			if in[point] == 0 {
				queue = append(queue, point)
			}
		}

		// for point, degree := range in {
		//     if degree == 0 {
		//         queue = append(queue, point)
		//     }
		// }
	}

	for _, degree := range in {
		if degree != 0 {
			return false
		}
	}

	return true

}
