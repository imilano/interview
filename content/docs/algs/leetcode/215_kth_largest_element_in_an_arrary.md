---
title: 0215. Kth Largest Element in an Array
weight: 10
tags: [
	"sort",
	"heap sort",
	"quick sort"
]
---

## Description

> Given an integer array nums and an integer k, return the kth largest element in the array.
> 
> Note that it is the kth largest element in the sorted order, not the kth distinct element.


## Solutions

### Heap 
创建一个 k 个大小的最小堆，首先先将前 k 个元素压入堆中，而后面的元素只有当其比堆顶元素要大的时候才可以入堆。最后堆顶元素即为所求。
```go
import (
    "container/heap"
)


func findKthLargest(nums []int, k int) int {
    return findKthLargestUsingSort(nums, 0, len(nums)-1, k)
}

// 解法1， 使用最大堆来解
func findKthLargestUsingHeap(nums []int, k int) int {
    var maxHeap MaxHeap
    heap.Init(&maxHeap)
    
    for _, num := range nums {
        heap.Push(&maxHeap, num)
    }
    
    for i := 0; i < k-1; i++ {
        heap.Pop(&maxHeap)
    }
    
    return heap.Pop(&maxHeap).(int)
}

//--------------- max heap ----------------------
type MaxHeap []int

func (h MaxHeap) Len() int {
    return len(h)
}

func (h MaxHeap) Swap(i,j int) {
    h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Less(i,j int) bool {
    return h[i] > h[j]
}

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
    size := h.Len()
    if size == 0 {
        return -1
    }
    
    x := (*h)[size-1]
    *h = (*h)[:size-1]
    return x
}
```

### Quick Sort
可以使用快排的 partition 算法。快排的 partition 算法中，选中选择一个基点作为 pivot，然后将小于 pivot 的元素放在 pivot 左边，大于 pivot 的元素放在 pivot 右边，然后返回 pivot 的下标，这样 pivot 左边的元素都比 pivot 小，右边的都比 pivot 大，这样的话就确定了一个元素的位置。那么我们只需要当 pivot = k-1 时候返回即可。
```go
func findKthLargest(nums []int, k int) int {
    return findKthLargestUsingSort(nums, 0, len(nums)-1, k)
}

func findKthLargestUsingSort(nums []int, left, right, k int) int {
    for {
        pivot := partition(nums, left, right)
        if pivot == k - 1 {
            return nums[k-1]
        }
        if pivot >  k-1 {
            // 返回的 pivot 比 k-1 要大，说明查找太偏右了，需要往左查找
            right = pivot - 1
        } else {
            // 返回的 pivot 比 k-1 要小，说明查找太偏左了，需要往右查找
            left = pivot + 1
        }
    }
}

func partition(nums []int, left int, right int) int {
    pivot := nums[left]
    l,r := left + 1, right
    for l <= r {
        if (nums[l] < pivot && nums[r] > pivot) {
            nums[l], nums[r] = nums[r], nums[l]
            l++
            r--
        }
        
        if nums[l] >= pivot {
            l++
        }
        
        if nums[r] <= pivot {
            r--
        }
    }
    
    
    nums[left], nums[r] = nums[r], nums[left]
    return r
}
```
