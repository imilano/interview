---
title: 0232. Implement Queue using Stacks
weight: 10
tags: [
	"Queue"
]
---
## Description
> Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).
> 
> Implement the MyQueue class:
> 
> void push(int x) Pushes element x to the back of the queue.
> int pop() Removes the element from the front of the queue and returns it.
> int peek() Returns the element at the front of the queue.
> boolean empty() Returns true if the queue is empty, false otherwise.
> Notes:
> 
> You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty operations are valid.
> Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.

## Solutions
这个目前也很简单，主要实现如下：使用两个队列，队列1专用于push，队列用专用于pop 和 peek。push 时，入队 stack1； pop 时，如果 stack2 不为空，则弹出 stack2 栈顶即可，若 stack2 为空，则将 stack1 的所有元素压入 stack2， 然后将 stack2 的栈顶元素弹出。
```go
type MyQueue struct {
	stack1 []int
	stack2 []int
}

// comment this to resolve conflicts
func Constructor() MyQueue {
    return MyQueue{
        stack1: make([]int, 0),
        stack2: make([]int, 0),
    }
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	var res int
	len1, len2 := len(this.stack1), len(this.stack2)
	if len2 != 0 {
		ele := this.stack2[len2-1]
		res = ele
		this.stack2 = this.stack2[:len2-1]
	} else {
		for i := len1 - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		this.stack1 = this.stack1[:0]
		len2 = len(this.stack2)
		res = this.stack2[len2-1]
		this.stack2 = this.stack2[:len2-1]
	}

	return res
}

func (this *MyQueue) Peek() int {
	var res int
	len1, len2 := len(this.stack1), len(this.stack2)
	if len2 != 0 {
		ele := this.stack2[len2-1]
		res = ele
	} else {
		for i := len1 - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}

		this.stack1 = this.stack1[:0]
		len2 = len(this.stack2)
		res = this.stack2[len2-1]
	}

	return res
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}
```
