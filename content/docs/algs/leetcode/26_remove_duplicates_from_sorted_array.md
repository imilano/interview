---
title: 0026. remove duplicated from sorted array
weight: 10
tags: [
    "Array"
]
---

## Description

> Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.
> 
> Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
> 
> Return k after placing the final result in the first k slots of nums.
> 
> Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

## Solution

简单题。使用一个 pos 指针指向当前不重复元素应该存储的位置，然后遍历数组，对于重复元素则跳过，知道遇到不重复元素位置，将该元素放到 pos 处，然后让 pos 和 start 都自增，然后继续遍历
```go
func removeDuplicates(nums []int) int {
    size := len(nums)
    if size <= 1 {
        return size
    }
    
    start, pos := 1, 1
    for start < size {
        for start < size && nums[start] == nums[start-1] {
            start++
        }
        if start < size {
            nums[pos] = nums[start]
            pos++
        }

        start++
    }
    
    return pos
    
}
```
