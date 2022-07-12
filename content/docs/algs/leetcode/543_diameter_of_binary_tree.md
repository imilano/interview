---
title: 0543. Diameter of Binary Tree
weight: 10
tags: [
	"BInary Tree",
	"Recursive"
]
---

## Description
> Given the root of a binary tree, return the length of the diameter of the tree.
> 
> The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
> 
> The length of a path between two nodes is represented by the number of edges between them.

## Solutions
### Recursive
这是题主一开始想出来的方法：用一个全局变量 res 代表直径最大值，然后分别当前节点计算左右子树的最大高度，然后比较左右子树再加上当前节点拼起来的树的最大直径与 res 的大小，根据比较结果来更新 res。写得稍微有点复杂。
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    var res int
    res = max(helper(root, &res), res)
    return res
}


func helper(node *TreeNode, res *int) int {
    // 如果是叶节点或者空节点，那么直径为 0
    if node == nil || node.Left == nil && node.Right == nil {
        return 0
    }
    
    // 分别找到左右节点的最大高度
    left, right := helper(node.Left, res), helper(node.Right, res)
    // 如果左右节点中有一个节点是空节点，那么 res 就只需要与其中的非空子节点的高度进行比较
    if node.Left == nil || node.Right == nil {
        *res = max(*res, max(left, right))
    } else {
        // 如果左右节点均非空，那么直径就是左子树的最大高度再加上右子树的最大高度，然后再加上 2。
        // 将上述相加的结果与 res 进行比较
        *res = max(*res, left+right+2)    
    }
    
    // 返回以当前节点为跟的子树的最大高度，其值为左/右子树的最大高度再加上 1
    return max(left, right) + 1
}

func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}
```
