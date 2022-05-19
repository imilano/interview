---
title: 0075. Sort Colors
weight: 10
---

## Description
> Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
> 
> We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
> 
> You must solve this problem without using the library's sort function.


## Solutions

其实这里应该用任意的一个排序算法都可以。
### 冒泡排序的思路
参考冒泡排序算法。
```go
// 方法 1， 冒泡排序的思路
func sortColors(nums []int)  {
    size := len(nums)
    for i := size -1; i > 0; i-- {
        for j := 0; j < i; j++  {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
}
```

### 计数排序的思路

```go
func sortColors(nums []int)  {
    colors := make([]int, 3)
    for _, color := range nums {
        colors[color]++
    }
    
    var cur int
    for i := 0; i < 3; i++ {
        for j := 0; j < colors[i]; j++ {
            nums[cur] = i
            cur++
        }
    }
}
```
