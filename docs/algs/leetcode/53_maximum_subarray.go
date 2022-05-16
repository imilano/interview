package leetcode

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.
*/
// 扫描一次数组，在每一个扫描点计算以该点数值为结束点的子数列的最大和。该子数列由两部分组成：以前一个位置为结束点的最大子数列、该位置的数值。
// 这个算法用到了最佳子结构（以每个位置为重点的最大子数列都是基于前一位置的最大子数列计算得到）。
// see here: https://zh.wikipedia.org/wiki/%E6%9C%80%E5%A4%A7%E5%AD%90%E6%95%B0%E5%88%97%E9%97%AE%E9%A2%98
func maxSubArray(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	res, curSum := nums[0], nums[0]
	for i := 1; i < size; i++ {
		curSum = max(curSum+nums[i], nums[i])
		res = max(res, curSum)
	}

	return res
}
