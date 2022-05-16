package leetcode

/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
*/

// 这里特殊之处就是在首尾相连了，如何解决呢？首尾相连会出现的问题是，抢了第一家就不能抢最后一家，不抢第一家就能抢最后一家。
// 那么我们可以算两次，一次不抢第一家，另一次抢第一家，然后分别计算这两种情况下的最大收益，比较二者最大值即为所求。
func robIISolution(nums []int) int {
	size := len(nums)
	if size < 1 {
		return 0
	}

	if size == 1 {
		return nums[0]
	}
	return max(doRob(nums[1:]), doRob(nums[:size-1]))
}

func doRob(nums []int) int {
	size := len(nums)
	if size < 1 {
		return 0
	}

	if size == 1 {
		return nums[0]
	}

	dp := make([]int, size)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < size; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[size-1]
}
