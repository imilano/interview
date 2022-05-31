---
title: 0022. Generate Parentheses
weight: 10
tags: [
	"Recursive",
	"Branch Pruning"
]
---
## Description

> Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

## Solutions

### Recursive

参见： https://leetcode-cn.com/problems/generate-parentheses/solution/hui-su-suan-fa-by-liweiwei1419/

使用两个变量 left 和 right 来表示「左括号使用了集合」和「右括号使用了几个」，通过分析可以得到以下结论：
	- 只有当left 和 right 都比 n 小的时候，才产生分支。
	- 产生左分支的时候，只看当前还有左括号可用
	- 产生右分支的时候，还受到左括号数量的限制，left 要大于 right 的时候，才可以产生分支。否则就是无效分支，可以直接返回。
	- 在左边和右边的括号数都等于 n 的时候结算。

```go
func generateParenthesis(n int) []string {
    var res []string
	if n == 0 {
		return res
	}

	generateParenthesisHelper(n, 0,0,"", &res)
	return res
}

func generateParenthesisHelper(n, left, right int, cur string, res *[]string) {
	if left == n && right == n {
		*res = append(*res, cur)
		return
	}

	// 剪枝
	if left < right {
		return
	}

	if left < n {
		generateParenthesisHelper(n, left+1, right, cur + "(", res)
	}
	if right < n {
		generateParenthesisHelper(n, left, right+1, cur + ")", res)
	}
}
```
