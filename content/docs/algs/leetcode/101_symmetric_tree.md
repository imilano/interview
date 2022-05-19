---
title: 0101. Symmetric Tree
weight: 10
---

## Description

> Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

## Solutions
简单题，左子树的左节点的值要等于右子树的右子树的右节点的值，左子树的右节点的值要等于右子树的左节点的值。

```go
func isSymmetric(root *TreeNode) bool {
    return helper(root, root)
}

func helper(root1, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }
    
    if root1 == nil || root2 == nil {
        return false
    }
    
    if root1.Val != root2.Val {
        return false
    }
    
    return helper(root1.Left, root2.Right) && helper(root1.Right, root2.Left)
}
```
