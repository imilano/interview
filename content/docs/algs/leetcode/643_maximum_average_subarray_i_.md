---
title: 0643. Maximum Average Subarray I
weight: 10
tags: [
	"Sliding Window",
	"Array"
]
---

## Description
> You are given an integer array nums consisting of n elements, and an integer k.
> 
> Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.


## Solutions
### Sliding Window 
这题没什么难度，可以用一个固定长度的队列来模拟滑动窗口即可。
```go
func findMaxAverage(nums []int, k int) float64 {
    var queue []int
    res, size, sum := math.MinInt, len(nums), 0
    for i := 0; i < size; i++ {
        queue = append(queue, nums[i])
        sum += nums[i]
		// 注意第一个 if 语句和第二个 if 语句的先后关系
        if len(queue) > k {
            sum -= queue[0]
            queue = queue[1:]
        }
        if len(queue) == k && sum > res {
            res = sum
        }
    }
    
    return float64(res)/float64(k)
}
```
