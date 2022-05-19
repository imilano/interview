---
title: 0078. Subsets
weight: 10
---
## Descrition
> Given an integer array nums of unique elements, return all possible subsets (the power set).
>
> The solution set must not contain duplicate subsets. Return the solution in any order.

## Solutions
观察一下可以发现，其实子集合就是一个追加操作。你只需要把新元素不断的追加到就有的集合上构成一个新的集合，然后再将这个新集合添加到结果数组中即可。重复上述操作，直到你把原数组中的所有元素都遍历完一次即可。
```go
func subsets(nums []int) [][]int {
    res := [][]int{{}}
    size := len(nums)
    
    for i := 0; i < size; i++ {
        n := len(res)
        for j := 0; j < n; j++ {
            t := make([]int, len(res[j]))
            copy(t, res[j])
            t = append(t, nums[i])
            res = append(res, t)
            
            // 注意不能用下面这样的方式。
            // 因为 golang 中切片其实底层是一个指向实际数组的指针，所以下面的添加可能会旧元素上进行添加，而不是在新元素上进行添加。
            // t := res[j]
            // t = append(t, nums[i])
            // res = append(res, t)
        } 
    }
    
    return res
}
```
