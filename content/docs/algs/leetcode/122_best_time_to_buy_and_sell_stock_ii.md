---
title: 0122. Best Time to But And Sell Stock II
weight: 10
---

## Description
> You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
> 
> On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.
> 
> Find and return the maximum profit you can achieve.

## Solutions
如果要收益最大，那么最好是低价买入，高价卖出。所以只需要检查每一支股票，只要当前股票比前一天的大，那么就可以获得收益。
```go
  func maxProfit(prices []int) int {
    var res int
    size := len(prices)
    if size <= 1 {
        return res
    }
    
    for i := 1; i < size; i++ {
        if prices[i] > prices[i-1] {
            res += prices[i]- prices[i-1]
        }
    }
    return res
}

```
