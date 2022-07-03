---
title: 0314. Binary Tree Vertical Order Traversal
weight: 10
tags: [
	"Binary Tree",
	"Hash Table",
	"BFS"
]
---

## Description
> Given a binary tree, return the vertical order traversal of its nodes' values. (ie, from top to bottom, column by column).
> 
> If two nodes are in the same row and column, the order should be from left to right.


## Solutions
### Level Order Traversal
这里可以隐约看出来是要使用层序遍历，但是与一般的层序遍历的稍微有点不同。这里的主要思想是：对每个节点赋予一个值 col，然后对于其左节点，col 减 1，对于右节点， col 加 1。这样的话，同一列的节点都会有同一个值 col。在入队的时候，不仅需要将节点入队，还需要将节点的 row 入队，二者可以组成一个 pair。

```go
func vericalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	type pair struct {
		Node *TreeNode
		Col int
	}

	// 层序遍历需要用到的队列
	var queue []pair
	// 记录每个 col 上有哪些节点
	dict := make(map[int]*TreeNode)
	queue = append(queue, pair{root, 0})
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			p := queue[0]
			queue = queue[1:]
			dict[p.Col] = append(dict[p.Col], p.Node)
			if p.Node.Left != nil {
				queue = append(queue, pair{p.Node.Left, p.Val - 1})
			}

			if p.Node.Right != nil {
				queue = append(queue, pair{p.Node.Right, p.Val + 1})
			}
		}

		queue = queue[size:]
	} 


	// golang 的 map 不会根据 key 进行排序，所以需要手动进行排序
	var keys []int
	for key, _ := range dict {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	for _, key := range keys {
		nodes := dict[key]
		var level []int
		for _, node := range nodes {
			level = append(level, node.Val)
		}

		res = append(res, level)
	}

	return res
}
```
