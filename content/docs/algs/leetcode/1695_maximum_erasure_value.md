---
title: 1695. Maximum Erasure Value
weight: 10
tags: [
	"Sliding Window"
]
---

## Description
> You are given an array of positive integers nums and want to erase a subarray containing unique elements. The score you get by erasing the subarray is equal to the sum of its elements.
> 
> Return the maximum score you can get by erasing exactly one subarray.
> 
> An array b is called to be a subarray of a if it forms a contiguous subsequence of a, that is, if it is equal to a[l],a[l+1],...,a[r] for some (l,r).

## Solutions
### Sliding Window & Hash Table
这题一眼可以看出来要使用滑动窗口来解。而且还是我们之前经常使用过的滑动窗口配合 Hash Table 的套路，这题的解法跟 `3. Longest Substring Without Repeating Character`几乎一模一样。
```go
func maximumUniqueSubarray(nums []int) int {
    size := len(nums)
    // calculate prefix sum
	// 前缀和用于方便计算子数组和
    prefix := make([]int, size+1)
    for idx, num := range nums {
        prefix[idx+1] = num + prefix[idx]
    }
    
	// 注意 left 的初值，因为 left 总是取窗口左边界的前一个，所以这里只初值取-1 即可。
    res,left,right := 0,-1,0
    dict := make(map[int]int)
    for right < size  {
        if idx, ok := dict[nums[right]]; ok && idx > left {
            left = idx
        }
        
        res = max(res, prefix[right+1] - prefix[left+1])
        dict[nums[right]] = right
        right++
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
