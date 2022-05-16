package leetcode

/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
*/

func permute(nums []int) [][]int {
	var res [][]int
	size := len(nums)
	permuteSolution(0, size, nums, &res)
	return res
}

// TODO fix me
func permuteSolution(first int, size int, nums []int, res *[][]int) {
	if first == size {
		tmp := make([]int, size)
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := first; i < size; i++ {
		// 交换
		nums[i], nums[first] = nums[first], nums[i]
		// 注意这里递归是从哪一个元素开始的，不是从 i+1 开始，而是从 first + 1 开始。
		// permuteSolution(i+1, size, nums, res)
		permuteSolution(first+1, size, nums, res)
		// 撤销
		nums[i], nums[first] = nums[first], nums[i]
	}
}
