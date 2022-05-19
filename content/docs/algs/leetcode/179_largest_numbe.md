---
title: 0179. Largest Number
weight: 10
tags: [
	"sort",
	"string"
]
---

## Description

> Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.
> 
> Since the result may be very large, so you need to return a string instead of an integer.

## Solutions
这里的解法需要十分注意，跟一般的排序是有区别的。
```go
import (
    "strconv"
    "sort"
    "strings"
)
func largestNumber(nums []int) string {
    var rs []string
    for  _, num := range nums {
        rs = append(rs, strconv.Itoa(num))
    }
    
	// 这里应该进行连接排序
    sort.Slice(rs, func(i,j int) bool {
        return rs[i] + rs[j] > rs[j] + rs[i]
    })
    
    res := strings.Join(rs, "")
    size := len(res)
    
    // 如果全部都是 0，那么只需要保留最后一个 0 即可
    var pos int
    for pos < size-1 && res[pos] == '0' {
        pos++
    }
    
    return res[pos:]
}
```
