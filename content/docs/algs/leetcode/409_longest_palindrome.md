---
title: 0409. Longest Palindrome
weight: 10
tags: [
	"Palindrome",
	"String",
	"Hash Table"
]
---
## Description

> Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.
> 
> Letters are case sensitive, for example, "Aa" is not considered a palindrome here.

## Solutions
### Hash Table
回文串有两种形式，一种是`xxyy`，另一种是`xxyzz`，也就是说，两边的字符需要出现偶数次，那么这里就可以转换为统计出现偶数次的字符的数量。在统计的过程中，如果发现有一个字符出现了奇数次，那么这个奇数次的字符就可以放在回文串的中心，那么就需要记录一下，最后在结果上加上 1.
```go
func longestPalindrome(s string) int {
	// 统计每个字符出现的次数
    size := len(s)
    dict := make(map[byte]int)
    for i := 0; i < size; i++ {
        dict[s[i]]++
    }
    
    var res int
    var addMid bool
    for _, num := range dict {
        res += num
		// 如果出现了奇数次，那么只取最大的偶数，也就是要减去 1
		// 如果只是出现了一次的话，那么这里相当于没有加
        if num % 2 == 1 {
            res -= 1

			// 回文串有两种类型，一种是中间一个字符然后两边对称的形式，
			// 也就是说，如果有一个字符出现了奇数次，那么最后可以考虑把这个字符加上。
            addMid = true
        }
    }
    
    if addMid {
        res += 1
    }
    return res
}
```
