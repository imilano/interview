package leetcode

import (
	"sort"
)

/*
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.
*/

// 需要一个set或者map用来去重
// var m map[string]bool

// 因为已经对数组进行了排序，所以不需要单独再使用 map 来进行去重
func combinationSum2(candidates []int, target int) [][]int {
	// 由于 leetcode 编译器的问题，每次执行前都需要初始化环境变量，对上一次执行的全局变量清零处理，否则会出现错误
	// m = make(map[string]bool)

	var res [][]int
	size := len(candidates)
	sort.Ints(candidates)
	var arr []int

	recursiveCombinationSum2(candidates, target, 0, size, arr, &res)
	return res
}

func recursiveCombinationSum2(candidates []int, target int, start int, size int, arr []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		// if !inside(arr) {
		// 	// m[toString(arr)] = true
		// 	tmp := make([]int, len(arr))
		// 	copy(tmp, arr)
		// 	*res = append(*res, tmp)
		// }
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*res = append(*res, tmp)
		return
	}

	if start >= size {
		return
	}
	for i := start; i < size; i++ {
		// if candidates[i] <= target {
		// 	if i > start && candidates[i] == candidates[i - 1] {
		// 		continue
		// 	}
		// 	arr = append(arr, candidates[i])
		// 	recursiveCombinationSum2(candidates, target-candidates[i], i+1, size, arr, res)
		// 	arr = arr[:len(arr)-1]
		// }

		// TODO
		// 加上这一句，可以避免数组中出现多个连续重复项时出现超时的情况
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		arr = append(arr, candidates[i])
		recursiveCombinationSum2(candidates, target-candidates[i], i+1, size, arr, res)
		arr = arr[:len(arr)-1]
	}
}

// func inside(arr []int) bool {
// 	s := toString(arr)
// 	if _, ok := m[s]; ok {
// 		return true
// 	}

// 	return false
// }

// func toString(arr []int) string {
// 	var s string
// 	for _, i := range arr {
// 		s += strconv.Itoa(i)
// 	}

// 	return s
// }

// for test
func CombinationSum2(candidates []int, target int) [][]int {
	return combinationSum2(candidates, target)
}
