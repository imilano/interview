---
title: 0462. Minimum Moves to Equal Array Elements
weight: 10
tags: [
	"Array",
	"Math"
]
---

## Description
> Given an integer array nums of size n, return the minimum number of moves required to make all array elements equal.
> 
> In one move, you can increment or decrement an element of the array by 1.
> 
> Test cases are designed so that the answer will fit in a 32-bit integer.

## Solutions
### Math
这个题题主一开始以为是需要找到离平均数最近的那个数，然后用哪个数作为 target， 求这个数和其它所有数的差值的绝对值，对这些差值累加即可。结果提交之后，只能通过部分 OJ。
```go
func minMoves2(nums []int) int {
    size := len(nums)
    if size <= 1 {
        return 0
    }
    
    var sum int
    for _, num := range nums {
        sum += num
    }
    
    sort.Ints(nums)
    avg, index := float32(sum) / float32(size), 0
    for idx, _ := range nums {
        if absFloat(float32(nums[idx]) - avg) < absFloat(float32(nums[index]) - avg) {
            index = idx
        }
    }
    
    var res int
    for _, num := range nums {
        res += absInt(nums[index] - num)
    }
    
    return res
}


func absFloat(a float32) float32 {
    if a < 0 {
        return -a
    }
    
    return a
}

func absInt(a int) int {
    if a < 0 {
        return -a
    }
    
    return a
}
func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}
```

题主万万没想到，这里取的锚点是中位数。为什么中位数可以，而平均数就不可以呢？其实在对数组进行排序之后，就可以通过观察发现，要让所有元素相等，那么就需要增大左边的元素，减小右边的元素，直到所有元素相等。也就是说，所有元素需要中间相等，因此使用中位数比平均数更合适。这里可以用两个例子来说明，比如对于 [0,0,0,0,5] 以及 [0,0,0,4] 这两个数组，这二者取平均数作为锚点的话，所需要的 move 远大于取中值作为锚点来得大。也就是说，取均值肯定是不行的。

```go
func minMoves2(nums []int) int {
    sort.Ints(nums)
    res, left, right := 0, 0, len(nums) - 1
    for left < right {
        res += nums[right] - nums[left]
        left++
        right--
    }
    
    return res
}
```
