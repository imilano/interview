package leetcode

import (
	"container/heap"
)

/*
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.
*/

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

func findKthLargestUsingSort(nums []int, left, right, k int) int {
	if left <= right {
		pivot := partition(nums, left, right)
		if pivot < k-1 {
			findKthLargestUsingSort(nums, pivot+1, right, k-1-pivot-1)
		} else if pivot > k-1 {
			findKthLargestUsingSort(nums, left, pivot-1, k)
		} else if pivot == k-1 {
			return nums[k-1]
		}
	}

	return -1
}

func partition(nums []int, left int, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[left] > nums[pivot] {
			left++
		}

		for left < right && nums[right] < nums[pivot] {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}

	nums[left], nums[pivot] = nums[pivot], nums[left]
	return left
}

//--------------- max heap ----------------------
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Less(i, j int) bool {
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
