package leetcode

import "sort"
/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/

// 主要还是要排序，这样能够减少很多处理
func longestCommonPrefix(strs []string) string {
    size := len(strs)
    if size == 0 {
        return ""
    }
    if size == 1 {
        return strs[0]
    }
    
    sort.Strings(strs)
    longestLen := len(strs[0])
    if longestLen == 0 {
        return ""
    }
    
    
    for i := 1; i <= longestLen; i++ {
        for j := 1; j < size; j++ {
            if strs[j][:i] != strs[j-1][:i] {
                return strs[j][:i-1]
            }
        }
    }
    
    return strs[0]
}

// 给数组进行排序，这样的话，只需要比较排序后的第一个字符串和最后一个字符串即可，因为这两个字符串的区别是最大的，如果存在公共前缀的话，那么公共前缀必然也是这两个字符的公共前缀
func longestCommonPrefix2(strs []string) string {
	var res string
	size := len(strs)
	if size == 0 {
		return res
	}
	sort.Strings(strs)
	first, last := strs[0], strs[size-1]
	i, minLen := 0, min(len(first), len(last))
	for i < minLen && first[i] == last[i] {
		i++
	}

	return first[:i]
}
