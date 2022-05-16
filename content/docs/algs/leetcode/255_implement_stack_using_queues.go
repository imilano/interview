package leetcode

/*
Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).

Implement the MyStack class:

void push(int x) Pushes element x to the top of the stack.
int pop() Removes the element on the top of the stack and returns it.
int top() Returns the element on the top of the stack.
boolean empty() Returns true if the stack is empty, false otherwise.
Notes:

You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
Depending on your language, the queue may not be supported natively. You may simulate a queue using a list or deque (double-ended queue) as long as you use only a queue's standard operations.
*/

// 主要思路：使用两个队列，所有数据都主要存储在一个队列queue1中，然后使用另一个队列queue2来做辅助队。在做 push 操作时，将待入队的元素插入到非空的队列末尾（此时只有一个）；
// 在做 pop 操作时，装有元素的那个队列中除最后一个元素之外的元素都放到另一个队列中，记录最后一个元素的值并删除该元素，最后返回该值即可。
// 在做 top 操作时，将装有元素的那个队列中所有元素都放到另一个队列中，但是记录最后一个元素的值，最后返回该元素即可。

type MyStack struct {
	queue1 []int
	queue2 []int
}

// comment this to resolve conflicts
// func Constructor() MyStack {
//     return MyStack{
//         queue1: make([]int, 0),
//         queue2: make([]int, 0),
//     }
// }

func (this *MyStack) Push(x int) {
	if len(this.queue1) != 0 {
		this.queue1 = append(this.queue1, x)
	} else {
		this.queue2 = append(this.queue2, x)
	}
}

func (this *MyStack) Pop() int {
	var res int
	len1, len2 := len(this.queue1), len(this.queue2)
	if len1 != 0 {
		for i := 0; i < len1-1; i++ {
			ele := this.queue1[0]
			this.queue2 = append(this.queue2, ele)
			this.queue1 = this.queue1[1:]
		}

		res = this.queue1[0]
		this.queue1 = this.queue1[1:]
	} else {
		for i := 0; i < len2-1; i++ {
			ele := this.queue2[0]
			this.queue1 = append(this.queue1, ele)
			this.queue2 = this.queue2[1:]
		}

		res = this.queue2[0]
		this.queue2 = this.queue2[1:]
	}

	return res
}

func (this *MyStack) Top() int {
	var res int
	len1, len2 := len(this.queue1), len(this.queue2)
	if len1 != 0 {
		for i := 0; i < len1; i++ {
			ele := this.queue1[0]
			this.queue2 = append(this.queue2, ele)
			this.queue1 = this.queue1[1:]
			if i == len1-1 {
				res = ele
			}
		}
	} else {
		for i := 0; i < len2; i++ {
			ele := this.queue2[0]
			this.queue1 = append(this.queue1, ele)
			this.queue2 = this.queue2[1:]
			if i == len2-1 {
				res = ele
			}
		}
	}

	return res
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0 && len(this.queue2) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
