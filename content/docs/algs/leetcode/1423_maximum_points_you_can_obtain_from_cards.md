---
title: 1423. Maximum Points You Can Obrain fron Cards
weight: 10
tags: [
	"Prefix Sum",
	"DP"
]
---
## Description
> There are several cards arranged in a row, and each card has an associated number of points. The points are given in the integer array cardPoints.
> 
> In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.
> 
> Your score is the sum of the points of the cards you have taken.
> 
> Given the integer array cardPoints and the integer k, return the maximum score you can obtain.

## Solutions
### Prefix Sum
leetcode 有一题的解法跟这一题几乎完全一样，但是可惜题主忘了到底是哪一题了。这个题要求每次从左边或者右边取一个数，取出这个数之后数组长度就会减少，要求取出 k 个数，并且让这个 k 个数的和最大。这里的关键就在于如何决策，取的时候到底是要从左边取还是要从右边取。这里最先想到的就是使用枚举的办法，不过这个办法是指数级别的，佷容易就会超时了。

其实这个题，可以转换为使用前缀和来做。我们不去管决策，而是在这个数组中，找长度为 size -k 的一个子数组，让这个子数组的和最小，那么剩下的几个数的和自然就是相对最大的，于是问题就转换为找一个长度为 size -k 的和最小的子数组，这就是前缀和问题啦。
```go
// 转换为找一个长度为 size - k 的和最小的子数组，于是便转换为一个前缀和问题
func maxScore(cardPoints []int, k int) int {
    size := len(cardPoints)
	// 计算前缀和
    prefix := make([]int, size + 1)
    for i := 1; i <= size; i++ {
        prefix[i] = prefix[i-1] + cardPoints[i-1]
    }
    
	// 注意 res 要设置为最大值，因为下面使用了 min 函数，如果初始化 res 为 0 的话，结果就总会是整个数组的和
    res, left := math.MaxInt, 0
    for left + size - k <= size {
		// 取长度为 size -k 的子数组的最小和
        res = min(res, prefix[left+size-k] - prefix[left])
        left++
    }
    
    
	// 用整个数组的和减去子数组最小和，即为剩余数字最大和
    return prefix[size] - res
}

func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```
