package leetcode

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/

// 要知道一个位置上的值，只需要知道这个数前面的数的乘积以及这个数后面的数乘积。所以可以使用两个数组，一个从左向右计算，第 i 个位置存储第 i 个元素的前面的元素的乘积。
// 一个从右向左计算，第 j 个位置存储第 j 个位置之后的元素的乘积。最后结果就是两个数组的积
func productExceptSelf(nums []int) []int {
	size := len(nums)
	if size < 2 {
		return nums
	}

	left, right := make([]int, size), make([]int, size)
	// nums =  append(append([]int{1}, nums...), 1)
	left[0], right[size-1] = 1, 1
	for i, j := 1, size-2; i < size; i, j = i+1, j-1 {
		left[i] = left[i-1] * nums[i-1]
		right[j] = right[j+1] * nums[j+1]
	}

	for i := 0; i < size; i++ {
		left[i] = left[i] * right[i]
	}

	return left
}

func ProductExceptSelf(nums []int) []int {
	return productExceptSelf(nums)
}
