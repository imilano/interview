---
title: 0104. Maximum Depth of Binary Tree
weight: 10
tags: [
	"Binary Tree",
	"Recursive",
	"Iterative"
]
---

## Description
> Given the root of a binary tree, return its maximum depth.
> 
> A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

## Solutions
### Recursive
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

### Iterative
这里只需要使用层次遍历，看最多可以遍历几层即可。
```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }   
    
    var queue []*TreeNode = []*TreeNode{root}
    var depth int
    for len(queue) != 0 {
        size := len(queue)
        depth++
        for i := 0; i < size; i++ {
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        
        queue = queue[size:]
    }
    
    return depth
}
```
