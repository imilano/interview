---
title: 0027. Remove Element
weight: 10
tags: [
    "sort",
	"two poiner"
]
---

## Description
> Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The relative order of the elements may be changed.
> 
> Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.
> 
> Return k after placing the final result in the first k slots of nums.
> 
> Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

## Solutions

参考比较排序挪位置的想法，使用双指针进行进行排序和挪位置即可。
```go
func removeElement(nums []int, val int) int {
    size := len(nums)
    start, end := 0, size -1
    for start <= end {
        if nums[start] != val {
            start++
            continue
        } 
        
        for i := start+1; i <= end; i++ {
            nums[i-1] = nums[i]
        }

        end--
    }
    
    return end+1
}
```
