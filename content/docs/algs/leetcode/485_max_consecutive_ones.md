---
title: 0485. Max Consecutive Ones
weight: 10
tags: [
	"Two Pointer",
	"Sliding Window"
]
---

## Description
> Given a binary array nums, return the maximum number of consecutive 1's in the array.

## Solutions
### Two Pointers
这个就是一个简单题啦，使用双指针来做就好了,直接看代码就好。
```go
func findMaxConsecutiveOnes(nums []int) int {
    res, size, left, right := 0, len(nums), 0, 0
    for right < size {
        // right 一直往右走，直到找到第一个非 1 的数才停下来
        for right < size && nums[right] == 1 {
            right++
            res = max(res, right - left)
        }
        
        // 跳过所有非 1 的数，同时更新 left 指向第一个 1
        for right < size && nums[right] != 1 {
            right++
            left = right
        }
    }
    
    return res
}

func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}

```
