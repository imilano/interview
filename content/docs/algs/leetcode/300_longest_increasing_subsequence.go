package leetcode

func lengthOfLIS(nums []int) int {
	return lengthOfLISSolution2((nums))
}

/*
动态规划，维护一个一维 dp 数组，其中 dp[i] 表示以nums[i]为结尾的最长递增子序列的长度，对于每一个 nums[i], 从第一个数再搜到 i,如果发现某个数小于 nums[i]，更新 dp[i].
更新方法为 dp[i] = max(dp[i], dp[j]+1)(仔细体会)，其实就是扫描的时候如果发现有一个数比当前 nums[i]小，那么以这个数为结尾的子序列就可以和 nums[i]构成一个递增子序列，
由于要求最大值，所以再对 dp[i] 和 dp[j]+1取个最值作为当前dp[i] 的值即可。这样不断的更新 dp 数组，到最后 dp 数组中最大的值就是要返回的 LIS 的长度。*/
func lengthOfLISSolution2(nums []int) int {
	// var res int
	// 由于第一层循环是从 1 开始的，所以这个地方需要初始化为 1，否则当数组长度为 1 时，可能会返回 0，而这是错误的
	res := 1
	size := len(nums)
	dp := make([]int, size)
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 1; i < size; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		// 注意这里，最后所求的并不是dp[size-1]
		res = max(res, dp[i])
	}

	return res
}

/*
错误解法： 实际上下面的解法是错误的，因为f[i]的值不仅仅收到 dp[i-1] 以及 nums[i] 和 nums[i-1] 的相对大小影响，还收到之前的最长子序列中的最大值影响。
也就是说，当前值是否能够为最大子序列长度增加长度，其实不仅仅与前一个位置有关，还与整个序列的最大值有关，单独使用nums[i] 与 nums[i-1]来比较大小是不对的。
使用动态规划，假设 f[i] 表示 在数组中第 i 个位置能取到最大子数组长度，那么可以知道，f[i] 的值受到 f[i-1]  和 nums[i] 与 nums[i-1] 之间相对大小的影响。
当 nums[i] > nums[i-1] 时， f[i] = f[i-1] + 1；否则 f[i] = f[i-1]。最终 f[i] 就是结果值。
以上就可以得出递归公式，并且很容易就可以知道, f[0] = 1
*/
func lengthOfLISSolution1(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	dp[0] = 1
	for i := 1; i < size; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[size-1]
}
