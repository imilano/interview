---
title: 0487. Max Consecutive Ones II
weight: 10
tags: [
	"Two Pointer",
	"Sliding Window"
]
---

## Description
> Given a binary array, find the maximum number of consecutive 1s in this array if you can flip at most one 0.
> 
> Example 1:
> 
> Input: [1,0,1,1,0]
> Output: 4
> Explanation: Flip the first zero will get the the maximum number of consecutive 1s.
>     After flipping, the maximum number of consecutive 1s is 4.
>  
> 
> Note:
> 
> The input array will only contain 0 and 1.
> The length of input array is a positive integer and will not exceed 10,000
>  
> 
> Follow up:
> What if the input numbers come in one by one as an infinite stream? In other words, you can't store all numbers coming from the stream as it's too large to hold in memory. Could you solve it efficiently?

## Solutions
### Sliding Window
这个题就是 1004 题 Max Consecutive Ones III 的具体化，可以直接使用该题的解法来解答。

可以维护一个窗口 [left,right] 来容纳至少k个0。当遇到了0，就累加 zero 的个数，然后判断如果此时0的个数大于k，则右移左边界left，如果移除掉的 nums[left] 为0，那么 zero 自减1。如果不大于k，则用窗口中数字的个数来更新 res.
```go
func longestOnes(nums []int, k int) int {
    res, size, left, right, cnt := 0, len(nums), 0, 0, 0
    for right < size {
        // 对 0 计数
        if nums[right] == 0 {
            cnt++
        }
        
        // 如果 cnt 大于 k 了，那么缩小 left 窗口，如果此时nums[left] 为 0， 则更新 cnt 值
        // 注意这里需要让 left 等于 right，因为下面计算 res 的时候加了 1
        for left <= right && cnt > k {
            if nums[left] == 0 {
                cnt--
            }
            
            left++
        }
        
        // 更新 res 值
        res = max(res, right - left + 1)
        right++
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
