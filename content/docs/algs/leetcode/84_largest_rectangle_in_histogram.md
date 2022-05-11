---
title: 84. Largest Rectangle in Histogram
weight: 10
---


## Description

> Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.

## Solutions

寻找每一个局部峰值 i (heights[i] > heights[i+1])，然后从这个局部峰值 i 开始向前遍历所有的值，算出共同的矩形面积，每次对比保留最大值。 很不幸，超时了 :)
```go
func largestRectangleArea(heights []int) int {
    var res int
    size := len(heights)
    if size == 0 {
        return res
    }
    
    for i := 0; i < size; i++ {
        if i+1 < size && heights[i] <= heights[i+1] {
            continue
        }
        
        curMin := heights[i]
        for j := i; j >= 0; j-- {
            if heights[j] < curMin {
                curMin = heights[j]
            }
            res = max(res, curMin * (i-j+1))
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


