---
title: 0098. Validate Bianry Search Tree
weight: 10
---

## Description
> Given the root of a binary tree, determine if it is a valid binary search tree (BST).
> 
> A valid BST is defined as follows:
> 
> The left subtree of a node contains only nodes with keys less than the node's key.
> The right subtree of a node contains only nodes with keys greater than the node's key.
> Both the left and right subtrees must also be binary search trees.

## Solutions

### Inorder Traversal
首先可以使用中序遍历的的方式遍历一下，然后记录每个遍历到的节点值到一个数组里，然后判断这个数组是否有序即可。这里还可以优化一下，没必要存储整个数组，只需要使用一个指针 pre，记录中序遍历的前一个值即可，然后使用这个值与当前值进行比较，如果当前值小于等于 pre，那么说明这个序列不是递增的，则返回 false。全部遍历结束之后，返回 true。
```go
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    var stack []*TreeNode
    cur := root
    pre := math.MinInt
    for cur != nil || len(stack) != 0 {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        
        size := len(stack)
        cur = stack[size-1]
        stack = stack[:size-1]
        if cur.Val <= pre {
            return false
        }
        
        pre = cur.Val
        cur = cur.Right
    }
    
    return true
}
```

有迭代当然有递归，下面是递归的方式：
```go
 func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    return helper(root, math.MinInt, math.MaxInt)
    
}

func helper(root *TreeNode, mn,mx int) bool {
    if root == nil {
        return mn < mx
    }
    
    if root.Val < mn || root.Val > mx {
        return false
    }
    
    return helper(root.Left, mn, root.Val) && helper(root.Right, root.Val, mx)
}
```

