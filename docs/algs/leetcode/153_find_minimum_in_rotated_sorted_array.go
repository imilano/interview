package leetcode

/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.

*/

//TODO 复习这里
// 不是用 left 和 right 进行比较，而是对 mid 和 right 进行比较，如果 mid 小于 right，则 right 等于 mid；如果 mid 大于 right，则 left 等于 mid+1
// 注意，上述一个是等于 mid，一个是等于 mid+1，由于要找的是最小值，稍微思考一下还是会觉得很合理的
func findMin(nums []int) int {
	size := len(nums)
	left, right := 0, size-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

// 错误解法，会出现死循环
// 取左边界 left 和右边界 right， 如果 right 比 left 大，则 left = mid；如果 right 比 left 小，则 right = mid。由于每个元素都是独一无二的，所以不需要考虑其它的情形
func findMinWrong(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}

	left, right := 0, size-1
	for left <= right {
		if left == right {
			break
		}
		if nums[right] > nums[left] {
			right = left + (right-left)/2
		} else {
			left = left + (right-left)/2
		}
	}

	return nums[left]
}

func FindMin(nums []int) int {
	return findMin(nums)
}
