package leetcode

/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/

func wordBreak(s string, wordDict []string) bool {
	return wordBreakSolution1(s, wordDict)
}

// 动态规划。假设dp[i]表示字符串s从0到i（不包括）的子串是否可以被拆分，那么我们就可以使用两个循环，外层循环i，内层循环j，其中内层循环j把0
// 到i的子串分为了 0到j和j到i两个字串。每次遍历时，dp[j]我们可以直接查表得出，那么只需要再看 s[j:i]是否在给出的wordDict中即可，如果在，说明0到i是可以拆分的，那么结束此次遍历，继续求解dp[i+1]
// 最后所求即为 dp[n]。注意我们的dp定义，从定义看，我们需要申请 n+1个长度的数组，并且很容易知道，dp[0]=true
func wordBreakSolution2(s string, wordDict []string) bool {
	size := len(s)
	dp := make([]bool, size+1)

	// 这个初始化条件很重要
	dp[0] = true
	m := make(map[string]bool, len(wordDict))
	for _, s := range wordDict {
		m[s] = true
	}

	// 注意子问题的求解方向
	for i := 0; i <= size; i++ {
		for j := 0; j < i; j++ {
			_, exist := m[s[j:i]]
			if dp[j] && exist {
				dp[i] = true
				// 只需要确定可解就行，但是我们并不关注有几种可解的方案
				break
			}
		}
	}

	return dp[size]
}

// 暴力遍历很容易超时，优化就是可以加上一个记忆数组
func wordBreakSolution1(s string, wordDict []string) bool {
	size := len(wordDict)
	m := make(map[string]bool, size)
	memo := make(map[int]bool, size)
	for _, s := range wordDict {
		m[s] = true
	}

	return recursiveWordBreak(0, s, m, &memo)
}

func recursiveWordBreak(start int, s string, m map[string]bool, memo *map[int]bool) bool {
	size := len(s)
	if start >= size {
		return true
	}

	dividable, ok := (*memo)[start]
	if ok {
		return dividable
	}

	for i := start + 1; i <= size; i++ {
		_, ok := m[s[start:i]]
		if ok && recursiveWordBreak(i, s, m, memo) {
			(*memo)[start] = true
			return true
		}
	}

	// TODO 这个地方需要特别注意下，当你 s[:start] 不可分时，也需要做标记
	(*memo)[start] = false
	return false
}

// for test
func WordBreak(s string, wordDict []string) bool {
	return wordBreak(s, wordDict)
}
