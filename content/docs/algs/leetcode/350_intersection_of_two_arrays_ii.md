---
title: 0350. Intersection of Two Arrays II
weight: 10
tags: [
	"Hash Table",
	"Array"
]
---
## Description
> Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

## Solutions
这里跟 349 题的区别是，一个元素可能会出现多次，所以我们在统计完两个书中的元素的出现次数之后，需要确定把该元素放几次到结果数组中，那么需要放几次呢？这个次数应该跟该元素在两个数组中出现的最小次数相等。
```go
func intersect(nums1 []int, nums2 []int) []int {
    d1,d2 := make(map[int]int), make(map[int]int)
    for _, num := range nums1 {
        d1[num]++
    }
    
    for _, num := range nums2 {
        d2[num]++
    }
    
    var res []int
    for key, value := range d1 {
        if cnt, ok := d2[key]; ok {
            m := min(cnt, value)
            for i := 0; i < m; i++ {
                res = append(res, key)
            }
        }
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
