package leetcode

import (
	"container/heap"
)

/*
The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value and the median is the mean of the two middle values.

For example, for arr = [2,3,4], the median is 3.
For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
Implement the MedianFinder class:

MedianFinder() initializes the MedianFinder object.
void addNum(int num) adds the integer num from the data stream to the data structure.
double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.
*/

// 可以使用最大最小堆来做。维护两个堆，左边是最大堆，右边是最小堆，每当新来一个元素，优先放到左边小根堆，然后将小根堆堆顶元素放入大根堆，如果小根堆元素数量小于大根堆元素数量，
// 那么将大根堆元素出堆放入小根堆，维持两边堆的元素数量平衡。
type MedianFinder struct {
	minHeap MinHeapArr
	maxHeap MaxHeap
}

// comment this to avoid conflicts
// func Constructor() MedianFinder {
// 	var minHeap MinHeapArr
// 	var maxHeap MaxHeap
// 	heap.Init(&minHeap)
// 	heap.Init(&maxHeap)

// 	return MedianFinder{
// 		minHeap: minHeap,
// 		maxHeap: maxHeap,
// 	}
// }

func (this *MedianFinder) AddNum(num int) {
	// 关键还是在这里的调整，这里的调整保证了左边大根堆的元素数量大于等于右边小根堆的数量：当元素总数为偶数时，二者相等；为奇数时，左边比右边大 1
	heap.Push(&(*this).minHeap, num)
	heap.Push(&(*this).maxHeap, heap.Pop(&this.minHeap))
	if (*this).minHeap.Len() < (*this).maxHeap.Len() {
		heap.Push(&this.minHeap, heap.Pop(&this.maxHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() > this.maxHeap.Len() {
		x := heap.Pop(&this.minHeap).(int)
		heap.Push(&this.minHeap, x)
		return float64(x)
	} else {
		x := heap.Pop(&this.minHeap).(int)
		heap.Push(&this.minHeap, x)

		y := heap.Pop(&this.maxHeap).(int)
		heap.Push(&this.maxHeap, y)

		return float64(x+y) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
