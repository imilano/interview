---
title: 0454. 4 Sum II
weight: 10
tags: [
	"Divide",
	"Hash Table",
	"Array"
]
---

## Description
> Given four integer arrays nums1, nums2, nums3, and nums4 all of length n, return the number of tuples (i, j, k, l) such that:
> 
> - `0 <= i, j, k, l < n`
> - `nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0`

## Solutions
### Hash Table && Divide
首先题主肯定是写了一个四重循环，提交上去然后不出意外的超时了（笑。

然后经过一波搜索，发现可以这么做：这里使用的也算是分治思想，与其直接写一个四重循环，那么我们可以写两个二重循环，这样就可以有效降低时间复杂度。首先使用一个 map 记录 A 数组和 B 数组中每个元素组合之和出现的次数，然后再遍历 C 和 D，求出 C 和 D 当前元素组合的和，然后判断其相反数有没有在 map 中出现，如果出现了，则在结果上加上其出现的次数即可。不得不说，这个解法还是很巧妙的。
```go
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    var res int
    dict := make(map[int]int)
    for _, num1 := range nums1 {
        for _, num2 := range nums2 {
            dict[num1+num2]++
        }
    }
    
    for _, num1 := range nums3 {
        for _, num2 := range nums4 {
            if cnt,ok := dict[-num1-num2]; ok {
                res += cnt
            }
        }
    }
    
    return res
}
```
