---
title: 0268. Missing Number
weight: 10
tags: [
	"Math",
]
---
## Description
> Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.


## Solutions
这题很简单，只需要使用求和公式来进行计算就可以了。直接给出代码。
```go
func missingNumber(nums []int) int {
    n := len(nums)
    sum := (0+n) * (n+1)/2
    for _, num := range nums {
        sum -= num
    }
    
    return sum
}
```
