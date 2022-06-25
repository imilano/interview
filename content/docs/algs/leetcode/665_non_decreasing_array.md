---
title: 0665. Non-decreasing Array
weight: 10
tags: [
	"Array",
	"Greedy"
]
---

## Description
> Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most one element.
> 
> We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).

## Solution
### Greedy
题主这题做了好久，多次提交都不正确，只好上网上看各路大神们的解法。这里用的是贪心思想，这里你会发现，对于一个序列 abcd，如果 c 比 b 要小，那么你要么更改 c 的值，让当前元素等于前一个元素，要么更改 b 的值，让前一个元素（相对于 c 而言）小于等于当前元素，那改谁的值会让你受益更多呢？其实是改 b 的值，也就是 nums[i-1] 的值，因为减少值没有风险，但是增加值却可能会带来风险。但是，如果你发现 a > c 的时候，那么你就需要更改 c 的值，也就是 nums[i]，否则的话，你就需要同时更改 a 和 b 的值，而这显然是不被允许的。
```go
func checkPossibility(nums []int) bool {
    size := len(nums)
    if size <= 1 {
        return true
    }
    
    cnt, i := 1, 1
    for i < size {
        if nums[i] < nums[i-1] {
            cnt--
            if cnt < 0 {
                return false
            }
            // 优先更改前一个元素
            if i == 1 || nums[i-2] <= nums[i] {
                nums[i-1] = nums[i]
            } else {
                // 否则更改当前元素
                nums[i] = nums[i-1]
            }
        }
        i++
    }
    
    return true
}
```
