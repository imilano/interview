package leetcode

/*
Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}


Test case format:

For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.

An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.

*/
// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	set := make(map[*Node]*Node)
	result := cloneGraphDFS(node, &set)
	return result
}

// 其实就是对图进行一次遍历
// 比较重要的是，为了避免重复添加，需要使用 map 来进行去重，来对应原图中的节点和新生成的克隆图的节点
func cloneGraphDFS(node *Node, set *map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	// 注意这里的去重操作
	if _, ok := (*set)[node]; ok {
		return (*set)[node]
	}

	res := new(Node)
	res.Val = node.Val
	// 这一句应该放在这里
	(*set)[node] = res
	neighbors := node.Neighbors
	for _, neighbor := range neighbors {
		res.Neighbors = append(res.Neighbors, cloneGraphDFS(neighbor, set))
	}

	// 这一句不应该放在这里，放在这里会导致内存爆炸
	// (*set)[node] = res
	return res
}
