---
title: 0540. Single Element in a Sorted Array
weight: 10
tags: [
	"Array",
	"Search",
	"Binary Search"
]
---

## Description
> You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once.
> 
> Return the single element that appears only once.
> 
> Your solution must run in {{< katex >}}\Omicron(\log n) {{ < /katex >}} time and {{ < katex > \Omicron(1) {{ < /katex > }}} space.

### Bit Manipulation
这个题一眼看到，首先想到的就是使用异或操作，因为异或会导致相同的元素相互抵消为零，所以最后只会剩下哪个只出现一次的元素。
```go
func singleNonDuplicate(nums []int) int {
    var res int
    for _, num := range nums {
        res ^= num
    }
    
    return res
}
```


### Binary Search
这里因为数组有序，并且题目也要求 {{< katex >}} \Omicron(\log n) {{< /katex>}} 的解法，所以可以想到要使用二分来解，但是这里怎么二分呢？这里用到了一个小 trick，我们将坐标两两归为一对，比如 0 和 1 归为一对，2 和 3 归为一对等，那么对应一个下标 i， 我们将 i 异或 1 就可以找到它对应的 peer 的下标。比如对于 0， 异或 1 就是 1，而对于 1，异或 1 就是 0。此时如果该 peer 跟当前数字不相等，那么说明左侧的顺序出现了紊乱，那么就需要向左侧收缩；如果相等的话，说明左侧顺序没问题，那么就需要向右收缩。
```go
func singleNonDuplicate(nums []int) int {
    size := len(nums)
    left, right := 0, size -1
    for left < right {
        mid := left + (right - left)/2
        // 异或 1 可以找到它对应的 peer 下标，如果二者相等，说明左侧有序，则需要向右收缩；
        // 如果二者不等，说明左侧无序，需要向左收缩。
        if nums[mid] == nums[mid^1] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return nums[left]
}
```
