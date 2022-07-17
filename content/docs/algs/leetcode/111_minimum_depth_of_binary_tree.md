---
title: 0111. Minimum Depth of Binary Tree
weight: 10
tags: [
	"Binary Tree",
	"Iterative",
	"Recursive"
]
---

## Description

> Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
>
> A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

## Solutions
### Iterative
只需要简单的层次遍历，终止条件就是当队列中出现一个左右子节点均为空节点的节点即可。
```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    var queue []*TreeNode = []*TreeNode{root}
    var depth int
    for len(queue) != 0 {
        depth++
        size := len(queue)
        for i := 0; i < size;i++ {
            if queue[i].Left == nil && queue[i].Right == nil {
                return depth
            }
            
            if queue[i].Left != nil {
                queue  = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        
        queue = queue[size:]
    }
    
    return 0
}
```
### Recursive
递归解法需要主体返回条件，首先当根节点为空时，需要返回 0，然后还需要加上两个条件”如果根节点的左节点为空，则返回右节点的遍历值；如果根节点的右节点为空，则需要返回左节点的遍历值“，如果不加上上面两个条件，那么对于 [1, null,2] 这样的树，结果会是 1 而不是 2，但 2 才是正确答案。
```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    if root.Left == nil && root.Right == nil {
        return 1
    }
    
    if root.Left == nil {
        return minDepth(root.Right) + 1
    }
    
    if root.Right == nil {
        return minDepth(root.Left) + 1
    }
    
    
    
    return min(minDepth(root.Left), minDepth(root.Right)) + 1
}


func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```
