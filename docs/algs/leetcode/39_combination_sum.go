package leetcode

/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
*/

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	size := len(candidates)
	for i := 0; i < size; i++ {
		if candidates[i] <= target {
			var arr []int
			arr = append(arr, candidates[i])
			recursiveCombinationSum(candidates, target-candidates[i], i, arr, &res)
		}
	}

	return res
}

func recursiveCombinationSum(candidates []int, target int, start int, arr []int, res *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		*res = append(*res, arr)
		return
	}

	size := len(candidates)
	for i := start; i < size; i++ {
		if candidates[i] <= target {
			newArr := make([]int, len(arr))
			copy(newArr, arr)
			newArr = append(newArr, candidates[i])
			recursiveCombinationSum(candidates, target-candidates[i], i, newArr, res)
		}
	}
}

// for test
func CombinationSum(candidates []int, target int) [][]int {
	return combinationSum(candidates, target)
}

// 方法 2，也是回溯。与前一个不一样的是，这里的解法是一次性同时添加多个同一元素，这样能够提升一些性能
func combinationSumSolution2(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	combinationSumHelper(candidates, 0, target, cur, &res)
	return res
}

func combinationSumHelper(candidates []int, start int, target int, cur []int, res *[][]int) {
	size := len(candidates)
	if target == 0 {
		// 基于切片的性质，这里需要添加其拷贝
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	if start >= size {
		return
	}

	for i := start; i < size; i++ {
		for j := target / candidates[i]; j > 0; j-- {
			//一次性添加多个同一元素
			for k := 0; k < j; k++ {
				cur = append(cur, candidates[i])
			}
			combinationSumHelper(candidates, i+1, target-j*candidates[i], cur, res)
			// 复原
			cur = cur[:len(cur)-j]
		}
	}
}
