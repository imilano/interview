---
title: 0560. Subarray Sum Equals K
weight: 10
tags: [
	"Array",
	"Prefix Sum"
]
---

## Description
> Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
> 
> A subarray is a contiguous non-empty sequence of elements within an array.

## Solutions
### Prefix Sum
这个题可以使用前缀和的技巧来解决。前缀和相关的介绍可以看[这里](https://oi-wiki.org/basic/prefix-sum/)。
```go
func subarraySum(nums []int, k int) int {
    var res int
    size := len(nums)
    prefix := make([]int, size+1)

	// 计算前缀和
    idx := 1
    for _, num := range nums {
        prefix[idx] = prefix[idx-1] + num
        idx++
    }
    
	// 计算子数组
    for i := 1; i <= size; i++ {
        for j := 0; j < i; j++ {
            if prefix[i] - prefix[j] == k {
                res++
            }
        }
    }
    
    return res
}
```
