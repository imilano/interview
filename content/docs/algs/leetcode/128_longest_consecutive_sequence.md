---
title: 0128. Longest Consecurive Sequence
weight: 10
tags: [
	"Hash Table"
]
---

## Description
> Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
> 
> You must write an algorithm that runs in O(n) time.


## Solutions
使用一个集合存入所有的数字，然后遍历数组中的每个数字，如果其在集合中存在，那么将其移除，然后 pre 再自减 1，直至pre 不在集合之中，对 next 采用同样的方法，那么 next - pre -1 就是当前数字的最长连续序列。之所以要移除数字，是为了避免重复计算。比如说对于 4、3、2，如果计算 4 的时候不把 3 和 2 移除掉，那么计算 3 和 2 的时候就会出现重复计算的情况。
```go
func longestConsecutive(nums []int) int {
    var res int
    dict := make(map[int]bool)
    for _, num := range nums {
        dict[num] = true
    }
    
    for _, num := range nums {
        left, right := num -1, num + 1
        for dict[left] == true {
            delete(dict, left)
            left--
        }
        
        for dict[right] == true {
            delete(dict, right)
            right++
        }

        
        res = max(res, right-left-1)
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
