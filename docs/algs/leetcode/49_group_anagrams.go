package leetcode

import (
	"sort"
	"strings"
)

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
*/

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)
	for _, str := range strs {
		str_splited := strings.Split(str, "")
		sort.Strings(str_splited)
		s := strings.Join(str_splited, "")
		m[s] = append(m[s], str)
	}

	for _, strs := range m {
		res = append(res, strs)
	}

	return res
}

// 解法 2，不用排序
// 用一个大小为 26 的int 数组来统计每个单词中字符出现的次数，然后将 int 数组转换为唯一一个的字符串，跟字符串数组进行映射，这样就不用给字符串排序了。
