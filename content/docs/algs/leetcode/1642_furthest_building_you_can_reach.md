---
title: 1642. Furthest Building you Can Reach
weight: 10
tags: [
	"Array",
	"Heap",
	"Priority Queue"
]
---
## Description
> You are given an integer array heights representing the heights of buildings, some bricks, and some ladders.
> 
> You start your journey from building 0 and move to the next building by possibly using bricks or ladders.
> 
> While moving from building i to building i+1 (0-indexed),
> 
> - If the current building's height is greater than or equal to the next building's height, you do not need a ladder or bricks.
> - If the current building's height is less than the next building's height, you can either use one ladder or (h[i+1] - h[i]) bricks.
> Return the furthest building index (0-indexed) you can reach if you use the given ladders and bricks optimally.

## Solutions
### Greedy && Heap
这道题的贪心思路应该是比较容易想到的：好钢一定要用在刀刃上，而梯子就是一块好钢，我们要尽可能用在高度差比较大的地方。或者也可以这么说：尽可能用完每一块砖块，只有砖块实在不够的时候，才考虑来使用梯子。那么接下来问题就来了，既然要把梯子尽可能用在高度差比较大的地方，那么什么地方才算是高度差比较大呢？也就是说，我们如何才能让计算机知道这里高度差比较大呢？很不幸，计算机不知道，没有后视之明。不过这里可以用一个小技巧：刚开始优先使用梯子，对于每个高度差，都将该高度差放入到优先度列中；当队列长度超过 l 时，说明此时的梯子已经被用完了，那么接下来我们就需要把那些大材小用的梯子用砖块替换掉，统计替换这个梯子所需要的砖块数并与我们已有的砖块数进行比较，如果某一个时刻这个累加值超过了我们拥有的砖块数，那么说明我们此时就无法再向前移动了，返回此时的下标减去 1 即可。如果我们能到达最后一个建筑物，那么直接返回长度减 1  即可。

```go
func furthestBuilding(heights []int, bricks int, ladders int) int {
    size := len(heights)
    var deltaH int
    var minHeap MinHeap
    heap.Init(&minHeap)
    
    for i := 1; i < size; i++ {
        // 计算高度差，如果是高到低或者平移，则直接继续遍历
        x := heights[i] - heights[i-1]
        if x > 0 {
            // 这里优先使用梯子，只有当没有梯子使用的时候，才开始考虑使用砖块
            heap.Push(&minHeap, x)
            if minHeap.Len() > ladders {
				// 这里就相当于把那些大材小用的梯子用砖块来代替掉，需要用多少砖块来代替呢？
				// 最小堆中的堆顶元素就是我们所需要的砖块
                deltaH += heap.Pop(&minHeap).(int)
                // deltaH += t.(int)
            }
            
			// 当我们有足够的砖块可以用来代替梯子的时候，我们就可以继续进行比那里；
			// 否则当砖块不足的时候，说明此时我们已经无法继续向前走了
            if deltaH > bricks {
                return i - 1
            }
        }
    }
    
	// 我们可以走到最后一个建筑物
    return size - 1
}

type MinHeap []int

func (m MinHeap) Len() int {
    return len(m)
}

func (m MinHeap) Less(i,j int) bool {
    return m[i] < m[j]
}

func (m MinHeap) Swap(i,j int) {
    m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
    *m = append(*m, x.(int))
}

func (m *MinHeap) Pop() interface{} {
    size := m.Len()
    x := (*m)[size-1]
    *m = (*m)[:size-1]
    
    return x
}
```
