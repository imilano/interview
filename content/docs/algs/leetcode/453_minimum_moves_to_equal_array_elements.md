---
title: 0453. Minimum Moves to Equal Array Elements
weight: 10
tags: [
	"Math",
	"Array"
]
---

## Description
> Given an integer array nums of size n, return the minimum number of moves required to make all array elements equal.
> 
> In one move, you can increment n - 1 elements of the array by 1.

## Solutions
### Math
这题的意思就是，对于一个长度为 n 的数组，每次可以给 n-1 个数组加 1，问最少要多少次，才能让整个数组的数字相等。那这里要给哪些数字加 1 呢？很明显可以想到，每次都需要给最小的 n-1 个数加 1，这也就意味着，每次 move 都需要排序来找出最小的 n - 1 个数。可惜的是，依照上面的思路，是无法通过 OJ 的。

这里的解法比较 tricky，其实给 n-1 个数字加 1，就等同于给最大的数字减去 1. 那么每次 move 对最小的 n - 1 个数字加 1，就相当于每次 move 对最大的那个数减 1，那么问题也就转化为，求每个数字跟最小数字之间的差值，然后对这写差值进行累加即可。
```go
func minMoves(nums []int) int {
    var idx,res int
    // find min
    for i, _ := range nums {
        if nums[i] < nums[idx] {
            idx = i
        }
    }
    
    // get sum
    for _, num := range nums {
        res += num - nums[idx]
    }
    
    
    return res
}
```
