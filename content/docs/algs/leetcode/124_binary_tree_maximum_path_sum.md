---
title: 0124. Binary Tree Maximum Path Sum
weight: 10
tags: [
    "Binary Tree",
    "Recursive"
]
---

## Description
> A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.
> 
> The path sum of a path is the sum of the node's values in the path.
> 
> Given the root of a binary tree, return the maximum path sum of any non-empty path.

## Solutions
比较明显的是，对于每个节点来说，我们需要知道经过其左子节点的 path 之和大还是经过其右子节点的 path 之和大。
```go
func maxPathSum(root *TreeNode) int {
    res := math.MinInt
    maxPathSumRecur(root, &res)
    return res
}

func maxPathSumRecur(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    
    // 这里之所以在 max 中放一个 0，其实是做了一些剪枝操作。如果把这个与 0 的剪枝操作去掉，那么下面的代码还是比较好理解的。
    // 如果把 0 去掉，那么下面就是分别求得当前节点左右子树的最大路径和，然后比较 res 与 左右节点和当前节点构成的路径之和的大小并
    // 更新 res，最后返回左右路径中的较大者与当前节点值的和。
    // 与 0 进行比较的话去除了路径和为负数的情况，做了一个不错的剪枝机制。
    // 左子树的最大路径和，如果为负数，则返回 0
    left := max(maxPathSumRecur(root.Left, res), 0)
    // 右子树的最大路径和，如果为负数，则返回 0
    right := max(maxPathSumRecur(root.Right, res), 0)
    // 更新当前最大路径和
    *res = max(*res, left + right + root.Val)

    // 选择当前能让路径和最大的分支
    return max(left, right) + root.Val
}


func max(a, b int) int {
    if a < b {
        return b
    }
    
    return a
}
```
