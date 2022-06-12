---
title: 1658. Minimum Operations to Reduce X to Zero
weight: 10
tags: [
	"Backtracing",
	"Prefix Sum",
	"Sliding Window"
]
---

## Description
> You are given an integer array nums and an integer x. In one operation, you can either remove the leftmost or the rightmost element from the array nums and subtract its value from x. Note that this modifies the array for future operations.
> 
> Return the minimum number of operations to reduce x to exactly 0 if it is possible, otherwise, return -1.

## Solutions

### Backtrace
这是题主一开始想出来的解法，直接暴力回溯，但是超时了。仔细分析一下，发现时间复杂度确很高得离谱...
```go
func minOperations(nums []int, x int) int {
    res := math.MaxInt
    helper(nums, x, 0, &res)
    if res == math.MaxInt {
        return -1
    }
    return res
}

func helper(nums []int, x int, cur int, res *int) {
    size := len(nums)
    if x < 0 {
        return
    }
    
    if x == 0 {
        *res = min(*res, cur)
        return
    }
    
    if size == 0 {
        return
    }
    
    helper(nums[1:], x - nums[0], cur+1, res)
    helper(nums[:size-1], x - nums[size-1], cur+1, res)
    
} 

func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```

### Prefix Sum
楼主在思考上面的回溯法优化中，其实还是有看到这题的一点前缀和的特征的，但是在想不出来具体的解法....最后还是看了网上的解法...

其实我们可以转换一下思路，不要总想着「什么时候选择左边，什么时候选择右边」。其实这个问题可以转换一下就是，在原数组中找到一个子数组，让这个子数组的和等于整个数组的和减去题目要求的 x，如果存在这样的最长子数组，那么久返回总数组的长度减去子数组的长度，否则，返回-1。而因为这里要求的最小操作次数，那么意味着我们只需要找最长的子数组即可。这么一转换的话，是不是马上又回到了我们熟悉的前缀和了呢？

这是题主首先写出的前缀和解法，结果也是超时了...
```go
func minOperations(nums []int, x int) int {
    size := len(nums)
    prefix := make([]int, size+1)
    for idx, num := range nums {
        prefix[idx+1] = prefix[idx] + num
    }

    // corner case
    if prefix[size] < x {
        return -1
    }
    
    res := math.MaxInt
    for i := 1; i <= size; i++ {
        for j := 0; j <= i; j++ {
            if prefix[j] + prefix[size] - prefix[i] == x {
                res = min(res, size - i + j)
            }
        }
    }
    
    if res == math.MaxInt {
        return -1
    }
    
    return res
}


func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```

稍微优化一下，提交之后勉强通过：
```go
func minOperations(nums []int, x int) int {
    size := len(nums)
    prefix := make([]int, size+1)
    for idx, num := range nums {
        prefix[idx+1] = prefix[idx] + num
    }

    // corner case
    if prefix[size] < x {
        return -1
    }
    if prefix[size] == x {
        return size
    }
    
    // 连续子数组的和
    sumOfSubArr := prefix[size] - x
    // 使用一个 map， key 为前缀和，value 为达成该前缀和的下标，用来记录指定前缀和可以在哪个下标求得
    dict := make(map[int]int)
    res := math.MinInt
    for i, preSum := range prefix {
        // 如果存在一个子数组[j,i]使得 prefix[i] - prefix[j] = sumOfSubArr， 那么也就只需要比较 res 和 i-j 的大小即可。
        // 这里我们使用一个 map 来表示 prefix[j]，这样的话只O(n) 的遍历。
        s := preSum - sumOfSubArr 
        if idx, ok := dict[s]; ok {
            res = max(res, i - idx)
        }
        
        if _, ok := dict[preSum]; !ok {
            dict[preSum] = i
        }
    }
    
    if res == math.MinInt {
        return -1
    }
    
    return size - res
}

func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}
```

### Sliding Window
这里也可以使用滑动窗口来解。思路还是跟前面相似，都是要找到最大子数组，其和等于整个数组的和减去 x，只不过使用了滑动窗口的方式来解决。方法是维持一个滑动窗口，每当窗口中的值大于 subArrSum 时，收缩左边界；如果相等，则说明找到了一个连续子数组，使其和等于了整个数组和减去 x，那么更新 res。每次遍历都需要不断移动右边界。
```go
func minOperations(nums []int, x int) int {
    size := len(nums)
    var sum int 
    for _, num := range nums {
        sum += num
    }
    
    // corner case
    if sum == x {
        return size
    }
    if sum < x {
        return -1
    }
    
    // 维护一个滑动窗口，如果滑动窗口中的值之和大于连续子数组和，则左边界右移；
    // 如果滑动窗口中的值之和等于最大连续子数组和，则说明找到了一个子数组，其和加上 x 等于整个数组和，那么更新 res
    // 每次遍历，都需要不断移动右边界
    res := math.MaxInt
    left, right, windowSum, subArrSum := 0, 0, 0, sum - x
    for right < size {
        windowSum += nums[right]
        for windowSum > subArrSum {
            windowSum -= nums[left]
            left++
        }
        
        if windowSum == subArrSum {
            res = min(res, size - (right-left+1))
        }
        right++
    }
    
    if res == math.MaxInt {
        return -1
    }
    
    return res
}

func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```


