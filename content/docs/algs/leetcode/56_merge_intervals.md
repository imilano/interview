---
title: 56. Merge Intervals
weight: 10
---
## Description

> Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.


## Solutions

典型的区间合并问题。

首先先对区间按照开始时间从小到大进行排序，然后使用一个指针 cur 指向当前 intervals 数组遍历到的位置，使用一个指针 tail 指向 res 数组最后一个元素。然后开始对 intervals 数组进行遍历，如果当前遍历到的区间跟 res 数组的最后一个区间没有重叠，则直接将这个区间插入 res 数组，然后 tail 自增；否则的话就合并区间，更新 res[tail] 的结束时间；每次遍历结束，cur 自增。
```go
func merge(intervals [][]int) [][]int {
    var res [][]int
	size := len(intervals)
    if size <= 1 {
        return intervals
    }
    sort.Slice(intervals, func(i,j int)bool {
        return intervals[i][0] < intervals[j][0]
    })

    res = append(res, intervals[0])
    var tail int
    cur := 1
    for cur < size {
        // 如果开始时间相等，或者 cur 的开始时间小于 tail 的结束时间，则说明有重叠，则需要合并
        if intervals[cur][0] <= res[tail][1] || intervals[cur][0] == res[tail][0]{
            res[tail][1] = max(res[tail][1], intervals[cur][1])
        } else {
            // 如果没有重叠，则直接插入
            res = append(res, intervals[cur])
            tail++
        }
        
        
        cur++
    }
    
    
    return res
}
//*************** util ****************
func min(a,b int) int {
    if a < b {
        return a
    }
    
    
    return b
}


func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}

```
