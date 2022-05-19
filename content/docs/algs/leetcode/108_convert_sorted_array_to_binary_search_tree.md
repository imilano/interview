---
title: 0108. Convert Sorted Array to Binary Search Tree
weight: 10
---

## Description

> Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
>
> A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

## Solutions
简单题，直接递归构建即可。
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    size := len(nums)
    if size == 0 {
        return nil
    }
    
    return helper(nums, 0, size-1)

}


func helper(nums []int, start, end int) *TreeNode {
    if start > end {
        return nil
    }
    
    mid := (start+end)/2
    node := new(TreeNode)
    node.Val = nums[mid]
    node.Left, node.Right = helper(nums, start, mid-1), helper(nums, mid+1, end)
    return node
}
```
