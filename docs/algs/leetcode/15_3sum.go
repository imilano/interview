package leetcode

import (
	"sort"
	"strconv"
)

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/

/*
 * 注意这里其实并没有说不能更改原数组，那么我们可以对原数组先排个序，然后遍历数组，每次遍历到数 x，从 x 的后面的子数组找出两个和为-x 的数即可。
 * 这里可以注意到右以下几个优化技巧：
 * 	- 遍历的时候只需要遍历到倒数第三个即可。
 *	- 由于限制了不能重复，那么遍历时候对于重复出现的数字需要跳过，策略是，从第二个数字开始，如果这个数字和前一个数字相同，则跳过这个数字，继续遍历下一个。
 *	- 如果当前固定的数 x 是个正数，那么也可以直接跳过这个数，因为既然数组已经排好序了，那么后面的数都只会比 x 大，他们的和也比 x 大
 * 	- 那么如何在子数组中查找-x 呢，可以在子数组中使用双指针，假设一个下标是 i 另一个是 j，那么如果二者之和大于 target，则 j 前移，否则 i 后移。
 *	- i 和 j 移动的过程中需要注意，二者均需要跳过重复数字
 */

func threeSum(nums []int) [][]int {
	return solution(nums)
}

func solution(nums []int) [][]int {
	var res [][]int
	size := len(nums)
	if size < 3 {
		return res
	}
	sort.Ints(nums)

	if nums[0] > 0 || nums[size-1] < 0 {
		return res
	}
	for i := 0; i < size-2; i++ {
		// 正数不需要再遍历
		if nums[i] > 0 {
			break
		}

		// 跳过重复数字
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, size-1
		target := 0 - nums[i]
		for j < k {
			sum := nums[j] + nums[k]
			if sum == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				// 跳过重复
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}

				// 注意跳过重复之后还需要改变下标
				j++
				k--
			} else if sum < target {
				j++
			} else if sum > target {
				k--
			}
		}
	}

	return res
}

func solution2(nums []int) [][]int {
	var res [][]int
	size := len(nums)
	if size < 3 {
		return nil
	}

	for index, value := range nums {
		subAns := subTwoSum(nums[index+1:], -value)
		for _, arr := range subAns {
			arr = append(arr, nums[index])
			res = append(res, arr)
		}
	}

	return res
}

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

var exclude map[string]bool

func subTwoSum(nums []int, target int) [][]int {
	var res [][]int
	m := make(map[int]int)
	for index, value := range nums {
		idx, ok := m[target-value]
		if ok && idx != index {
			var key string
			if nums[idx] < nums[index] {
				key = strconv.Itoa(nums[idx]) + strconv.Itoa(nums[index])
			} else {
				key = strconv.Itoa(nums[index]) + strconv.Itoa(nums[idx])
			}

			if _, ok := exclude[key]; !ok {
				res = append(res, []int{nums[idx], nums[index]})

			}
		}

		m[value] = index
	}

	return res
}
