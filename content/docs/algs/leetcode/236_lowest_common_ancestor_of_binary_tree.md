---
title: 0236. Lowest Common Ancestor of Binary Tree
weight: 10
tags: [
	"Binary Tree",
	"Recursive"
]
---
## Description
> Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
> 
> According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

## Solutions
### Recursive
这里如果是一棵二叉排序树，那么问题就会简单很多，但是这里只是普通的二叉树，所以稍微就会稍微难一点。注释都写在代码中了，判断的过程值得好好体会一下。
```go

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// lowestCommonAncestor 判断 root 节点
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 首先空子树肯定不包含我们所要的
     if root == nil {
         return nil
     }

	// 递归遍历左右子节点
     left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	 // 如果当前根节点刚好等于其中一个节点的值，那么返回根节点
    if root.Val == p.Val || root.Val == q.Val {
         return root
     }
	// 否则，如果左右节点中有一个非空，说明已经在左/右子树中找到了 lcs，那么返回该非空节点
     if left == nil {
         return right
     }
     
     if right == nil {
         return left
     }
     
	 // 此时左右节点都非空，说明当前根节点刚好是 lcs，则直接返回 root
    return root
 }
 ```
