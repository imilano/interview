---
title: 0366. Find Leaves of Binary Tree
weight: 10
tags: [
	"Binary Tree",
	"BFS"
]
---

## Description
> Given a binary tree, collect a tree's nodes as if you were doing this: Collect and remove all leaves, repeat until the tree is empty.
> 
> Example:
> 
> Input: [1,2,3,4,5]
>   
>           1
>          / \
>         2   3
>        / \     
>       4   5    
> 
> Output: [[4,5,3],[2],[1]]

## Solutions
### BFS
这个题要求，每次都把叶节点给出调，然后再剩下的树中再把叶节点除掉，重复上述过程直到数中没有节点为止。仔细观察一下，你会发现，每次除掉的都是出度为 0 的节点。反过来说，如果你将这个树视作一个以叶节点为出发点的图，那么这个问题就会变成一个拓扑排序题。如果你能这么想，那问题也就很简单了。首先先将这个树视作一个以叶节点为出发点、以根节点为结束点的的有向图，然后统计每个节点的入度，每次把入度为 0 的节点放到数组中，然后将从该节点可达的节点的入度减去 1.循环上述过程即可，典型的拓扑排序问题。

同样是上面的思路，但是可以不把这个图视作有向图，可以视作无向图，然后对无向图的每个节点统计出入度，上面的入度为 0 的条件可以替换为入度为 1 。

Golang 本身没有 set， 确实很让人蛋疼。

```go
func findLeaves(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// 将树视为无向图，从根节点开始进行 BFS，对于每条边，统计两个节点的入度。
	// 然后将入度为 1 的节点入队(因为无向图入度为 1 表示该节点为叶节点)，开始进行无向图的拓扑排序。
	degree := make(map[int]map[int]bool)

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if _, ok := degree[node]; !ok {
				degree[node.Val] = make(map[int]bool)
			}
			if node.Left != nil {
				if _, ok := degree[node.Left.Val]; !ok {
					degree[node.Left.Val] = make(map[int]bool)
				}

				degree[node.Val][node.Left.Val] = true
				degree[node.Left.Val][node.Val] = true
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				if _, ok := degree[node.Right.Val]; !ok { 
					degree[node.Right.Val] = make(map[int]bool)
				}
				degree[node.Val][node.Right.Val] = true
				degree[node.Right.Val][node.Val] = true
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]
	}

	// 将入度为 1 的节点放入队列（这些入度为 1 的节点其实就是叶节点），再次开始 BFS
	var arr []int
	for key, peer := range degree {
		if len(peer) == 1 {
			arr = append(arr, key)
		}
	}

	// 再次开始 BFS， 每次在队列中的节点就是叶节点
	var res [][]int
	for len(arr) != 0 {
		var level []int
		size := len(arr)
		for i := 0; i < size; i++ {
			nodes := degree[arr[i]]
			level = append(level, arr[i])
			for node, _ := range nodes {
				delete(degree[node], queue[i])
				if len(degree[node]) == 1 {
					queue = append(queue, node)
				}
			}
		}


		res = append(res, level)
		arr = arr[size:]
	}

	return res
}
```
