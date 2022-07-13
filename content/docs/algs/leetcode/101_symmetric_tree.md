---
title: 0101. Symmetric Tree
weight: 10
tags: [
    "Binary Tree",
    "Recursive",
    "Inorder Traversal"
]
---

## Description

> Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

## Solutions
### Recursive
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

### Iterative
这里也可以使用中序遍历的方式，对这个树进行一次中序遍历，然后将所有节点值都放在数组中，然后检查这个数组的元素是否左右对称。这里也比较简单，就不写代码了。
