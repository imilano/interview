---
title: 0346. Moving Average from Data Stream
weight: 10
tags: [
	"Queue",
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
```go
type MovingAverage struct {
	nums []int
	size int
	sum int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{nums: []int{}, size: size, sum: 0}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.size < 3 {
		*this.nums  = append(*this.nums, val)
		sum += val
		this.size++
	} else {
		num := (*this).nums[0]
		*this.nums = (*this).nums[1:]
		sum -= num
		*this.nums = append(*this.nums, val)
		sum += val
	}

	return float64(sum/this.num)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
 ```
