---
title: 0226. Invert Binary Tree
weight: 10
tags: [
	"Binary Tree",
	"Recursive"
]
---

## Description
> Given the root of a binary tree, invert the tree, and return its root.
## Solutions
### Recursive
太简单了，递归交换每个节点的左右子节点即可。
```go
func invertTree(root *TreeNode) *TreeNode {
    if root == nil || root.Left == nil && root.Right == nil {
        return root
    }
    
    left, right := invertTree(root.Left), invertTree(root.Right)
    root.Left, root.Right = right, left
    return root
}
```

### Iterative
或者也可以用中序遍历的的迭代方式。
```go
func invertTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}
```
