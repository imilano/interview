---
title: 0767. Reorganize String
weight: 10
tags: [
	"String",
	"Heap"
]
---
## Description
> Given a string s, rearrange the characters of s so that any two adjacent characters are not the same.
> 
> Return any possible rearrangement of s or return "" if not possible.

## Solutions
这里的解法还是很巧妙的，需要注意一下。
```go
type Pair struct {
    char byte
    cnt int
}

func reorganizeString(s string) string {
    dict := make(map[byte]int)
    for idx, _ := range s {
        dict[s[idx]]++
    }
    
    var maxHeap MaxHeap
    heap.Init(&maxHeap)
    for char, cnt := range dict {
        if cnt > (len(s)+1)/2 {
            return ""
        }
        
        heap.Push(&maxHeap, Pair{char, cnt})
    }
    
    var res string
    for maxHeap.Len() > 1 {
        first, second := heap.Pop(&maxHeap).(Pair), heap.Pop(&maxHeap).(Pair)
        res += string(first.char)
        res += string(second.char)
        
        first.cnt--
        second.cnt--
        if first.cnt > 0 {
            heap.Push(&maxHeap, first)
        }
        
        if second.cnt > 0 {
            heap.Push(&maxHeap, second)
        }
    }
    
    if maxHeap.Len() > 0 {
        first := heap.Pop(&maxHeap).(Pair)
        res += string(first.char)
    }
    
    return res
}

type MaxHeap []interface{}

func(m MaxHeap) Len() int {
    return len(m)
}

func (m MaxHeap) Swap(i,j int) {
    m[i], m[j] = m[j], m[i]
}

// > 才是最大堆， < 是最小堆
func (m MaxHeap) Less(i,j int) bool {
    return m[i].(Pair).cnt > m[j].(Pair).cnt
}

func (m *MaxHeap) Push(x interface{}) {
    *m = append(*m, x)
}

func (m *MaxHeap) Pop() interface{} {
    x := (*m)[m.Len()-1]
    *m = (*m)[:m.Len()-1]
    return x
}
```
