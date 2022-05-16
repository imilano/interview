package leetcode



/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
*/

// build map
var dict map[byte][]string = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var res []string
	size := len(digits)
	if size == 0 {
		return res
	}

	letterCombinationsHelper(digits,size, 0, "", &res)
	return res
}

func letterCombinationsHelper(digits string, size int, start int, cur string, res *[]string) {
	if start > size {
		return
	}

	if start == size {
		*res = append(*res, cur)
		return
	}

	strs := dict[digits[start]]
	for _, str := range strs {
		letterCombinationsHelper(digits, size, start+1, cur+str, res)
	}
}
