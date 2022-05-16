package leetcode

/*
A message containing letters from A-Z can be encoded into numbers using the following mapping:

'A' -> "1"
'B' -> "2"
...
'Z' -> "26"
To decode an encoded message, all the digits must be grouped then mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, "11106" can be mapped into:

"AAJF" with the grouping (1 1 10 6)
"KJF" with the grouping (11 10 6)
Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".

Given a string s containing only digits, return the number of ways to decode it.

The test cases are generated so that the answer fits in a 32-bit integer.
*/

// TODO fix me
func numDecodings(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	var res int
	numDecodingsRecur(s, &res)
	return res
}

func numDecodingsRecur(s string, res *int) {
	size := len(s)
	if size == 0 {
		*res += 1
	}

	if size == 1 && s != "0" {
		*res += 1
	}

	if size >= 2 {
		if s[:1] == "0" {
			return
		}
		if s[:2] >= "10" && s[:2] <= "26" {
			numDecodingsRecur(s[2:], res)
		}

		numDecodingsRecur(s[1:], res)
	}
}

// for test
func NumDecodings(s string) int {
	return numDecodings(s)
}
