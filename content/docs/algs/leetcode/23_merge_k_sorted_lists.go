package leetcode

import (
	"container/heap"
)

/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.
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

// 方法2，多路归并，可以使用败者树
