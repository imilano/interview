package leetcode

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}

func maxProfit(prices []int) int {
	return maxProfitSolution3(prices)
}

// 方法 3， 扫描一遍数组，用一个值记录当前扫描到的最小值，每个当前扫描到的值都减去最小值，然后取其最大结果即可
func maxProfitSolution3(prices []int) int {
	var res int
	size := len(prices)
	if size < 2 {
		return res
	}

	minVal := prices[0]
	for _, value := range prices {
		res = max(res, value-minVal)
		if value < minVal {
			minVal = value
		}
	}

	return res
}

// 错误方法，因为这两个值不一定是一个在左半边，一个在右半边。很可能出现的情况是，两个值都出现在左半边
// 方法 2，左指针从左往右扫，找最小值，右指针从右往左扫，找最大值，直到两个指针相交
func maxProfitSolution2(prices []int) int {
	size := len(prices)
	if size < 2 {
		return 0
	}

	// for 循环在 init 之后检查 condition，condition 通过之后会执行 body，body 执行之后会进行 increment，然后再比较 condition，如果 condition 通过，继续执行 body
	left, right := 0, size-1
	for l, r := 0, size-1; l <= r; l, r = l+1, r-1 {
		if prices[l] < prices[left] {
			left = l
		}

		if prices[r] > prices[right] {
			right = r
		}
	}

	if prices[right] < prices[left] {
		return 0
	}

	return prices[right] - prices[left]
}

// 方法 1，暴力循环
func maxProfitSolution1(prices []int) int {
	var res int
	length := len(prices)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			res = max(res, prices[j]-prices[i])
		}
	}

	return res
}
