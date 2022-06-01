---
title: 0347. Top K Frequent Elements
weight: 10
tags: [
	"Hash Table",
	"Sort",
	"Heap",
	"Min Heap"
]
---
## Description
> Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

## Solutions
### Hash Table & Sort
这里首先使用 map 来统计每个元素出现的次数，然后将元素和其出现的次数组成一个 pair 对，根据每个 pair 对的元素出现次数对 pair 对进行排序，最后取前 k 个即可。
```go
func topKFrequent(nums []int, k int) []int {
    dict := make(map[int]int)
    for _, num := range nums {
        dict[num]++
    }
    
    type Pair struct {
        num int
        cnt int
    }
    
    var pairs []Pair
    for key, cnt := range dict {
        pairs = append(pairs, Pair{key, cnt})
    }
    
    sort.Slice(pairs, func(i,j int) bool {
        return pairs[i].cnt > pairs[j].cnt
    })
    
    var res []int
    for i := 0; i < k; i++ {
        res = append(res, pairs[i].num)
    }
    
    return res
}
```

### Hash Table & Min Heap
还是先对数组中元素的出现次数进行统计，然后将每个元素以及该元素出现的次数作为 pair 对插入最大堆中，最后只需要最大堆的前 k 个元素即可。
```go
func topKFrequent(nums []int, k int) []int {
	var res []int
    m := make(map[int]int)
	for _, value  := range nums {
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

func (h CustomizedMaxHeap) Less(i,j int) bool {
	return h[i][0] > h[j][0]
}

func (h CustomizedMaxHeap) Swap(i,j int) {
	h[i], h[j]= h[j], h[i]
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
```
