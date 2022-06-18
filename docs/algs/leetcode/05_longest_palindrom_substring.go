package leetcode

/*
Given a string s, return the longest palindromic substring in s.
*/

// 方法 1： 中心扩散
func longestPalindromeSubstring(s string) string {
	size := len(s)
	if size <= 1 {
		return s
	}

	var res string
	for i := 0; i < size; i++ {
		r1 := longestPalindromeHelper(s, i, i, size)
		r2 := longestPalindromeHelper(s, i-1, i, size)
		res = getMaxString(r1, r2, res)
	}

	return res
}

func getMaxString(a, b, s string) string {
	if len(a) < len(b) {
		if len(b) < len(s) {
			return s
		} else {
			return b
		}
	} else {
		if len(a) < len(s) {
			return s
		} else {
			return a
		}
	}
}
func longestPalindromeHelper(s string, left, right, size int) string {
	var res string
	for left >= 0 && right < size {
		if s[left] == s[right] {
			res = s[left : right+1]
			left--
			right++
		} else {
			break
		}
	}

	return res
}

// 方法 2，如果将这个字符串逆转，那么求最长回文串就会变成求两个字符串的最长公共子串
// 方法 3： 动态规划
// 使用一个二维数组 dp[i][j]表示字符串 i 到 j 是否为回文串，则有：
// 	dp[i][j] = true if i == j
// 	dp[i][j] = s[i] == s[j]  if j-i == 1
// 	dp[i][j] = dp[i+1][j-1] && s[i] == s[j] if j - i > 1
func longestPalindromeSolution3(s string) string {
	size := len(s)
	if size <= 1 {
		return s
	}

	dp := make([][]bool, size)
	for idx, _ := range dp {
		dp[idx] = make([]bool, size)
	}

	left, step := 0, 1
	for i := 0; i < size; i++ {
		dp[i][i] = true
		for j := 0; j < i; j++ {
			dp[j][i] = s[i] == s[j] && (i-j == 1 || dp[j+1][i-1])
			if dp[j][i] && step < i-j+1 {
				step = i - j + 1
				left = j
			}
		}
	}

	return s[left : left+step]
}

// 方法 3： 马拉车算法

// for test
func LongestPalindromeSubstring(s string) string {
	return longestPalindromeSubstring(s)
}
