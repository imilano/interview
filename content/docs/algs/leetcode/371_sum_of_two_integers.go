package leetcode

import "math"

/*
Given two integers a and b, return the sum of the two integers without using the operators + and -.
*/

/*
 * 使用二进制运算来进行。对于二进制异或而言， 如果不考虑进位： 0 + 0 = 0， 0 + 1 = 1， 1 + 0 = 1， 1 + 1 = 0，而这刚好与二进制的异或相对应，因此可以用异或操作来模拟二进制相加每一位的和。
 * 而如果只考虑进位： 1 + 0 = 0， 1 + 1 = 1， 0 + 1 = 0， 0 + 0 = 0， 而这刚好与与运算对应，因而可以用来模拟二进制相加每一位的进位。
 * 将每一位的和与对应的进位相加即是结果。需要注意的是，进位是加在更高位的，所以进位结果需要左移一位再与和相加。而每一次相加又会产生进位，那么什么时候计算停止呢？当进位为 0 的时候停止即可。
 */
func getSum(a int, b int) int {
	if b == 0 {
		return a
	}

	sum := a ^ b
	carry := (a & b) << 1
	return getSum(sum, carry)
}

// 取巧做法（：
func getSumSolution2(a int, b int) int {
	return int(math.Log2(math.Exp2(float64(a)) * math.Exp2(float64(b))))
}
