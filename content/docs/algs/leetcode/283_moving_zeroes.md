---
title: 0283. Moving Zeroes
weight: 10
tags: [
	"Two Pointer",
]
---
## Description
> Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
> 
> Note that you must do this in-place without making a copy of the array.

## Solutions
### Two Pointer
简单题，直接上代码：
```go
func moveZeroes(nums []int)  {
    size := len(nums)
    if size <= 1 {
        return
    }
    
    left, right := 0, 0
    for right < size {
        if nums[right] != 0 {
            nums[left] = nums[right]
            left++
        }
        
        right++
    }
    
    for left < size {
        nums[left] = 0
        left++
    }
}
```
