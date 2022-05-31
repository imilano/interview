---
title: 0033. Search in Rotated Sorted Array
weight: 10
tags: [
	"Binary Search",
	"Search"
]
---

## Description

> There is an integer array nums sorted in ascending order (with distinct values).
> 
> Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
> 
> Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.
>
> You must write an algorithm with O(log n) runtime complexity.


## Solutions

### Binary Search

引用别人的话：
> 将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
> 此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环. 

还是使用二分查找的方法来。首先，如果这个数组翻转了，那么至少有有一侧是占多数数字的，具体那哪一侧还要进一步判断。经过举例观察可以看到，如果中间数小于右边的数，那么右边一定是有序的；如果中间数大于右边的数，那么左半段是有序的。那么我们只需要在有序的区间内进行二分查找即可。详见代码。

```go
func search(nums []int, target int) int {
	size := len(nums)
	left, right := 0, size-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} 
		
		if nums[mid] < nums[right] {
			// 如果右半边有序，则对右半边进行二分
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 如果左半边有序，则对左半边进行二分
			if nums[left] <= target && nums[mid] > target {
				right = mid -1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
```
