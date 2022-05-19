---
title: 0094. Bianry Tree Inorder Traversal
weight: 10
---
## Description
> Given the root of a binary tree, return the inorder traversal of its nodes' values.

## Solutions
二叉树中序遍历，又分为迭代和递归两种方式。

### Recursive
```go
func inorderTraversal(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }
    
    helper(root, &res)
    return res
}

func helper(root *TreeNode, res *[]int) {
    if root == nil {
        return 
    }
    
    helper(root.Left, res)
    *res = append(*res, root.Val)
    helper(root.Right, res)
}
```

### Iterative

中序遍历需要借助栈来实现。
```go
func inorderTraversal(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    }
    
    var stack []*TreeNode
    cur := root 
    for cur != nil || len(stack) != 0 {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        
        size := len(stack)
        cur = stack[size-1]
        stack = stack[:size-1]
        res = append(res, cur.Val)
        cur = cur.Right
    }
    return res
}
```
