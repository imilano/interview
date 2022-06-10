---
title: 0162. Find Peak Element
weight: 10
tags: [
	"Array",
	"Search",
	"Binary Search"
]
---

## Description
> A peak element is an element that is strictly greater than its neighbors.
> 
> Given an integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.
> 
> You may imagine that nums[-1] = nums[n] = -∞.
> 
> You must write an algorithm that runs in {{< katex >}} Omicron(\log n) {{< /katex >}} time.
## Solutions

### One Pass Iteration
这里可以通过一次数组的一次遍历来完成。为了方便，我们可以在头部和尾部分别添加上一个最小值，这样能够减少一些边界值判断。当然，这种方法并不符合题目要求的 {{< katex >}} \Omicron(n\log n) {{< /katex >}} 的要求。
```go
func findPeakElement(nums []int) int {
    // 可以在前后都添加一个最小值，这样可以减少一些判断
    nums = append([]int{math.MinInt}, append(nums, math.MinInt)...)
    size, start, res := len(nums), 1, 0
    for start + 1 <= size {
        if nums[start] > nums[start-1] && nums[start] < nums[start+1] {
            res = start
        }
        
        start++
    }
    
    
    return res
}

```

### Binary Search
