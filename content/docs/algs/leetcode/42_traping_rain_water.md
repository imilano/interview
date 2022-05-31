---
title: 0042. Traping Rain Water
weight: 10
tags: [
    "Array"
]
---

## Description

> Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
> 
> detail see: https://leetcode.com/problems/trapping-rain-water/


## Solutions

### Array
使用两个数组扫描两趟，第一个数组 fromLeft 从左到右扫描，记录到目前为止遇到的最大值，第二个数组fromRight 从右向左扫描，记录当前为止遇到的最大值。最后扫描一下这两个数组，对于每个位置 i，取 fromLeft 和 fromRight 二者中的较小者与当前 height[i] 的差值作为当前格子所能盛水量。
```go
func trap(nums []int) int {
    var res int
    size := len(nums)
    if size <= 1 {
        return res
    }
    
    curMax := nums[0]
    fromLeft := make([]int, size)
    for idx, num := range nums {
        curMax = max(num, curMax)
        fromLeft[idx] = curMax
    }
    
    curMax = nums[size-1]
    fromRight := make([]int, size)
    for i := size-1; i >= 0; i-- {
        curMax = max(nums[i], curMax)
        fromRight[i] = curMax
    }
    
    for i := 0; i < size; i++ {
        res += min(fromLeft[i], fromRight[i]) - nums[i]
    }
    
    return res
}

func  min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}

func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}
```


还有一种解法，就是先找出整个数组的最大值，这个最大值将整个数组切分为两部分，然后分别处理左半边和右半边。处理单边的时候，记录自己当前遇到的历史最大值，如果当前值比历史最大值要大，则更新历史最大值为当前值；否则的话说明历史最大值比当前值要大，那么当前值是可以盛水的，则更新结果盛水量。为什么只需要记录当前历史最大值就可以了呢，毕竟能不能盛水其实要靠两边而不是一边啊？答案是我我们已经找出了整个数组的最大值，那么当前的历史最大值只会比整个数组的最大值要小，能盛多少水完全取决于较短的一方（参考短板理论），所以只需要维持一个最大值即可。
```go
// 先找到数组中的最大值，这个最大值将数组分为左右两部分，然后分别计算左边的所能盛的水和右边所能盛的水
func trap(nums []int) int {
    var res int
    size := len(nums)
    if size <= 1 {
        return res
    }
    
    // 找到当前最大高度，这个高度将数组一分为二
    var maxHeight int
    for idx, value := range nums {
        if nums[maxHeight] < value {
            maxHeight = idx
        }
    }
    
    // 处理左半边
    peakL := nums[0] // 存储当前扫描到的历史
    left := 0
    for left < maxHeight {
        // 当前元素更大，则更新历史最大值
        if nums[left] > peakL {
            peakL = nums[left]
        } else {
            // 历史最大值还是大于当前元素，则可以盛水
            res += peakL - nums[left]
        }
        
        left++
    }

    // 处理右半边
    peakR := nums[size-1]
    right := size-1
    for right > maxHeight {
        if nums[right] > peakR {
            peakR = nums[right]
        } else {
            res += peakR - nums[right]
        }
        
        right--
    }
    

    return res
}
```


