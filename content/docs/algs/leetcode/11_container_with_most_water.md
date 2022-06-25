---
title: 0011. Container with Most Water
weight: 10
tags: [
	"Two Pointer",
	"Greedy"
]
---

## Description
> You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
> 
> Find two lines that together with the x-axis form a container, such that the container contains the most water.
> 
> Return the maximum amount of water a container can store.
> 
> Notice that you may not slant the container.

## Solutions
### Two Pointers
使用两个指针 i、j分别指向起点中终点，此时容器的盛水量 res=整个数组的长度*min（起点高度nums[i]，终点高度nums[j]）。接下来继续遍历，保留最大的 res 即可。
那么如何遍历呢 ，也就是所，i 和 j 如何移动呢。这里的想法是，最大盛水量取决于短板高度，如果要取得最大盛水量，那么就需要弥补短板，那么也就说，只需要高度以较低的一方移动即可。
当高度相等时，理论上移动任何一方都是可以的，我们可以固定移动左边，也可以固定移动右边。

```go
func maxArea(height []int) int {
    var res int
    size := len(height)
    left, right := 0, size -1
    for left <right {
        res = max(res, min(height[left], height[right]) * (right - left))
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    
    return res
}

func max(a,b int) int  {
    if a < b {
        return b
    }
    
    return a
}

func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```
