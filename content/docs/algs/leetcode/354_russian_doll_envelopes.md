---
title: 0354. Russian Doll Envelopes
weight: 10
tags: [
	"DP"
]
---
## Description
> You are given a 2D array of integers envelopes where envelopes[i] = [wi, hi] represents the width and the height of an envelope.
> 
> One envelope can fit into another if and only if both the width and height of one envelope are greater than the other envelope's width and height.
> 
> Return the maximum number of envelopes you can Russian doll (i.e., put one inside the other).
> 
> Note: You cannot rotate an envelope.

## Solutions
这题的直白翻译就是“俄罗斯套娃，问最多能套几个”。题主想出了一个{{< katex >}} \Omicron(n^2) {{< /katex >}} 的DP 解法，奈何时间复杂度太高，死在最后几个测试用例上了。基本思路就是，先给信封根据宽度从小到大排一下序，然后创建一个 dp 数组，因为最少也能够装一个，所以 dp 数组的每个位置都先初始化为 1；然后从下标 1 开始遍历信封数组，对于每个 envelopes[i]，找出其当前能够装的最多的信封数，然后继续往后遍历。怎么找呢，对于每个信封 i，又新开一个从 0 到 i-1 的循环 j，然后比较信封 i 和信封 j 的宽度和长度，如果 i 能够装得下 j，那么就更新{{< katex >}} dp[i] = max(dp[i], dp[i-j]+1) {{< /katex >}}，这里需要注意遍历的方向，一定是需要先从 0 到 i-1 的。先贴一下代码：
```go
func maxEnvelopes(envelopes [][]int) int {
    //corner case
    size := len(envelopes)
    if size <= 1 {
        return size
    }
    
    res := 1
    sort.Slice(envelopes, func(i,j int) bool {
        return envelopes[i][0] < envelopes[j][0]
    })
    
    dp := make([]int, size)
    for i := 0; i < size; i++ {
        dp[i] = 1
    }
    
    for i := 1; i < size; i++ {
        for j := i; j > 0; j-- {
            if envelopes[i][0] > envelopes[i-j][0] && envelopes[i][1] > envelopes[i-j][1] {
                dp[i] = max(dp[i], dp[i-j] + 1)
                res = max(res, dp[i])
            }
        }
    }
    
    return res
}


func max(a,b int) int {
    if a < b {
        return b
    }
    
    
    return a
}
```
