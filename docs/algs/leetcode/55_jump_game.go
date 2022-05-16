package leetcode

/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.
*/

func canJump(nums []int) bool {
	return dynamicProgrammingCanJump(nums)
}

// 贪心算法。
// 我们只对最远能到到的位置感兴趣，那么可以维护一个变量reach，表示在当前位置能跳到的最远位置，初始化为0 。所有小于reach的位置都是可以到达的，此时可以用
// reach 和 reach + num[i] 中的最大值来更新reach，一旦有一个地方出现了reach大于n-1，那么就直接break，否则的话继续更新reach。最后如果reach值小于n-1，则说明无法到达最后位置，否则说明能够到达。
func greedyCanJump(nums []int) bool {
	var reach int
	size := len(nums)
	for i := 0; i < size; i++ {
		if reach >= size-1 || i > reach {
			break
		}
		reach = max(reach, i+nums[i])

	}

	return reach >= size-1
}

// 动态规划
// 维护一个dp数组，其中dp[i]表示到达i位置所剩余的跳力，若到达某个位置时跳力为负，则说明无法到达该位置。
// 当前位置的跳力跟上一个位置的跳力以及上一个位置剩余的跳力有关，所以当前位置的剩余跳力和当前位置新的跳力中较大的那个数决定了当前能达到
// 的最远距离，而下一个位置的剩余跳力等于当前位置的跳力减去1，所以 dp[i] = max(dp[i-1], nums[i-1]) - 1。如果当前dp为负，则说明无法抵达当前位置，返回false。
func dynamicProgrammingCanJump(nums []int) bool {
	size := len(nums)
	dp := make([]int, size)
	for i := 1; i < size; i++ {
		dp[i] = max(dp[i-1], nums[i-1]) - 1
		if dp[i] < 0 {
			return false
		}
	}

	return true
}

// 递归跳
func canJumpRecursive(nums []int) bool {
	size := len(nums)
	return canJumpHelper(nums, 0, size)

}

func canJumpHelper(nums []int, start int, size int) bool {
	if start >= size {
		return true
	}

	if start == size-1 && nums[start] == 0 {
		return true
	}

	for i := 1; i <= nums[start]; i++ {
		if canJumpHelper(nums, start+i, size) {
			return true
		}
	}

	return false
}
