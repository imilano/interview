package leetcode

/*
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.
*/

// 动态规划。定义 dp[i][j] 表示字符串 text1 的前 i 个字符与 text2 的前 j 个字符的最长公共子序列。那么如何更新 dp[i][j] 呢？
// 如果 text1[i-1] == text2[j-1] 相等，那么说明最长公共子序列在 dp[i-1][j-1] 的基础上又增加了一个，此时 dp[i][j] = dp[i-1][j-1] + 1.
// 如果 text1[i] != text2[j]， 那么还可以进行错位比较， dp[i][j] = max(dp[i-1][j], dp[i][j-1])。
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for index, _ := range dp {
		dp[index] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[m][n]
}
