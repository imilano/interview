---
title: 0973. K Closest Points to Origin
weight: 10
tags: [
	"Sort",
	"Heap"
]
---

## Description
> Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).
> 
> The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).
> 
> You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).


## Solutions
### Sort
这道题说到底就是一个排序问题，就是比较各个点到原点的距离，那么我们只需要根据这个距离对各个点进行排序即可。
```go
func kClosest(points [][]int, k int) [][]int {
    var res [][]int
    sort.Slice(points, func(i,j int) bool {
        return points[i][0]*points[i][0] + points[i][1]*points[i][1] < points[j][0]*points[j][0]+points[j][1]*points[j][1]
    })
    
    for i := 0; i < k; i++ {
        res = append(res, points[i])
    }
    return res
}

```

### Heap
这里也可以使用最大堆来做。创建一个 k 个大小的最大堆，然后把距离和 points 下标组成的 Pair 对插入堆中，根据距离进行排序，当堆中数量超过 k 时，将堆顶元素出堆。最后剩下的 k 个元素就是 k 个最小的元素。

大致思想就是这样，代码就不写了。
