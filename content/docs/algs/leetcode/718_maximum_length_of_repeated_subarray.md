---
title: 0718. Maximum Length of Repeated Subarray
weight: 10
tags: [
	"String",
	"Longest Common Substring"
]
---
## Description
> Given two integer arrays nums1 and nums2, return the maximum length of a subarray that appears in both arrays.

## Solutions
### Dynamic Programming
定义 `dp[i][j]`表示 `nums1` 中以 `i` 为结尾的子数组和 `nums2` 中以 `j` 为结尾的子数组当前的匹配的最长公共子数组的长度。如果 `nums1[i] == num2[j]`，说明当前两个数字相同，那么当前最长公共子数组的长度就是 `nums1` 以 `i-1` 为结尾的子数组和 `nums2` 以 `j-1` 为结尾的子数组的最长公共子数组的长度加 `1`，即 `dp[i][j] = dp[i-1][j-1] + 1`；否则说明二者当前无公共子子数组或者公共子数组在当前位置不连续，则`dp[i][j] = 0`。
```go
func findLength(nums1 []int, nums2 []int) int {
    m,n := len(nums1), len(nums2)
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int, n + 1)
    }
    
    var res int
    for i := 1; i <= m;i++ {
        for j := 1; j <= n; j++ {
            if nums1[i-1] == nums2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
                res = max(res, dp[i][j])
            } else {
                dp[i][j] = 0
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
