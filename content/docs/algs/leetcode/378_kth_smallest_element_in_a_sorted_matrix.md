---
title: 0378. Kth Smallest Element in a Sorted Matrix
weight: 10
tags: [
	"Matrix",
	"Heap",
	"Min Heap"
]
---
## Description
> Given an n x n matrix where each of the rows and columns is sorted in ascending order, return the kth smallest element in the matrix.
> 
> Note that it is the kth smallest element in the sorted order, not the kth distinct element.
> 
> You must find a solution with a memory complexity better than {{< katex >}} \Omicron(n^2) {{< /katex >}}


## Solutions
这里可以使用最小堆，遍历矩阵，不断往堆中加入元素，遍历完之后，取堆中第 k 个元素就可。或者使用一个大小为 k 的最大堆，当堆中元素数量不足 k 时，将元素入堆，当堆中元素比 k 大时，将堆顶元素出堆即可。
```go
import "container/heap"
func kthSmallest(matrix [][]int, k int) int {
    var maxHeap MinHeap
    heap.Init(&maxHeap)
    m, n := len(matrix), len(matrix[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            heap.Push(&maxHeap, matrix[i][j])
        }
    }
    
    var res int
    for i := 0; i < k; i++ {
        res = heap.Pop(&maxHeap).(int)
    }
    
    return res
}

type MinHeap []int
func (m MinHeap) Len() int {
    return len(m)
}

func (m MinHeap) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

// 这里控制是构造最小堆还是最大堆，如果使用 < 的话，构造的就是最小堆
func (m MinHeap) Less(i,j int) bool {
    return m[i] < m[j]
}

func (m *MinHeap) Push(x interface{}) {
    *m = append(*m, x.(int))
}

// 这里永远应该弹出的是最后一个元素
func (m *MinHeap) Pop() interface{} {
    x := (*m)[m.Len()-1]
    *m = (*m)[:m.Len()-1]
    return x
}
```
