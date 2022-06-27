---
title: 0424. Longest Repeating Character Replacement
weight: 10
tags: [
	"Two Pointer",
	"Sliding Window"
]
---

## Description
> You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.
> 
> Return the length of the longest substring containing the same letter you can get after performing the above operations.

## Solutions
### Sliding Window
这个题题主又不会了，在网上看到可以使用滑动窗口来解。假设没有 k 的限制，也就是说可以替换任意多次，把一个字符串变为只有只有一个字符重复的字符串需要的最少置换次数是多少呢？其实就是扫描一遍字符串，记录出现次数最多的字符串的出现个数k，然后用总长度减去 k ，所得值即为最少置换次数。

而如果带上了 k，那么就要满足: 子字符串的长度减去该子字符串中出现次数最多的字符的长度要小于等于 k。具体做法：
- 使用一个变量 start 记录滑动窗口左边界，然后遍历字符串，记录每个单词的出现次数。
- 然后判断是否满足上面所说的那个条件，如果不满足，则前移 start，同时还需要注意减去该字符的出现次数，直到满足上面所说的条件，然后更新 res 即可。
需要注意的是，当滑动窗口的左边界向右移动了后，并不需要更新 maxCnt 的值。原因在于既然我们求的是最大值，那么既然此次已经出现了 maxCnt，那么后续我们只会希望出现比 maxCnt 更大的值，所以更新 maxCnt 没有意义。maxCnt 表示子字符串中出现次数最多的字符的出现个数。

```go
func characterReplacement(s string, k int) int {
	var res, start, maxCnt int
	size := len(s)
	m := make(map[byte]int, 26)
	for i := 0; i < size; i++ {
		// 记录每个字符的出现次数
		m[s[i]] += 1
		// 更新子串中出现次数最多的字符的出现个数
		maxCnt = max(maxCnt, m[s[i]])
		for i-start+1-maxCnt > k { // 子字符串的长度减去该子字符串中出现次数最多的字符的长度要小于等于 k
			m[s[start]] -= 1
			start += 1
		}

		res = max(res, i-start+1)
	}

	return res
}
```
