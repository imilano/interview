package leetcode

/*
Given a string s, find the longest palindromic subsequence's length in s.

A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.
*/

// 子序列，意味着不需要连续，比如对于 bbbab，最长回文子序列是bbbb，而不是bbb。
// dp[i][j] 表示字符s从i到j（闭区间）之间的最大回文串，则如果s[i]=s[j]，那么有 dp[i][j] = dp[i+1][j-1] + 2,否则， dp[i][j] = max(dp[i+1][j], dp[i][j-1])
// 另外注意，回文串是加2不是加1；此外，注意遍历的方向，求解dp[i][j]时， dp[i+1][j-1]的子问题需要完全求解才行
func longestPalindromeSubseq(s string) int {
	size := len(s)
	dp := make([][]int, size)
	for index, _ := range dp {
		dp[index] = make([]int, size)
	}

	// 最小回文长度至少是1，所以二维数组对角线的每一个位置都可以预先填充一个1
	for i := 0; i < size; i++ {
		dp[i][i] = 1
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				// dp[j][i] = dp[i-1][j+1] + 1
				dp[j][i] = dp[j+1][i-1] + 2
			} else {
				// dp[j][i] = max(dp[i-1][j], dp[i][j+1])
				dp[j][i] = max(dp[j+1][i], dp[j][i-1])
			}
		}
	}

	return dp[0][size-1]
}

func LongestPalindromeSubseq(s string) int {
	return longestPalindromeSubseq(s)
}

// 解法2，转化为求 s 和 s 的逆序的最长公共子序列问题
func longestPlaindromeSubseqUsingLCS(s string) int {
	rs := reverseStr(s)
	return longestCommonSubsequence(s, rs)
}
