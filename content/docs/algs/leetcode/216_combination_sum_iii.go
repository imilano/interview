package leetcode

/*
Find all valid combinations of k numbers that sum up to n such that the following conditions are true:

Only numbers 1 through 9 are used.
Each number is used at most once.
Return a list of all possible valid combinations. The list must not contain the same combination twice, and the combinations may be returned in any order.
*/

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var out []int
	if n <= 0 {
		return res
	}

	recursiveCombinationSum3(k, n, 1, out, &res)
	return res
}

func recursiveCombinationSum3(k int, n int, start int, out []int, res *[][]int) {
	if n < 0 {
		return
	}

	if n == 0 && len(out) == k {
		tmp := make([]int, len(out))
		copy(tmp, out)
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= 9 && i <= n; i++ {
		out = append(out, i)
		recursiveCombinationSum3(k, n-i, i+1, out, res)
		out = out[:len(out)-1]
	}
}

// for test
func CombinationSum3(k int, n int) [][]int {
	return combinationSum3(k, n)
}
