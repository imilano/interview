package leetcode

import (
	"container/heap"
)

/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
*/

// 使用最大堆.
// 先对每个数字的出现次数排序，然后使用最大堆，其中 key 为出现次数，value 为数字，基于 key 排序即可
func topKFrequent(nums []int, k int) []int {
	var res []int
	m := make(map[int]int)
	for _, value := range nums {
		m[value] += 1
	}

	var maxHeap CustomizedMaxHeap
	heap.Init(&maxHeap)
	for num, count := range m {
		heap.Push(&maxHeap, []int{count, num})
	}

	for i := 0; i < k; i++ {
		ele := heap.Pop(&maxHeap).([]int)
		res = append(res, ele[1])

	}
	return res
}

// 自定义堆
type CustomizedMaxHeap [][]int

func (h CustomizedMaxHeap) Len() int {
	return len(h)
}

func (h CustomizedMaxHeap) Less(i, j int) bool {
	return h[i][0] > h[j][0]
}

func (h CustomizedMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CustomizedMaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *CustomizedMaxHeap) Pop() interface{} {
	size := len(*h)
	res := (*h)[size-1]
	*h = (*h)[:size-1]
	return res
}
