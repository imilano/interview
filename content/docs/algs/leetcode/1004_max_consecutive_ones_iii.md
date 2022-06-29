---
title: 1004. Max Consecutive Ones III
weight: 10
tags: [
	"Sliding Window",
	"Two Pointer"
]
---

## Description
> Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
## Solutions
### Sliding Window
这题一看就是要用滑动窗口来解啦，但是窗口的边界一定要控制好。题主初始写出了下面的解法，结果只通过了部分case，调来调去总是有一部分 case 不能照顾到。
```go
func longestOnes(nums []int, k int) int {
    res, size, left, right,cnt := 0, len(nums), 0, 0, 0
    for right < size {
		// 如果 0 的数量比 k 小，那么继续扩大右边界
        for right < size && cnt <= k {
            if nums[right] == 0 {
                cnt++
            }
			// 更新窗口值
            res = max(res, right - left)
            right++
            
        }
        
		// 缩小左边界
        for left < right && cnt > k {
            if nums[left] == 0 {
                cnt--
            }
            left++
        }
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
于是题主最后还是去看了网上大神的解答，发现其实是自己搞复杂了，这个题完全没有必要写得像上面那么复杂。可以维护一个窗口 [left,right] 来容纳至少k个0。当遇到了0，就累加 zero 的个数，然后判断如果此时0的个数大于k，则右移左边界left，如果移除掉的 nums[left] 为0，那么 zero 自减1。如果不大于k，则用窗口中数字的个数来更新 res。
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
