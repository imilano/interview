package leetcode

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
*/
func containsDuplicate(nums []int) bool {
	size := len(nums)
	if size < 2 {
		return false
	}

	m := make(map[int]int)
	for _, value := range nums {
		if _, ok := m[value]; ok {
			return true
		}

		m[value] = 1
	}

	return false
}
