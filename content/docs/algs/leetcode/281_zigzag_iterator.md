---
title: 0281. Zigzag Iterator
weight: 10
tags: [
	"Queue",
]
---
## Description
> Given two 1d vectors, implement an iterator to return their elements alternately.
> 
> For example, given two 1d vectors:
> ```
>   v1 = [1, 2]        
>   v2 = [3, 4, 5, 6]
> ```
> By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1, 3, 2, 4, 5, 6].
> 
> Follow up: What if you are given k 1d vectors? How well can your code be extended to such cases?
> 
> Clarification for the follow up question:
> The "Zigzag" order is not clearly defined and is ambiguous for k > 2 cases. If "Zigzag" does not look right to you, replace "Zigzag" with "Cyclic". For example:
> 
> Input:
> ```
> [1,2,3]
> [4,5,6,7]
> [8,9]
> ```
> Output: 
> ```
> [1,4,8,2,5,9,3,6,7]
> ```


## Solutions
这里可以用队列。首先可以对每一个数组生成一个 pair 对，每个 pair 对包含 两个元素 start 和 end， 其中 start 表示当前应该打印的数字下标，end 表示当前数组的结尾位置（也就是 size），然后按照顺序将每个 pair 入队。读的时候就取出队头元素，如果 start 比 end 要小，那么取出 start 下标的值并输出，然后将 pair 的 start 加 1，放到队尾去；如果 start 等于 end 了，则说明这个数组已经遍历完成了，则从队列总删除这个 pair。

因为不是 premium 会员，所以看不到代码签名，智能子集随便写一下了。
```go
// 下标 0 记录是哪一个数组，下标 1 记录该数组的 start， 下标 2 记录该数组的 end
type Pair [3]int
 
type ZigZagIterator struct {
	queue []Pair
	nums [][]int
}

func Construtor(nums ...[]int) ZigZagIterator {
	var queue []Pair
	for idx, num := range  nums {
		size := len(num)
		queue  = append(queue, []int{idx, 0, size})
	}

	return ZigZagIterator{queue: queue, nums: nums}
}

func (this *ZigZagIterator) Next() int {
	size := len(*this.queue)

	top := (*this).queue[0]
	(*this).queue = (*this).queue[1:]
	for size != 0  && top[1] == top[2] {
		top = (*this).queue[0]
		(*this).queue = (*this).queue[1:]
		size--
	}

	// handle error
	if size == 0 {
		return -1
	}

	x := (*this).nums[top[0]][top[1]]
	(*this).queue = append(*this.queue, []int{top[0], top[1]+1, top[2]})
	return x
}

```
