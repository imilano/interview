---
title: 0362. Design Hit Counter
weight: 10
tags: [
	"Queue"
]
---
## Description
> Design a hit counter which counts the number of hits received in the past 5 minutes.
> 
> Each function accepts a timestamp parameter (in seconds granularity) and you may assume that calls are being made to the system in chronological order (ie, the timestamp is monotonically increasing). You may assume that the earliest timestamp starts at 1.
> 
> It is possible that several hits arrive roughly at the same time.

## Solutions

简单，使用队列即可。没产生一次 hit，检查队尾元素是否比当前 hit 的 second 要小，如果小，就把 second 入队，同时检查对首元素是否比 second - 300 要大，如果不是，则将这些元素出队即可。

```go
type HitCounter struct {
	queue []int
}

func Constructor() HitCounter{
	return HitCounter{queue: []int{}}
}

func (this *HitCounter) Hit(second int) {
	size := this.Len()
	if size == 0 {
		(*this).queue = append(*this.queue, second)
		return
	} else {
		tail := (*this).queue[size-1]
		if second <= tail {
			return
		}

		var i int
		for i < size {
			if (*this).queue[i] < second - 300 {
				i++
				continue
			}

			break
		}

		(*this).queue = (*this).queue[i:]
		(*this).queue = append(*this.queue, val)
	}
}

func (this *HitCounter) Len() int {
	return len(*this.queue)
}

func (this *HitCounter) GetHits(second int) int {
	return this.Len()
}
```
