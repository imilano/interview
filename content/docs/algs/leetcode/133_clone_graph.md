---
title: 0133. Clone Graph
weight: 10
tags: [
	"Hash Table",
	"Graph",
	"BFS",
	"DFS"
]
---
## Description
> Given a reference of a node in a connected undirected graph.
> 
> Return a deep copy (clone) of the graph.
> 
> Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.
> ```go
> class Node {
>     public int val;
>     public List<Node> neighbors;
> }
>  ```
> 
> Test case format:
> 
> For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.
> 
> An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.
> 
> The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.

## Solutions
### DFS
图克隆或者树克隆，最先想到的就是要使用 Hash Table，这里用 BFS 或者 DFS 来做都可以。
```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    dict := make(map[*Node]*Node)
    return dfs(node, &dict)
}

func dfs(root *Node, dict *map[*Node]*Node) *Node {
    // 如果为空，直接返回
    if root == nil {
        return nil
    }
    
    // 如果这个节点已经遍历过了，则返回其副本节点
    if _, ok := (*dict)[root]; ok {        
        return (*dict)[root]
    }
    
    // 如果这个节点没遍历过，则根据当前节点创建一个副本节点
    // 同时复制当前节点和其它节点的关系，最后返回复制的节点
    dummy := new(Node)
    dummy.Val = root.Val
    (*dict)[root] = dummy
    for _, node := range root.Neighbors {
        // 这里需要对返回的节点进行判空处理，空节点不能直接放入副本的邻居中
        r := dfs(node, dict)
        if r != nil {
            (*dict)[root].Neighbors = append((*dict)[root].Neighbors , r)    
        }
    }
    
    // 返回副本节点
    return dummy
}
```
