package leetcode

/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
*/

// 求一个连续数组的最大值，很容易能够想到用dp来求解。问题是怎么计算状态转移方程。
// 首先， 对于一个房子i，只有两种选择，要么抢，要么不抢，只有这两种选择。定义一个数组dp,dp[i]表示从[0,i]区间内能抢到的最大金额。
// 那么如果不抢当前房子的话， dp[i] = dp[i-1]，而如果抢的话，则dp[i] = dp[i-2] +  nums[i]，注意这里是 i-2，因为不能抢相邻的房子。
// 基于以上分析，就可以得到 dp[i] = max(dp[i-1], dp[i-2]+nums[i])。
// 初始条件 dp[0] = nums[0], dp[1] = max(nums[0], nums[1])
func rob(nums []int) int {
	return robSolution1(nums)
}

// 动态规划
func robSolution1(nums []int) int {
	size := len(nums)
	if size < 1 {
		return 0
	}

	if size == 1 {
		return nums[0]
	}

	// 注意这里的初始化条件
	dp := make([]int, size)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < size; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[size-1]
}

// 核心思想还是dp，分别维护两个变量 robEven 和 robOdd，表示抢偶数位置的房子和抢奇数位置的房子。
func robSolution2(nums []int) int {
	size := len(nums)
	var robEven, robOdd int

	for i := 0; i < size; i++ {
		if i%2 == 0 {
			robEven = max(robEven+nums[i], robOdd)
		} else {
			robOdd = max(robOdd+nums[i], robEven)
		}
	}

	return max(robEven, robOdd)
}
