---
title: 0346. Moving Average from Data Stream
weight: 10
tags: [
	"LinkedList",
	"Two Pointer"
]
---
## Description
> Given a stream of integers and a window size, calculate the moving average of all integers in the sliding window.
> 
> Example:
>''' 
> MovingAverage m = new MovingAverage(3);
> m.next(1) = 1
> m.next(10) = (1 + 10) / 2
> m.next(3) = (1 + 10 + 3) / 3
> m.next(5) = (10 + 3 + 5) / 3
> '''

## Solutions

很明显是使用队列啦。
TODO
```go
type MovingAverage struct {

}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {

}

func (this *MovingAverage) Next(val int) float64 {

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
 ```
