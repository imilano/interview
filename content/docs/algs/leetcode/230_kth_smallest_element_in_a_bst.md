---
title: 0230. Kth Smallest Element in a BST
weight: 10
tags: [
	"Binary Search Tree",
	"DFS",
	"Inorder Traversal"
]
---

## Description
> Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

## Solutions
这个题很简单，只需要使用中序遍历即可。一下分别给出递归和迭代的解法。
### Recursive
```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
    if root == nil {
        return math.MinInt
    }
    
    return helper(root, &k)
}

func helper(root *TreeNode, cur *int) int {
    if root == nil {
        return math.MinInt
    }
    
    val := helper(root.Left, cur)
    if val != math.MinInt {
        return val
    }
    
    *cur--
    if 0 == *cur {
        return root.Val    
    }
    
    return helper(root.Right, cur)
}
```
### Iterative
```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
    var queue []*TreeNode
    cur := root
    var cnt int
    for len(queue) != 0 || cur != nil {
        for cur != nil {
            queue  = append(queue, cur)
            cur = cur.Left
        }
        
        size := len(queue)
        // 注意这里不要覆盖了 cur，要使用 = 而不是 :=
        cur = queue[size-1]
        queue = queue[:size-1]
        cnt++
        if cnt == k {
            return cur.Val
        }
        
        
        cur = cur.Right
    }
    
    return -1
}
```
