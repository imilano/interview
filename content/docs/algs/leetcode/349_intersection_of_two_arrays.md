---
title: 0349. Intersection of Two Arrays
weight: 10
tags: [
	"Hash Table"
]
---
## Description
> Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

## Solutoins
简单题，只需要使用 map 即可。
```go
func intersection(nums1 []int, nums2 []int) []int {
    d1,d2 := make(map[int]int), make(map[int]int)
    for _, num := range nums1 {
        d1[num]++
    }
    
    for _, num := range nums2 {
        d2[num]++
    }
    
    var res []int
    for key, _ := range d1 {
        if _, ok := d2[key]; ok {
            res = append(res, key)
        }
    }
    
    return res
}
```
