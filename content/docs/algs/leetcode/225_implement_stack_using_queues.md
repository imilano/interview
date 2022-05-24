---
title: 0225. Implement Stack using Queues
weight: 10
tags: [
	"Stack",
	"Queue"
]
---

## Description
> Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).
> 
> Implement the MyStack class:
> 
> void push(int x) Pushes element x to the top of the stack.
> int pop() Removes the element on the top of the stack and returns it.
> int top() Returns the element on the top of the stack.
> boolean empty() Returns true if the stack is empty, false otherwise.
> Notes:
> 
> You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
> Depending on your language, the queue may not be supported natively. You may simulate a queue using a list or deque (double-ended queue) as long as you use only a queue's standard operations.
## Solutions
简单题，可以使用数组模拟双向队列，二双线个队列本来就是一端进一端出的，那么这里只需要把出的一端统一到入的一端即可。
```go
type Queue = []int

type MyStack struct {
    nums Queue
}


func Constructor() MyStack {
    return MyStack{nums: []int{}}
}


func (this *MyStack) Push(x int)  {
    (*this).nums = append((*this).nums, x)
}


func (this *MyStack) Pop() int {
    size := len((*this).nums)
    x := (*this).nums[size-1]
    (*this).nums = (*this).nums[:size-1]
    return x
}


func (this *MyStack) Top() int {
    size := len((*this).nums)
    return (*this).nums[size-1]
}


func (this *MyStack) Empty() bool {
    return len((*this).nums) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
 ```
