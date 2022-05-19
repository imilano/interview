---
title: 0034. Find First and Last Position of Element in Sorted Array
weight: 10
---

## Description

> Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
> 
> If target is not found in the array, return [-1, -1].
> 
> You must write an algorithm with O(log n) runtime complexity.


## Solutions

### Binary Search

二分查找再加上中间扩散，很容易想出来。

```go
func searchRange(nums []int, target int) []int {
    res := []int{-1, -1}
    size := len(nums)
    if size == 0 {
        return res
    }
    
    left, right := 0, size -1
    for left <= right {
        mid := (left+right)/2
        if nums[mid] == target {
            l,r := mid, mid
            res[0] = l
            res[1] = r
            for l >= 0 && nums[l] == target {
                res[0] = l
                l--
            }
            
            for r < size && nums[r] == target {
                res[1] = r
                r++
            }
            
            break
        } else if nums[mid] > target {
            right = mid -1
        } else if nums[mid] < target {
            left = mid + 1
        }
    }
    
    return res
}
```
