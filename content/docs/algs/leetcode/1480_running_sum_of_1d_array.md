---
title: 1480. Running Sum of 1d Array
weight: 10
tags: [
	"Array",
]
---

## Description
> Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
> 
> Return the running sum of nums.

## Solutions
简单题，直接看代码就好了。
```go
func runningSum(nums []int) []int {
    var res []int
    var pre int
    for _, num := range nums {
        pre += num
        res = append(res, pre)
    }
    
    return res
}
```
