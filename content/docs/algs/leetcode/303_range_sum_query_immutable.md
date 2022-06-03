---
title: 0303. Range Sum Query - Immutable
weight: 10
tags: [
	"DP",
	"Prefix Sum"
]
---
## Solutions

### Brute Force & Memorial
本来题主以为这里考的是使用记忆化数组来避免重复计算，结果一提交结果，发现自己还是太天真了...
```go
type NumArray struct {
    nums []int
    dict map[string]int
}


func Constructor(nums []int) NumArray {
    return NumArray{nums, make(map[string]int)}
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == right {
        return (*this).nums[left]
    }
    
    if left > right {
        return 0
    }
    
    target := string(left) + string(right)
    if _, ok := (*this).dict[target]; ok {
        return (*this).dict[target]
    }
    
    res := this.nums[left] + this.nums[right] + this.SumRange(left+1, right-1)
    (*this).dict[target] = res
    return res
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
 ```

### Prefix Sum
网上看到有用前缀和的解法，只能说我实在是想不到....。总的来说就是，创建一个数组 predix 记录 [0,right] 的和，那么当要求计算 [left,right] 的和时，只需要使用[0,right]的和减去[0,left-1]即可。而为了方便计算，我们的数组增加了一个位置，这样在计算结果的时候，下标也后移 1，就会变成[0,right+1]的和减去[0,left] 的和
```go
type NumArray struct {
    nums []int
    prefix []int
}


func Constructor(nums []int) NumArray {
    prefix := make([]int, len(nums)+1)
    idx := 1
    var res int
    // [0,right] 的和为 prefix[right+1]
    for _, num := range nums {
        res += num
        prefix[idx] = res
        idx++
    } 
    return NumArray{nums, prefix}
}

// [0,right] 的和为 prefix[right+1]，而[left,right] 的和为[0, right]的和 减去 [0,left-1] 的和，
// 对应到这里，即为 prefix[right+1] - prefix[left]
func (this *NumArray) SumRange(left int, right int) int {
    return (*this).prefix[right+1] - (*this).prefix[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
 ```
