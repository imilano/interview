package leetcode

/*
Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.
*/

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (1 + n) / 2
	for _, value := range nums {
		sum -= value
	}

	return sum
}
