---
title: 0108. Convert Sorted Array to Binary Search Tree
weight: 10
tags: [
    "Binary Search Tree",
    "Tree"
]
---

## Description

> Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
>
> A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

## Solutions
简单题，直接递归构建即可。
```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


 func sortedArrayToBST(nums []int) *TreeNode {
    size := len(nums)
    if size == 0 {
        return nil
    }
    
    root := new(TreeNode)
    root.Val = nums[size/2]
    root.Left, root.Right = sortedArrayToBST(nums[:size/2]), sortedArrayToBST(nums[size/2+1:])
    return root
}
```
