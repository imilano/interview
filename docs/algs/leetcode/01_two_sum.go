package leetcode

import (
	"fmt"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int {
	var res []int
	if len(nums) <= 1 {
		return res
	}

	dict := make(map[int]int)
	for index, value := range nums {
		dict[value] = index
	}
	for index, num := range nums {
		t := target - num
		if val, ok := dict[t]; ok && val != index {
			return []int{index, val}
		}
	}
	return res
}

func ExplicitTwoSum(nums []int, target int) {
	fmt.Println(twoSum(nums, target))
}

func twoSum2(nums []int, target int) []int {
	if len(nums) <= 1 {
		return nil
	}

	m := make(map[int]int)
	// 两个循环是可以合并的
	// for index, value := range nums {
	// 	m[value] = index
	// }

	// for index, value := range nums {
	// 	idx, ok := m[target-value]
	// 	if ok && idx != index {
	// 		return []int{index, idx}
	// 	}
	// }

	for index, value := range nums {
		t := target - value
		if idx, ok := m[t]; ok {
			return []int{idx, index}
		}
		m[value] = index
	}

	return nil
}
