package leetcode

/*
Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

A subarray is a contiguous subsequence of the array.
*/

func maxProduct(nums []int) int {
	return maxProductSolution2(nums)
}

// 方法 2，使用动态规划。从左向右扫描，使用两个数组 greater 和 lesser 分别记录包含当前元素 nums[i] 的最大子数组和最小子数组乘积，二者初始化都是 nums 的第一个元素。
// 那么从第二个元素开始，当前的最大/小子数组乘积只会在 greater[i-1]*nums[i]、lesser[i-1]*nums[i]  和 nums[i] 这三者中产生。使用一个 res 不断标记当前最大值即为最后结果。
func maxProductSolution2(nums []int) int {
	size := len(nums)

	greater, lesser := make([]int, size), make([]int, size)
	greater[0], lesser[0] = nums[0], nums[0]
	res := nums[0]
	for i := 1; i < size; i++ {
		greater[i] = max(greater[i-1]*nums[i], max(lesser[i-1]*nums[i], nums[i]))
		lesser[i] = min(greater[i-1]*nums[i], min(lesser[i-1]*nums[i], nums[i]))
		res = max(res, greater[i])
	}

	return res
}

// 方法 3，遍历两次，正向遍历一次数组，更新累积积的最大值，然后在反向遍历一次，更新累积积的最大值。
// 之所以要遍历两次，是因为如果负数出现的次数为奇数，那么单独一次正向遍历的结果可能并不会反应真实的最大值，也就说，负数出现的次数和位置很重要，所以要反向遍历一次勘误。
func maxProductSolution3(nums []int) int {
	size := len(nums)
	res, prod := nums[0], 1
	for i := 0; i < size; i++ {
		prod *= nums[i]
		res = max(res, prod)
		if prod == 0 {
			prod = 1
		}
	}

	prod = 1
	for i := size - 1; i >= 0; i-- {
		prod *= nums[i]
		res = max(res, prod)
		if prod == 0 {
			prod = 1
		}
	}

	return res
}

// 错误方法
// 方法 1，根据第 53 题的解法，可以发现这二者还是存在一定的相似之处的，只不过 53 题是加法，而这里是乘法。53 题的解法可以应对出现 0 或者负数的情况，但是对于乘法而言，这两种情况会有所不同。
// 盲目套用第 53 题的方法的都的答案是错误的。
func maxProductSolution1(nums []int) int {
	res, curProduct := nums[0], nums[0]
	size := len(nums)
	for i := 1; i < size; i++ {
		curProduct = max(curProduct*nums[i], nums[i])
		res = max(res, curProduct)
	}

	return res
}
