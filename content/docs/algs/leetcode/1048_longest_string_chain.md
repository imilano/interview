---
title: 1048. Longest String Chain
weight: 10
tags: [
	"DP",
	"String"
]
---
## Description
> You are given an array of `words` where each word consists of lowercase English letters.
> 
> `wordA` is a predecessor of `wordB` if and only if we can insert exactly one letter anywhere in `wordA` without changing the order of the other characters to make it equal to `wordB`.
> 
> - For example, `"abc"` is a predecessor of `"abac"`, while `"cba"` is not a predecessor of `"bcad"`.
> A word chain is a sequence of words `[word1, word2, ..., wordk]` with `k >= 1`, where `word1` is a predecessor of `word2`, `word2` is a predecessor of `word3`, and so on. A single word is trivially a word chain with `k == 1`.
> 
> Return the length of the longest possible word chain with words chosen from the given list of `words`.

## Solutions
### DP
题主这里首先想到的 DP 的解法，写出解法之后，一度以为会超时，结果提交之后竟然 ac 了。首先需要明白的是，题目并不要求你选的字符串之间需要连续，而且我们也可以看出来，对字符串进行排序之后其实更方便我们进行处理。于是首先我们先对数组进行排序：按照长度进行排序，长度相同的按照字典序进行排序。然后我们定义 `dp[i]`表示数组中 `[0,i]` 之间的最长 `predecessor` 长度，然后我们来一个双重循环，如果 `j` 是 `i` 的 `predecessor` 的话，就进行状态转移，转移方程为 `dp[i] = max(dp[i], dp[j]+1)`。因为最长 `predecessor` 长度可能并不一定会在最后一个字符串处取到，所以我们还需要一个 `res` 来记录我们当前计算到的最长 `predecessor` 长度。最后放回 `res` 即可.

```go
func longestStrChain(words []string) int {
    // 先排序，方便比较：首先根据长度进行排序，如果长度相同，则按照字典序排序
    size := len(words)
    sort.Slice(words, func(i,j int) bool {
        m, n := len(words[i]), len(words[j])
        if m == n {
            return words[i] < words[j]
        }
        
        return m < n
    })
    
    // dp[i] 表示 [0,i] 内的最长 predecessor 个数，初始化为 1
    dp := make([]int, size)
    for i := 0; i < size; i++ {
        dp[i] = 1
    }
    
    res := 1
    for i := 0; i < size; i++ {
        for j := 0; j < i; j++ {
            // 如果 j 是 i 的 predecessor， 那么更新 dp[i]，同时更新结果 res
            if isPredecessor(words[j], words[i]) {
                dp[i] = max(dp[i], dp[j]+1)   
                res = max(res, dp[i])
            }
        }
    }
    
    return res
}

// 判断两个字符是不是 predecessor 关系
func isPredecessor(s1,s2 string) bool {
    // s1 总是长度比较小的字符串
    m, n := len(s1), len(s2)
    if m > n {
        s1, s2 = s2, s1
        m,n = n, m
    }
    
    // 如果二者长度差不是 1，那么肯定没有 predecessor 关系
    if n != m + 1 {
        return false
    }
    
    // 遍历长字符串s2，对于位置 i，判断 i 前面的字符和 i 后面的字符合起来的字符是不是与 s1 相等
    for i := 0; i < n; i++ {
        // 对于第一个位置需要进行特殊处理
        if i == 0 && i + 1 < n && s2[i+1:] == s1 {
            return true
        }
        
        // 对于最后一个位置需要特殊处理
        if i == n-1 && s2[:i] == s1 {
            return true
        }
        
        // i 前面的字符串和 i 后面的字符串合起来是不是等于 s1
        // 如果是， 怎说明二者是 predecessor 关系
        if s2[:i] + s2[i+1:] == s1 {
            return true
        }
    }
    
    // 没有找到 predecessor 关系，返回 false
    return false
}

func max(a,b int) int {
    if a < b {
        return b
    }
    
    return a
}
```
