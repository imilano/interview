package leetcode

import "sort"

/*
Given an array of distinct integers nums and a target integer target, return the number of possible combinations that add up to target.

The test cases are generated so that the answer can fit in a 32-bit integer.
*/

func combinationSum4(nums []int, target int) int {
	return combinationSum4Solution2(nums, target)
}

// 动态规划
// 由于需要考虑元素的选取顺序，所以这题需要计算的是选取元素的排列数
// 定义 dp[x] 表示选取的元素之和等于x的方案数，则目标是求dp[target]。
// 边界条件是，当不选取任何元素，也就是x=0时，选取方案只有1种（不选取），所以dp[0] = 1
// 当 1≤ i ≤target 时，如果存在一种排列，其中的元素之和等于 i，则该排列的最后一个元素一定是数组 nums 中的一个元素。假设该排列的最后一个元素是 num，则一定有 num≤i，对于元素之和等于 i−num 的每一种排列，在最后添加 num 之后即可得到一个元素之和等于 i 的排列，因此在计算 dp[i] 时，应该计算所有的 dp[i−num] 之和。

// 由此可以得到动态规划的做法：
// 	- 初始化 dp[0]=1；

// 	- 遍历 i 从 1 到 target，对于每个 i，进行如下操作：

// 		- 遍历数组 nums 中的每个元素num，当 num≤i 时，将 dp[i−num] 的值加到 dp[i]。
// 	- 最终得到 dp[target] 的值即为答案。

func combinationSum4Solution2(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}

// 回溯法
// 由于题目中相同序列的不同组合也视为不同的排列，所以这里需要对回溯得到的结果做一个全排列，全排列中的值都可能是我们需要的值，可以用 map 进行去重
func combinationSum4Solution1(nums []int, target int) int {
	dict := make(map[string]bool)
	size := len(nums)
	sort.Ints(nums)

	var out []int
	recursiveCombinationSum4(nums, 0, target, size, out, &dict)

	return len(dict)
}

func recursiveCombinationSum4(nums []int, start int, target int, size int, out []int, dict *map[string]bool) {
	if start >= size {
		s := arrToString(out)
		if _, ok := (*dict)[s]; !ok {

		}
	}

	for i := start; i < size; i++ {
		if nums[i] <= target {
			newArr := make([]int, len(out))
			copy(newArr, out)
			newArr = append(newArr, nums[i])
			recursiveCombinationSum4(nums, i, target-nums[i], size, newArr, dict)
		}
	}
}

// 做全排列
func recursiveCombinationSum4Helper(out []int, dict *map[string]bool) {
	// TODO
}

// for test
func CombinationSum4(nums []int, target int) int {
	return combinationSum4(nums, target)
}
