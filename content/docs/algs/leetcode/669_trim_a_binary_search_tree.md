---
title: 0669. Trim a Binary Search Tree
weight: 10
tags: [
	"Binary Search Tree",
	"Tree",
	"Recursive"
]
---

## Description
> Given the root of a binary search tree and the lowest and highest boundaries as low and high, trim the tree so that all its elements lies in [low, high]. Trimming the tree should not change the relative structure of the elements that will remain in the tree (i.e., any node's descendant should remain a descendant). It can be proven that there is a unique answer.
> 
> Return the root of the trimmed binary search tree. Note that the root may change depending on the given bounds.

## Solutions
### Recursive
因为根节点会发生改变，所以我们这里需要按照找最小公共父节点的方式，先找到 L 和 R 的最小公共父节点，然后再对该父节点来进行处理。其中该父节点的左节点应该是对以该左节点为根的子树进行 trim 后的子树，父节点的右节点应该是对以该节点右节点为根的子树进行 trim 后的子树。

对根节点的左右子树的处理可能不是很好理解，举几个例子印证一下就好了。
```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
    // 如果根节点为空，则直接返回即可
    if root == nil {
        return nil
    }
    
    // 因为根节点可能会发生改变，所以我们首先就是要找到 low 和 high 的最小公共祖先节点，这样 root 才是正确的。
    // 所以下面两个递归都是为了找到最小公共祖先
    if root.Val > high {
        return trimBST(root.Left, low, high)
    } 
    
    if root.Val < low {
        return trimBST(root.Right, low, high)
    }
    
    // 找到了 low 和 high 的最小公共祖先节点
    root.Left = trimBST(root.Left, low, high)
    root.Right = trimBST(root.Right, low, high)
    return root
}
```
