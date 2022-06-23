---
title: 0630. Course Schedule III
weight: 10
tags: [
	"Array",
	"Greedy"
]
---

## Description
> There are n different online courses numbered from 1 to n. You are given an array courses where courses[i] = [durationi, lastDayi] indicate that the ith course should be taken continuously for durationi days and must be finished before or on lastDayi.
> 
> You will start on the 1st day and you cannot take two or more courses simultaneously.
> 
> Return the maximum number of courses that you can take.

## Solutions
### Greedy Algorithms
这里的课程调度更前面的课程调度的题的解法都不太一样，这里采用的是贪心算法加最大堆。首先对按照结束时间对各个课程从小到达拍个序，然后遍历每个课程，将每个课程的耗时加入到 courseTaken 里面，并且将耗时也加入到最大堆中，当遍历到当前课程时，累计要花的时间 courseTaken 超过当前课程的结束时间时，我们就从最大堆中取出一个最大值，然后从 courseTaken 中减去这个最大值，这样的话，我们就可以尽量多上几门课。
```go
func scheduleCourse(courses [][]int) int {
    sort.Slice(courses, func(i,j int) bool {
        return courses[i][1] < courses[j][1]
    })
    
    var courseTaken int
    var maxHeap MaxHeap
    heap.Init(&maxHeap)
    for _, course := range courses {
        courseTaken += course[0]
        heap.Push(&maxHeap, course[0])
        if courseTaken > course[1] {
            c := heap.Pop(&maxHeap).(int)
            courseTaken -= c
        }
    }
    
    return maxHeap.Len()
}

type MaxHeap []int
func (m MaxHeap) Len() int {
    return len(m)
}

func (m MaxHeap) Less(i,j int) bool {
    return m[i] > m[j]
}

func (m MaxHeap) Swap(i,j int) {
    m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {
    *m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
    x := (*m)[m.Len()-1]
    *m = (*m)[:m.Len()-1]
    return x
}
```
