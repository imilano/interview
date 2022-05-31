---
title: 0041. First Missing Positive
weight: 10
tags: [
    "Hash Table",
    "Array"
]
---
## Description

> Given an unsorted integer array nums, return the smallest missing positive integer.
> 
> You must implement an algorithm that runs in O(n) time and uses constant extra space.

## Solutions

### Hash Table

可以使用一个 hash table 记录所以出现的值，然后找出当前值中的最大值。然后从 1 到最大值开始遍历，如果出现一个不在 hash table 中的数，则将该数返回即可。如果所有数都出现在了 hash table 中，那么返回最大值加 1 即可。但是最后这里需要注意的是，有可能整个数组都是负数，这样的话会导致前面的判断失效，所以返回值应该最小要返回 1.
```go
func firstMissingPositive(nums []int) int {
    dict := make(map[int]bool)
    curMax := nums[0]
    for _, num := range nums {
        dict[num] = true
        if num > curMax {
            curMax = num
        }
    }
    
    for i := 1; i <= curMax; i++ {
        if _, ok := dict[i]; !ok {
            return i
        }
    }
    
    return max(1, curMax + 1)
}

func max(a, b int) int {
    if a < b {
        return b
    }
    
    return a
}
```

进一步的，其实没必要遍历那么多数，只需要遍历 n 个数即可，n 为数组长度，上述改进版如下：
```go
func firstMissingPositive(nums []int) int {
    dict := make(map[int]bool)
    for _, num := range nums {
        dict[num] = true
    }
    
    n := len(nums)
    for i := 1; i <= n; i++ {
        if _, ok := dict[i]; !ok {
            return i
        }
    }
    
    return max(1, n+1)
}

func max(a, b int) int {
    if a < b {
        return b
    }
    
    return a
}
```


### Array

因为对空间复杂度有要求，所以上述的 hash table 的方法应该不能满足要求。这里的思路就是复用原数组，我们希望 nums[0]=1, nums[1] = 2, nums[i] = i+1。那么就需要遍历一次数组，遍历过程中如果发现 nums[i] != i+1, 也就是 nums[nums[i]-1] != nums[i]，那么我们就交换这两个元素。 最后扫描一遍数组，如果发现 nums[i] != i+1， 那么说明 i + 1 不在数组中，直接返回 i+1 即可。如果遍历完了发现从 1 到 n （为什么是 n，因为下标 0 存放的是 1，那么下标 n-1 存放的是 n）都在数组中，那么最后只需要返回 n+1 即可。
```go
func firstMissingPositive(nums []int) int {
    n := len(nums)
    for i := 0; i < n; i++ {
		// 注意这里是 while 循环，因为 nums[nums[i]-1] 和 nums[i] 交换的时候已经改变了 nums[i] 的值了
		// 所以这个 while 循环不能简单的用 if 来代替
        for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
            nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
        }
    }
    
    for i := 0; i < n; i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }
    
    return n+1
}
```
