---
title: 0124. Binary Tree Maximum Path Sum
weight: 10
---

## Description
> A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.
> 
> The path sum of a path is the sum of the node's values in the path.
> 
> Given the root of a binary tree, return the maximum path sum of any non-empty path.

## Solutions

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
