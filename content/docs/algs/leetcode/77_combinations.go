package leetcode

/*
Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].

You may return the answer in any order.
*/

func combine(n int, k int) [][]int {
	var res [][]int
	if k == 0 {
		return res
	}

	for i := 1; i <= n; i++ {
		arr := []int{i}
		combineSolution1(i+1, k-1, n, arr, &res)
	}

	return res
}

func combineSolution1(start int, k int, n int, arr []int, res *[][]int) {
	if k < 0 {
		return
	}

	if k == 0 {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= n; i++ {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		tmp = append(tmp, i)
		combineSolution1(i+1, k-1, n, tmp, res)
	}
}

// for test
func Combine(n int, k int) [][]int {
	return combine(n, k)
}
