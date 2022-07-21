---
title: 0700. Search In a Binary Search Tree
weight: 10
tags: [
	"Binary Search Tree",
	"Tree",
	"Recursive"
]
---
## Description
> You are given the root of a binary search tree (BST) and an integer val.
> 
> Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.

## Solutions
### Recursive
简单题，直接上代码。
```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    
    if root.Val < val {
        return searchBST(root.Right, val)
    }
    
    if root.Val > val {
        return searchBST(root.Left, val)
    }
    
    return root
}
```
### Iterative
```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    
    var queue []*TreeNode
    cur := root
    for cur != nil || len(queue) != 0 {
        for cur != nil {
            if cur.Val == val {
                return cur
            }
            queue = append(queue, cur)
            cur = cur.Left
        }
        
        size := len(queue)
        cur = queue[size-1]
        queue = queue[:size-1]
        cur = cur.Right
    }
    
    return nil
}
```
