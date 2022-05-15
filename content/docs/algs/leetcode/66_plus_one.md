---
title: 66. Plus One
weight: 10
---
## Description

> You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.
> 
> Increment the large integer by one and return the resulting array of digits.

## Solutions

简单题，直接计算即可。
```go
func plusOne(digits []int) []int {
    size := len(digits)
    if size == 0 {
        return digits
    }
    
    
    carry := 1
    for i := size-1; i >= 0; i-- {
        sum := digits[i] + carry
        carry = sum / 10
        digits[i] = sum % 10
    }
    
    if carry > 0 {
        digits = append([]int{carry}, digits...)
    }
    
    return digits
}
```

