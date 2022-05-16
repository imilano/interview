package leetcode

import (
	"math"
)

/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*/
// TODO brush it up
// 实际上不用处理负号
func reverse(x int) int {
	var res int
	for x != 0 {
		// 这里不能用 Math.MaxInt， 因为 int 有多少位是随着平台变化的，并不能固定的认为它是 32 位
		// int 在 32 位机器上是 32 位的，但在 64 位机器上却是 64 位的
		if abs(res) > math.MaxInt32/10 {
			return 0
		}

		res = res*10 + x%10
		x /= 10
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func reverse_str(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func Reverse(x int) int {
	return reverse(x)
}
