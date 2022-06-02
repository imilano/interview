---
title: 0318. Kth Smallest Element in a Sorted Matrix
weight: 10
tags: [
	"Heap",
	"Max Heap"
]
---

## Description
> Given an n x n matrix where each of the rows and columns is sorted in ascending order, return the kth smallest element in the matrix.
> 
> Note that it is the kth smallest element in the sorted order, not the kth distinct element.
> 
> You must find a solution with a memory complexity better than {{< katex >}} \Omicron(n^2) {{< /katex >}}.

## Solutions
这里要求空间复杂度要小于 {{< katex >}} \Omicron(n^2) {{< /katex >}}，那么其实可以构造一个大小为 k 的最大堆。不断将元素插入堆中，当堆大小大于 k 时，将堆顶元素出栈。等所有元素都遍历完之后，堆顶元素即为所求元素。
```go
import "container/heap"
func kthSmallest(matrix [][]int, k int) int {
    var maxHeap MaxHeap
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

type MaxHeap []int
func (m MaxHeap) Len() int {
    return len(m)
}

func (m MaxHeap) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

// 这里控制是构造最小堆还是最大堆，如果使用 < 的话，构造的就是最大堆
func (m MaxHeap) Less(i,j int) bool {
    return m[i] < m[j]
}

func (m *MaxHeap) Push(x interface{}) {
    *m = append(*m, x.(int))
}

// 这里永远应该弹出的是最后一个元素
func (m *MaxHeap) Pop() interface{} {
    x := (*m)[m.Len()-1]
    *m = (*m)[:m.Len()-1]
    return x
}
```
