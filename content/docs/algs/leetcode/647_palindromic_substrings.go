package leetcode

/*
Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.
*/

// 中心扩散法
func countSubstrings(s string) int {
	size := len(s)
	var res int
	for i := 0; i < size; i++ {
		countSubstringsHelper(s, size, i, i, &res)
		countSubstringsHelper(s, size, i-1, i, &res)
	}

	return res
}

func countSubstringsHelper(s string, size, left, right int, res *int) {
	for left >= 0 && right < size && s[left] == s[right] {
		left--
		right++
		*res += 1
	}
}
