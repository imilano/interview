---
title: 0053. Maximum Subarray
weight: 10
---

## Description

> Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
> 
> A subarray is a contiguous part of an array.


## Solutions

### Kadane 

Kadane 算法，最优子结构性质，当前位置的和取决于前一个位置的子数组和以及当前元素。
```python

def max_subarray(arr):
	max_sum_so_far, sum_ending_here = arr[0], arr[0]
	for e in range(arr[1:]):
		sum_ending_here = max(e, sum_ending_here + e)
		max_sum_so_far = max(max_sum_so_far, sum_ending_here)

	return max_sum_so_far
```

如果数组中有负数，并且允许返回长度为 0 的子数列，则该问题可以改为：
```python
def max_subarray(arr):
	max_sum_so_far, sum_ending_here = 0, 0
	for e in range(arr):
		sum_ending_here = max(sum_ending_here + e, e)
		max_sum_so_far = max(max_sum_so_far, sum_ending_here)
return max_sum_so_far
```

完整代码是：
```go
func maxSubArray(nums []int) int {
    size := len(nums)
	if size == 0 {
		return 0
	}

	res, curSum := nums[0], nums[0]
	for i := 1; i < size; i++ {
		curSum = max(curSum+nums[i], nums[i])
		res = max(res, curSum)
	}

	return res
}

func max(a,b int) int {
    if a > b {
        return a
    }
    
    return b
}
```

### DP
定义 dp[i] 表示以 i 为结尾的连续子数组的最大和，则很明显，dp[i] 的值受到 dp[i-1] 和 array[i] 的影响，稍微一想即可得到转态转移方程：{{< katex >}}$dp[i] = \max (dp[i-1] + array[i], array[i])${{< /katex >}}
```go
func maxSubArray(nums []int) int {
    // write code here
    size := len(nums)
    if size == 0  {
        return 0
    }
    
    res := nums[0]
    dp := make([]int, size)
    dp[0] = array[0]
    for i := 1; i < size; i++ {
        dp[i] = max(dp[i-1]+nums[i], nums[i])
        res = max(res, dp[i])
    }
    
    return res
}

func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}
```
