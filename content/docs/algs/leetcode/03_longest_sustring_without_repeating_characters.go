package leetcode

/*
Given a string s, find the length of the longest substring without repeating characters.
*/

func lengthOfLongestSubstring11(s string) int {
	return lengthOfLongestSubstringSolution2(s)
}

// TODO 这里需要好好注意一下
// 滑动窗口法
func lengthOfLongestSubstringWithoutRepeatingSolution2(s string) int {
	var res int
	m := make(map[byte]int)
	left := -1
	for index, _ := range s {
		// 为什么要要 m[s[index]] > left ？
		// 因为一些旧字符虽然出现在了 map 中，但是并不一定出现在我们维持的滑动窗口中.
		// 而只有在滑动窗口中出现了重复字符，才需要移动左指针。
		if _, ok := m[s[index]]; ok && m[s[index]] > left {
			left = m[s[index]]
		}

		m[s[index]] = index
		// left 指针指向的是滑动窗口左边界的前一个位置，所以不要在 index - left 的结果上再加 1
		res = max(res, index-left)
	}
	return res
}

// 中心遍历法，也就是暴力算法
func lengthOfLongestSubstringWithoutRepeatingSolution1(s string) int {
	// handle empty string
	var res int
	size := len(s)
	if size <= 1 {
		return size
	}

	// 中心扩散法
	for i := 0; i < size; i++ {
		res = max(res, expandFromMiddle(s, i))
	}

	return res
}

func expandFromMiddle(s string, start int) int {
	left, right := start-1, start+1
	size := len(s)
	m := make(map[byte]bool)
	m[s[start]] = true
	for left > 0 {
		if _, ok := m[s[left]]; ok {
			left += 1
			break
		}
		m[s[left]] = true
		left -= 1
	}

	for right < size {
		if _, ok := m[s[right]]; ok {
			right -= 1
			break
		}

		m[s[right]] = true
		right += 1
	}

	return len(m)
}
