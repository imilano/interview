---
title: 23. Merge K Sorted List
weight: 10
---

## Description

> You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
> 
> Merge all the linked-lists into one sorted linked-list and return it.


## Solutions

### Min Heap

很简单的想法，使用最小堆即可。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKListsSolution1(lists)
}

// 使用最小堆，将所有的节点数据都插入到最小堆中，然后再把所有元素从最小堆中弹出
func mergeKListsSolution1(lists []*ListNode) *ListNode {
	var minHeap MinHeapArr
	heap.Init(&minHeap)

	for _, list := range lists {
		t := list
		for t != nil {
			heap.Push(&minHeap, t.Val)
			t = t.Next
		}

	}

	res := new(ListNode)
	node := res
	for len(minHeap) > 0 {
		val := heap.Pop(&minHeap).(int)
		t := new(ListNode)
		t.Val = val
		node.Next = t
		node = node.Next
	}

	return res.Next
}


type MinHeapArr []int

func (h MinHeapArr) Len() int {
	return len(h)
}

func (h MinHeapArr) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeapArr) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeapArr) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeapArr) Pop() interface{} {
	size := len(*h)
	res := (*h)[size-1]
	*h = (*h)[:size-1]
	return res
}

func (h *MinHeapArr) Top() interface{} {
	// TODO
	return nil
}
```
