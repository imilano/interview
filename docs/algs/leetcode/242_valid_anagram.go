package leetcode

import (
	"reflect"
	"sort"
)

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
*/

func isAnagram(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n {
		return false
	}

	dict1, dict2 := make([]int, m), make([]int, m)
	for i, _ := range s {
		dict1[i] = int(s[i])
		dict2[i] = int(t[i])
	}

	sort.Ints(dict1)
	sort.Ints(dict2)
	if !reflect.DeepEqual(dict1, dict2) {
		return false
	}

	return true
}
