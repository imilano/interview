---
title: 0572. Subtree of Another Tree
weight: 10
tags: [
	"Tree",
	"Recursive"
]
---

## Description
> Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.
> 
> A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.


## Solutions
### Recursive
这里其实是简单题 sameTree 的一个应用。很明显可以看到，我们需要递归遍历 root 中的每个节点，然后看以这个节点为根节点的树是否和 subRoot 为根节点的树是 SameTree 关系。
```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    // 为什么这里单独针对 root 来进行判断呢？
    // 因为 isSubtree 其实是针对 root 节点进行的递归，所以这里理应将 root 节点的状态作为返回的判断条件，而不是 subRoot。
    if root == nil {
        return false
    }
    
    if sameTree(root, subRoot) {
        return true
    }
    
    
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}


func sameTree(root, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil {
        return true
    }
    
    if root == nil || subRoot == nil {
        return false
    }
    
    if root.Val != subRoot.Val {
        return false
    }
    
    
    return sameTree(root.Left, subRoot.Left) && sameTree(root.Right, subRoot.Right)
}
```
### Iterative
这里也有一种迭代的解法，那就是使用中序遍历的方式，将两棵树的节点值都存到数组中，然后查看 subRoot 的数组是否是 root 数组的子数组即可。
