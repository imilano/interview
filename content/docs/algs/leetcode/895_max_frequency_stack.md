---
title: 0895. Max Frequency Stack
weight: 10
tags: [
	"Stack",
	"Hash Table"
]
---

## Description
> Design a stack-like data structure to push elements to the stack and pop the most frequent element from the stack.
> 
> Implement the FreqStack class:
> 
> FreqStack() constructs an empty frequency stack.
> void push(int val) pushes an integer val onto the top of the stack.
> int pop() removes and returns the most frequent element in the stack.
> If there is a tie for the most frequent element, the element closest to the stack's top is removed and returned.


## Solutions
首先可以很明确的是，一定要统计每个数字出现的次数。那么现在我们知道了每个数字出现的次数之后，如何知道对于指定的出现次数中会有哪些数字呢，并且这写数字还有一定的先后顺序？这个也可以通过一个 Hash Table 来解决，其中 key 为出现的次数，而 value 为一个数组，表示对应于该出现次数的数字。每次出现次数到达 key 时，我们就可以将对应的值 append 到 value 数组中，这样我们就可以知道对于特定的出现次数都包含了哪些数字。

这里维持一个最大出现频率 maxFreq，那么接下来对应 push 操作，只需要将freq 中的对应数字的计数值增加，然后在 dict 中使用这个计数值作为 key，将当前值 append 到 value 中；对于 pop 操作，直接从 dict[maxFreq] 中取出最后一个数字 x，然后检查该频率对应的数组中是否还有数字，如果没有数字了，说明 maxFreq 对应频率的数字已经被取完了，接下来应该去 maxFreq - 1 的数组中去取，那么将 maxFreq 减去 1，此外还需要在 freq 中将对应的值的频率减去 1，最后返回 x 即可。
```go
type FreqStack struct {
    maxFreq int
    freq map[int]int
    dict map[int][]int
}


func Constructor() FreqStack {
    var maxFreq int
    dict := make(map[int][]int)
    freq := make(map[int]int)
    
    return FreqStack{maxFreq, freq, dict}
}


func (this *FreqStack) Push(val int)  {
	//  更新频率
    (*this).freq[val]++
	// 更新当前最大频率
    (*this).maxFreq = max((*this).maxFreq, (*this).freq[val])
	// 将 val 添加到对应频率的数组中
    (*this).dict[(*this).freq[val]] = append((*this).dict[(*this).freq[val]], val)
}


func (this *FreqStack) Pop() int {
	// 在最大频率数组中取出最后一个数字
    size := len((*this).dict[(*this).maxFreq])
    x := (*this).dict[(*this).maxFreq][size-1]
	// 缩小最大频率数组的长度
    (*this).dict[(*this).maxFreq] = (*this).dict[(*this).maxFreq][:size-1]
	// 如果最大频率数组中已经没有数字中，那么最大频率应该降低
	// 因为频率是连续的，所以下面这一步可以确保上面的最大频率数组中不会取到空值
    if len((*this).dict[(*this).freq[x]]) == 0 {
        (*this).maxFreq--
    }
	// 将对应数组的频率降低
    (*this).freq[x]--
    return x
    
}

func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}
/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */

```
