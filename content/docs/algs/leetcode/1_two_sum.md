---
title: 0001. Two Sum
weight: 10
tags: [
	"Sort",
	"Hash Table",
	"Two Pointer"
]
---

## Description
> Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
> 
> You may assume that each input would have exactly one solution, and you may not use the same element twice.
> 
> You can return the answer in any order.

## Solutions

### Hash Table
这里因为答案唯一，那么也就意味着数组里的元素是唯一的，所以可以使用 map 来做，只需要一次遍历即可，具体看下面代码就好。
```go
func twoSum(nums []int, target int) []int {
    dict := make(map[int]int)
    for idx, num := range nums{
        t := target - num
        if exist(dict, t) {
            return []int{idx, dict[t]}
        }
        
        dict[num] = idx
    }
    
    
    return nil
    
}

func exist(dict map[int]int, target int) bool {
    if _, ok := dict[target]; ok {
        return true
    }
    
    return false
}
```

### Sort && Two Pointer
这里也可以使用排序来做，根据值进行排序，但是需要保存该值对应的下标。排好序之后可以使用双指针，一个 left 指向开头，一个 right 指向结尾，计算二者和，如果小于 target，则 left++，否则 right++，直到二者之和等于 target 即可。思路大概就这样，代码就不写了。
