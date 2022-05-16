package leetcode

/*
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.
*/

func countBits(n int) []int {
	return countBitsSolution2(n)
}

// 方法 1，很容易想到
func solution1(n int) []int {
	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = countOne(i)
	}

	return arr
}

func countOne(n int) int {
	var res int
	for n != 0 {
		if n&1 == 1 {
			res++
		}

		n >>= 1
	}

	return res
}

// 方法 2，规律： 从 1 开始，遇到偶数时，其 1 的个数和该偶数除以 2 得到的数字的1 的个数相同；遇到奇数时，其 1 的个数与该奇数除以 2 得到的数字的 1 的个数再加 1 相同。
func countBitsSolution2(n int) []int {
	res := []int{0}
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			res = append(res, res[i/2])
		} else {
			res = append(res, res[i/2]+1)
		}
	}

	return res
}
