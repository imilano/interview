package leetcode

import (
	"math"
	"sort"
)

/*
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.
*/

// 这题乍一看需要使用贪心，但其实贪心算法只能给出局部最优解，而不能给出全局最优解。所以使用贪心可能会出现的情况是：amount 其实完全是可以换完的，但是最后计算出来的结果却是无法换完。
// 所以这意味着我们不能在找到一个可行解之后就停止，而是要寻找到所有的可行解，然后取这些可行解中的最小值
func coinChange(coins []int, amount int) int {
	return coinChangeGreedy(coins, amount)
	// iterative
	// return iterateCoinsChange(coins, amount)
}

func coinChangeGreedy(coins []int, amount int) int {
	res := math.MaxInt
	if amount == 0 {
		return 0
	}

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	coinChangeGreedyAllSolutions(coins, amount, 0, 0, &res)
	if res == math.MaxInt {
		return -1
	}
	return res
}

func coinChangeGreedyAllSolutions(coins []int, amount, cur, index int, res *int) {
	if amount == 0 {
		*res = min(*res, cur)
		return
	}

	if index >= len(coins) {
		return
	}

	for k := amount / coins[index]; k >= 0 && k+cur < *res; k-- {
		coinChangeGreedyAllSolutions(coins, amount-k*coins[index], cur+k, index+1, res)
	}
}

// 动态规划，假设 dp[i] 表示 amount 为 i 时所需要的最小硬币数，那么如何更新 dp[i]呢，结果是遍历每个硬币币值，如果 coins[j] < i，那么可以用 dp[i- coins[j]] + 1 来更新 dp[i] 的值，
// 所以 dp 的状态转移方程是： dp[i] = min(dp[i], dp[i - coins[j]] + 1) if coins[j] <= i。
// 这里需要注意的是，由于求的是 amount 的找零硬币个数，所以创建 dp 的时候，需要创建 amount + 1 个长度。
// 另外，由于硬币的最小单位是 1，所以可能的最多的找零硬币个数是 amount + 1， 所以初始化可以把 dp[i] 设置为任意比 amount 大的值，最后用这个值和 amount 比较，就知道是否需要返回-1
func coinChangeDP(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		memo[i] = math.MaxInt
	}
	return recursiveCoinsChange(coins, amount, &memo)
}
func iterateCoinsChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	size := len(coins)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := 0; j < size; j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 递归解法
// TODO 相关的语法还需要调试
func recursiveCoinsChange(coins []int, target int, memo *[]int) int {
	if target < 0 {
		return -1
	}
	if (*memo)[target] != math.MaxInt {
		return (*memo)[target]
	}
	size := len(coins)
	for i := 0; i < size; i++ {
		tmp := recursiveCoinsChange(coins, target-coins[i], memo)
		if tmp >= 0 {
			(*memo)[target] = min((*memo)[target], tmp+1)
		}
	}

	if (*memo)[target] == math.MaxInt {
		return -1
	}
	return (*memo)[target]
}

// 错误的贪心算法
func greedyCoinChange(coins []int, amount int) int {
	var res int
	sort.Ints(coins)
	index := len(coins) - 1
	for amount != 0 {
		if index < 0 {
			return -1
		}
		count := amount / coins[index]
		res += count
		amount -= count * coins[index]
		index--
	}

	return res
}
