---
title: 1060. Missing Element in Sorted Array
weight: 10
tags: [
	"Search",
	"Binary Search"
]
---
## Description
> Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
> 
> For example,
> Given nums = [0, 1, 3] return 2.
> 
> Note:
> Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?

## Solutions
### Math
这个题还是很简单的，首先需要求出求出 [0,n] 的累和，然后用这个累和减去数组中出现的每个数，最后剩下的数就是结果。
```go
func missingNumber(nums []int) int {
	size := len(nums)
	res := (1 + n)*n/2
	for _, num := range nums {
		res -= num
	}

	return num
}
```

### Bit Manipulation
这个题也可以使用异或来做。既然从 0 到 n 中缺了一个，那么我们就可以使用从 0 到 n 的每个数字来跟数组中的数字进行异或操作，最后的那个数字肯定就是缺的那个数字啦。
```go
func missingNumber(n int) int {
	var res int
	for idx, num := range nums {
		res ^= (i+1)^nums[i]
	}

	return res
}
```
