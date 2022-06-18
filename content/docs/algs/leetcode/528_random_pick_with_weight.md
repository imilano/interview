---
title: 0528. Random Pick with Weight
weight: 10
tags: [
	"Search",
	"Binary Search"
]
---

## Description
> You are given a 0-indexed array of positive integers w where w[i] describes the weight of the ith index.
> 
> You need to implement the function pickIndex(), which randomly picks an index in the range [0, w.length - 1] (inclusive) and returns it. The probability of picking an index i is w[i] / sum(w).
> 
> - For example, if w = [1, 3], the probability of picking index 0 is 1 / (1 + 3) = 0.25 (i.e., 25%), and the probability of picking index 1 is 3 / (1 + 3) = 0.75 (i.e., 75%).

## Solutions
### Binary Search
这里的解法比较巧妙。首先这里每个数字被取到的比例是按照该数字的权重来计算的。比如说对于[1,3,2] 这是三个数字，每个数字取到的比例就是 1/6、3/6、2/6，这样的话，我们就不能按照平常的以下标作为取随机数的范围来计算。这里用到了一个小技巧：既然这里的和是 6，那么我们就以 6 作为 random 的边界，然后取一个随机数，如果取出的数是 0，那么就选第一个数 1；如果取出的数是 1、2、3，那么就取第二个数 3；如果取出的数是 4、5，那么就取第三个数 2。那么，我们就需要计算一个累加数组，并且让累加数组的下标与原数组的下标对应，**这样的话只需要查找第一个大于随机数 x 数字的下标即可**，该下标即为结果。查找的过程，我们可以使用二分查找来进行。
```go
type Solution struct {
    sums []int
}


func Constructor(w []int) Solution {
    sum := w
    size := len(w)
    for i := 1; i < size; i++ {
        sum[i] = sum[i-1] + w[i]
    }
    
    return Solution{sum}
}


func (this *Solution) PickIndex() int {
    size := len((*this).sums)
    r := rand.Intn((*this).sums[size-1])
    left, right := 0, size - 1
    for left < right {
        mid := left + (right - left) /2
        if (*this).sums[mid] <= r {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return right
}
/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

```
