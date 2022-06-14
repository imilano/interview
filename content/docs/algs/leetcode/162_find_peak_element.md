---
title: 0162. Find Peak Element
weight: 10
tags: [
	"Array",
	"Search",
	"Binary Search",
    "Monotone Stack",
    "Stack"
]
---

## Description
> A peak element is an element that is strictly greater than its neighbors.
> 
> Given an integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.
> 
> You may imagine that nums[-1] = nums[n] = -∞.
> 
> You must write an algorithm that runs in {{< katex >}} Omicron(\log n) {{< /katex >}} time.
## Solutions

### One Pass Iteration
这里可以通过一次数组的一次遍历来完成。为了方便，我们可以在头部和尾部分别添加上一个最小值，这样能够减少一些边界值判断。当然，这种方法并不符合题目要求的 {{< katex >}} \Omicron(n\log n) {{< /katex >}} 的要求。
```go
func findPeakElement(nums []int) int {
    // 可以在前后都添加一个最小值，这样可以减少一些判断
    nums = append([]int{math.MinInt}, append(nums, math.MinInt)...)
    size, start, res := len(nums), 1, 0
    for start + 1 <= size {
        if nums[start] > nums[start-1] && nums[start] < nums[start+1] {
            res = start
        }
        
        start++
    }
    
    
    return res
}

```

### Monotone Stack
这里还可以使用单调栈来解决。维护一个单调递增栈，如果当前遍历到的元素比栈顶元素要小，那么返回栈顶元素；如果栈为空或者当前元素比栈顶元素要大，那么将当前元素入栈，然后继续遍历。遍历过程中需要注意整个数组为递增数组或者数组中只有一个元素的情况，这种情况可以通过在遍历的时候将下标遍历到 size 来解决。
```go
func findPeakElement(nums []int) int {
    size := len(nums)
    if size == 1 {
        return 0
    }
    // 因为测试用例中包含了开头就是最小值的测试用例，所以下面这个策略其实并不会奏效
    // 可以在前后都添加一个最小值，这样可以减少一些判断
    // nums = append([]int{math.MinInt64}, append(nums, math.MinInt64)...)    
    
    // idx 之所以要遍历到 size ，是因为如果整个数组都是一个递增的数组的话，那么这个此时应该返回的是最后的一个数组下标
    // 也就是说，这里其实是一个 corner case，需要进行特殊处理
    var monoStack []int
    for idx := 0; idx <= size; idx++ {
        // 处理整个数组都为递增数组的情况，或者整个数组只有一个元素的情况（其实这个情况也包含在前一个情况中）
        if idx == size {
            return size - 1
        }
        
        // 如果栈为空或者当前元素比栈顶元素要大，那么说明当前是递增的，将当前元素入栈
        stackSize := len(monoStack)
        if stackSize == 0 || nums[idx] > nums[monoStack[stackSize - 1]] {
            monoStack = append(monoStack, idx)
            continue
        }
        
        // 如果栈顶元素比当前元素要大，那么说明找到了一个局部峰值，此时返回这个局部峰值即可
        if nums[idx] < nums[monoStack[stackSize-1]] {
            return monoStack[stackSize-1]
        }
        
    }
    
    // 如果找不到局部峰值，那么返回-1
    return -1
}
```

### Binary Search
