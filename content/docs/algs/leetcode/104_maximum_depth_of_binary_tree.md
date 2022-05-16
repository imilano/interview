---
title: 104. Maximum Depth of Binary Tree
weight: 10
---

## Description
> Given the root of a binary tree, return its maximum depth.
> 
> A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

## Solutions
简单题，无需多说。
```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}


func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}
```
