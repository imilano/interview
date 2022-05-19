---
title: 0103. Binary Tree ZigZag Level Order Traversal
weight: 10
---

## Description
> Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).

## Solutions
简单题，使用队列进行层次遍历即可。
```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    var res [][]int
    if root == nil {
        return res
    }
    
    var queue []*TreeNode
    queue = append(queue, root)
    var flag bool
    for len(queue) != 0 {
        size := len(queue)
        var nums []int
        for i := 0; i < size; i++ {
            node := queue[i]
            nums = append(nums, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        queue = queue[size:]
        if flag {
            reverse(nums)
            flag = false
        } else {
            flag = true
        }
        
        res = append(res, nums)
    }
    
    return res
}

func reverse(nums []int) {
    size := len(nums)
    start, end := 0, size -1
    for start < end {
        nums[start], nums[end]  = nums[end], nums[start]
        start++
        end--
    }
}
```
