---
title: 0003. Longest Substring Without Repeating Characters
weight: 10
tags: [
	"String",
	"Hash Table",
	"Sliding Window"
]
---
## Description
> Given a string s, find the length of the longest substring without repeating characters.

## Solutions
### Hash Table
中间扩散法，时间复杂度会比较高，因为会有很多的重复扫描。核心思想是，遍历一次字符串，然后对于每个遍历到的位置，从中心向两侧进行扫描，扫描时使用一个 map 来记录那些已经出现过的字符；在一侧扫描过程中，如果发现该字符已经出现过，则停止扫描该侧。最后 map 的长度即为以该字符为中心的不重复字符的长度。
```go
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
```


### Sliding Window & Hash Table
很明显也可以使用滑动窗口配合 Hash Table 的方式来解题。首先固定滑动窗口的左边界，然后将滑动窗口的右边界向右滑动，每滑动一个位置，就检查该位置的字符是否有在当前滑动窗口中出现过(也就是说，既要检查该字符是否在 Hash Table 中出现过，还要检查最近一次出现的下标是否位于滑动窗口之内)，如果出现过，则更新滑动串口左边界。另外，每次遍历都需要更新当前滑动窗口的最大长度。
```go
func lengthOfLongestSubstring(s string) int {
	// 用于存储字符以及该字符最近出现的下标位置
    dict := make(map[byte]int)
	// pre 表示滑动窗口左边界，其指向有效窗口的前一个位置
    res, pre, size := 0, -1, len(s)
    for i := 0; i < size; i++ {
		// 更新窗口不能放在这里。考虑「i 的出现引入了重复字符」以及「i 的出现没有引入重复字符」这两种情况下，
		// 两种情况下窗口的更新策略应该是不同的，这么写会导致二者共用一个更新策略，那很明显是错误的。
		// res = max(res, i - pre)
		// idx > pre 保证了出现的字符一定是位于滑动窗口中的，而不是那些之前已经出现过的不在滑动窗口中的字符
        if idx, ok := dict[s[i]]; ok && idx > pre {
			// 更新窗口不能放在这里。考虑当整个字符串都没有重复字符串的情况下，放在这里会导致窗口值永远不会更新
			// res = max(res, i - pre)
			// 更新滑动窗口左边界
            pre = idx
        }
		// 更新最大窗口值。因为 pre 始终指向的左边界的前一个位置，所以使用 i - pre 即可
        res = max(res, i - pre)
		// 更新当前字符的最近出现位置
        dict[s[i]] = i
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
