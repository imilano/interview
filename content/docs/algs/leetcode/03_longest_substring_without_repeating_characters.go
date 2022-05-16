package leetcode

/*
Given a string s, find the length of the longest substring without repeating characters.
*/

func lengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstringSolution2(s)
}

// 滑动窗口方法
func lengthOfLongestSubstringSolution2(s string) int {
	// handle empty string
	if len(s) <= 1 {
		return len(s)
	}

	var res int
	// left = -1 能处理字符从 0 开始计数的情况，这样在计算 count 的时候能正确计算
	left := -1
	dict := make(map[rune]int)
	for index, char := range s {
		i, ok := dict[char]
		if ok && i > left {
			left = i
		}

		dict[char] = index
		res = max(res, index-left)

	}

	return res
}

// 对于每个位置的字符，从中间向两侧扫描；使用一个 map 记录已经出现过的字符。在一侧扫描过程中，如果该字符已经出现过，则停止扫描该侧。
// 最后 map 的长度即为以该字符为中心的不重复字符的长度
func lengthOfLongestSubstringSolution1(s string) int {
	// handle empty string
	var res int
	if len(s) <= 1 {
		return len(s)
	}

	for i := 0; i < len(s); i++ {
		res = max(res, helper(s, i))
	}

	return res
}

func helper(s string, mid int) int {
	left, right := mid-1, mid+1
	dict := make(map[byte]bool)
	dict[s[mid]] = true
	for left > 0 {
		_, ok := dict[s[left]]
		if ok {
			break
		}

		dict[s[left]] = true
		left--
	}

	for right < len(s) {
		_, ok := dict[s[right]]
		if ok {
			break
		}
		dict[s[right]] = true
		right++
	}

	return len(dict)
}
