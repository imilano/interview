---
title: 0046. Permutations
weight: 10
tags: [
    "Backtracing"
]
---

## Description

> Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
>
## Solutions

### Array(swap)

基于交换进行 permutate 即可。
```go
func permute(nums []int) [][]int {
	var res [][]int
	size := len(nums)
	helper(0, size, nums, &res)
	return res
}

func helper(start,size int, nums []int, res *[][]int) {
    if start >= size {
        tmp := make([]int, size)
        copy(tmp, nums)
        *res = append(*res, tmp)
        return
    }
    
    
    for i := start; i < size; i++ {
        nums[i], nums[start] = nums[start], nums[i]
        // 注意这里是 start + 1， 而不是 i
        helper(start+1, size, nums, res)
        nums[i], nums[start] = nums[start], nums[i]
    }
}

```
