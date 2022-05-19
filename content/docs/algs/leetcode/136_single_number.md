---
title: 0136. Single Number
weight: 10
---

## Description
> Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
> 
> You must implement a solution with a linear runtime complexity and use only constant extra space.

## Solutions

简单题，直接使用异或运算符即可。注意 res 初始化应该为 0.
```go
func singleNumber(nums []int) int {
    var res int
    for _, num := range nums {
        res ^= num
    }
    
    return res
}
```
