---
title: 0049. Group Anagrams
weight: 10
---

## Description

> Given an array of strings strs, group the anagrams together. You can return the answer in any order.
> 
> An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Solutions

简单的方法就是，对具有相同字母以及字母出现次数也相同的字符串进行排序时，其排序结果都是一样的。所以可以创建一个 map，其中 key 为字符串的排序，value 为具有这样一个排序结果的字符串数组。
```go
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
```
