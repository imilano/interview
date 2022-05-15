---
title: 55. Jump Game
weight: 10
---

## Description
> You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
> 
> Return true if you can reach the last index, or false otherwise.

## Solutions
### Recursive
很容易想到递归的办法，但是也很容易就超时了 :)
```go
func canJump(nums []int) bool {
	size := len(nums)
    if helper(0, size, nums) {
        return true
    }
    
    return false

}

func helper(start int, size int, nums []int) bool {
    if start >= size -1 {
        return true
    }
    
    for i := 1; i <= nums[start]; i++ {
        if helper(start+i, size, nums) {
            return true
        }
    }
    
    return false
}
```
### Greedy
使用一个变量 remianed 表示当前剩余的跳力（指从当前节点最多还可以跳几步），然后使用一个值 maxDst 表示当前能够跳到的最远距离。每遍历到一个节点，判断是否能跳到当前节点，也就是 i 是否大于 Dst，如果大于，则说明跳不到当前节点，则可以直接返回 false； 否则说明能跳到当前节点，然后根据当前节点的跳力更新当前节点的最大跳力以及能跳到的最远距离。
```go
func canJump(nums []int) bool {
    size := len(nums)
    var remained, maxDst int
    for i := 0; i < size; i++ {      
		// 跳不到当前节点  
        if i > maxDst {
            return false
        }

		// 剩余的最大跳力
        remained = max(nums[i], remained)
		// 从当前节点跳是否能跳更远的距离
        maxDst = max(maxDst, i + remained)
        remained--
    }
    
    
    return true 
}


func max(a, b int) int {
    if a < b {
        return b
    }
    
    return a
} 

```
### DP
维护一个dp数组，其中dp[i]表示到达i位置所剩余的跳力，若到达某个位置时跳力为负，则说明无法到达该位置。

当前位置的跳力跟上一个位置的跳力以及上一个位置剩余的跳力有关，所以当前位置的剩余跳力和当前位置新的跳力中较大的那个数决定了当前能达到的最远距离，而下一个位置的剩余跳力等于当前位置的跳力减去1，所以 dp[i] = max(dp[i-1], nums[i-1]) - 1。如果当前dp为负，则说明无法抵达当前位置，返回false。

```go
func canJump(nums []int) bool {
	size := len(nums)
	dp := make([]int, size)
	for i := 1; i < size; i++ {
		dp[i] = max(dp[i-1], nums[i-1]) - 1
		if dp[i] < 0 {
			return false
		}
	}

	return true
}

func max(a, b int) int {
    if a < b {
        return b
    }
    
    return a
} 
```

