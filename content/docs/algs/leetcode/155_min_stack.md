---
title: 0155. Min Stack
weight: 10
tags: [
	"Stack",
	"Monotone Stack"
]
---
## Desctiption
> Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
> 
> Implement the MinStack class:
> 
> - MinStack() initializes the stack object.
> - void push(int val) pushes the element val onto the stack.
> - void pop() removes the element on the top of the stack.
> - int top() gets the top element of the stack.
> - int getMin() retrieves the minimum element in the stack.


## Solutions
除了维持一个正常的栈来维持压入的元素之外，还需要维持一个单调递减栈当前的最小数。具体做法是，当压入一个元素的时候，如果单调栈中元素为空，则直接压入元素到单调栈中；如果单调中元素不为空，那么又分为单调栈顶元素比当前压入元素大还是小两种情况，如果栈顶元素比当前压入元素大，那么将当前压入元素压入单调栈；如果栈顶元素比当前压入元素小，那么再次压入栈顶元素。弹出时，除了对正常栈进行弹出之外，还需要对单调栈进行弹出操作。getMin 函数直接返回单调栈顶元素即可。
```go
// 维持一个单调栈，单调递减栈即可
type MinStack struct {
    nums []int
    monoStack []int
}


func Constructor() MinStack {
    return MinStack{nums: []int{}, monoStack: []int{}}
}


func (this *MinStack) Push(val int)  {
    size := this.Len()
    if size == 0 {
        (*this).monoStack = append((*this).monoStack, val)
    } else {
        top := (*this).monoStack[size-1]
        if top <= val {
            (*this).monoStack = append((*this).monoStack, top)
        } else {
            (*this).monoStack = append((*this).monoStack, val)
        }
    }
    
    (*this).nums = append((*this).nums, val)
}

func (this *MinStack) Len() int  {
    return len(this.nums)
}


func (this *MinStack) Pop()  {
    size := this.Len()
    (*this).nums = (*this).nums[:size-1]
    (*this).monoStack = (*this).monoStack[:size-1]
}


func (this *MinStack) Top() int {
    size := this.Len()
    return (*this).nums[size-1]
}


func (this *MinStack) GetMin() int {
    size := this.Len()
    return (*this).monoStack[size-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
 ```
