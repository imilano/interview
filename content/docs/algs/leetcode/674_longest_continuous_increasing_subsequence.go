package leetcode

/*
Given an unsorted array of integers nums, return the length of the longest continuous increasing subsequence (i.e. subarray). The subsequence must be strictly increasing.

A continuous increasing subsequence is defined by two indices l and r (l < r) such that it is [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] and for each l <= i < r, nums[i] < nums[i + 1].

*/

func findLengthOfLCIS(nums []int) int {
	size := len(nums)
	if size <= 1 {
		return size
	}
	res, tmp := 1, 1
	for i := 1; i < size; i++ {
		if nums[i] <= nums[i-1] {
			tmp = 1
			continue
		}

		tmp += 1
		res = max(res, tmp)
	}

	return res
}
