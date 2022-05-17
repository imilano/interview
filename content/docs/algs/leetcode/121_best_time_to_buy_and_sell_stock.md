---
title: 121. Best Time to But And Sell Stock
weight: 10
---
## Description
> You are given an array prices where prices[i] is the price of a given stock on the ith day.
> 
> You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
> 
> Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

## Solutions

首先可以直接暴力求解，但是会超时。
```go
func maxProfit(prices []int) int {
   	var res int
	length := len(prices)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			res = max(res, prices[j]-prices[i])
		}
	}

	return res
}


func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}
```

这里也可以扫描一次数组，然后维护一个扫描过程中遇到的最小值，然后不断计算当前扫描到的值和该最小值的差值，取最大差值即可。
```go
func maxProfit(prices []int) int {
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


func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}
```


一种**错误解法** 是使用双指针，左指针从左到右找最小值，右指针从优到左找最大值。但是需要注意的是，这种解法有一个错误假设，就是最小值在最左边，而最大值在最右边，很明显这个假设是错误的，最大值和最小值可能出现在同一侧。
```go
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
```
