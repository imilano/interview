package leetcode

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.
*/

// TODO 这里的二分查找也要好好注意一下
// 因为数组原先是有序的，所以将数组一分为二之后，其中一定有一个是有序的，另一个可能有序，也可能部分有序。此时有序部分用二分查找。无序部分再一分为二，其中一个一定有序，而另一个可能有序，可能无序。循环即可。
func search(nums []int, target int) int {
	size := len(nums)
	left, right := 0, size-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		// 情况 1： 为什么这样可以
		if nums[mid] < nums[right] {
			// 如果右半边有序，则对右半边进行二分
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 如果左半边有序，则对左半边进行二分
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		// 情况 2：而这样不可以
		// 如果左半边有序，则对左半边进行二分
		// if nums[left] < nums[mid] {
		// 	if nums[left] <= target && nums[mid] > target {
		// 		right = mid -1
		// 	} else {
		// 		left = mid + 1
		// 	}
		// } else {
		// 	// 如果右半边有序，则对右半边进行二分
		// 	if nums[mid] < target && nums[right] >= target {
		// 		left = mid + 1
		// 	} else {
		// 		right = mid - 1
		// 	}
		// }

		// 情况 3： 而这样又可以了呢
		// if nums[left] <= nums[mid] {
		// 	if nums[left] <= target && nums[mid] > target {
		// 		right = mid -1
		// 	} else {
		// 		left = mid + 1
		// 	}
		// } else {
		// 	// 如果右半边有序，则对右半边进行二分
		// 	if nums[mid] < target && nums[right] >= target {
		// 		left = mid + 1
		// 	} else {
		// 		right = mid - 1
		// 	}
		// }

		// 原因在于可能会出现 left 和 mid 重叠的情况，比如在[3,1] 中寻找 1，情况 1 中能正确在右半边进行查找；而在情况 2 中，
		// 由于 left 和 mid 重叠了，所以会在右半边中把 right 变成 mid - 1， 从而反而导致将边界缩减到了左半边，这是不对的。
		// 所以我们可加个等号，将这种情况也统归到左半边来。
	}

	return -1
}

func Search(nums []int, target int) int {
	return search(nums, target)
}
