package leetcode

import "math"

/*
Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).

The algorithm for myAtoi(string s) is as follows:

- Read in and ignore any leading whitespace.
- Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
- Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
- Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
- If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
Return the integer as the final result.

Note:

Only the space character ' ' is considered a whitespace character.
Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.
*/

//TODO complete this
func myAtoi(s string) int {
	// ignore any leading whitespace
	if len(s) == 0 {
		return 0
	}
	var index int
	for idx, value := range s {
		index = idx
		if value != ' ' {
			break
		}
	}
	s = s[index:]

	// if it start with negative sign
	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	if len(s) == 0 && (s[0] == '+' || s[0] == '-') {
		return 0
	}
	// eliminate leading zero if any
	idx := 0
	for idx, _ = range s {
		if s[idx] != '0' {
			break
		}
	}

	s = s[idx:]
	size := len(s)
	if size == 0 {
		return 0
	}

	idx = 0
	for i := 0; i < size; i++ {
		if s[i] < '0' || s[i] > '9' {
			idx = i
			break
		}

		idx = i
	}

	if idx+1 != size {
		s = s[:idx]
	}

	var res int
	size = len(s)
	for i := 0; i < size; i++ {
		res *= 10
		res += int(s[i] - '0')
	}

	if negative {
		res = -res
		if res < math.MinInt32 {
			res = math.MinInt32
		}
	} else if res > math.MaxInt32 {
		res = math.MaxInt32
	}

	return res
}

// for test
func MyAtoi(s string) int {
	return myAtoi(s)
}
