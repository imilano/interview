---
title: 0304. Range Sum Query - Mutable
weight: 10
tags: [
	"Prefix Sum"
]
---

## Description
> Given an integer array nums, handle multiple queries of the following types:
> 
> - Update the value of an element in nums.
> - Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
> Implement the NumArray class:
> 
> - `NumArray(int[] nums)` Initializes the object with the integer array nums.
> - `void update(int index, int val)` Updates the value of nums[index] to be val.
> - `int sumRange(int left, int right)` Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

## Solutions

### Prefix Sum
这里参考 303 题和 304 题，很容易想出使用前缀和的解法。区别只是，在更新的时候，需要更新下标`index` 之后的每个前缀和，也正是因为这个更新操作，所以时间会高一些，但是可以通过 OJ。
```go
type NumArray struct {
    nums []int
    prefix []int
}


func Constructor(nums []int) NumArray {
    size := len(nums)
    prefix := make([]int, size +1)
    for idx, num := range nums {
        prefix[idx+1] = prefix[idx] + num
    }
    
    return NumArray{nums, prefix}
}


func (this *NumArray) Update(index int, val int)  {
    diff := val - (*this).nums[index]
    (*this).nums[index] = val
    for i := index+1; i <= len((*this).nums); i++ {
        (*this).prefix[i] += diff
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    return (*this).prefix[right+1] - (*this).prefix[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
```
