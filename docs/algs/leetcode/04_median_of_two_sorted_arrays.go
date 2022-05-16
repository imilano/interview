package leetcode

import "math"

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).
*/

// 参见这里 https://youtu.be/LPFhl65R7ww
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 总是保持 nums1 最短
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x,y := len(nums1), len(nums2)
	low, high := 0, x
	for low <= high {
		partitionX := (low+high)/2
		partitionY := (x+y+1)/2 - partitionX

		// 分区，让左右两边的元素个数基本一致（如果总数是偶数，左右相等，如果是奇数，左边比右边多 1）
		// 如果 partitionX 是 0，意味着左边已经没有元素了，那么我们使用 math.MinInt 来代替进行比较
		// 如果 partitionX 是 x，意味着右边已经没有元素了，那么我们使用 math.MaxInt 来代替进行比较
		// 同理对于 partitionY
		maxLeftX, minRightX := getMaxLeft(partitionX, nums1), getMinRight(partitionX, nums1)
		maxLeftY, minRightY := getMaxLeft(partitionY, nums2), getMinRight(partitionY, nums2)

		
		// 如果左边的所有元素都比右边的所有元素要小，说明我们找到了正确的切分点
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			// 如果总数是偶数
			if (x+y) % 2 == 0 {
				return float64(max(maxLeftX, maxLeftY) + min(minRightX, minRightY))/2
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			// nums1 的左半边最大值比 nums2 的右半边最小值还要大，说明nums1 切分不合理，过于偏右，于是需要往左进行偏移
			high = partitionX - 1
		} else {
			// nums2 的左半边最大值比 nums1 的右半边最小值还要大，说明 nums1 切分过于偏左，需要往右边走
			low = partitionX + 1
		}
	}

	return -1
}


func getMaxLeft(partition int, nums []int) int {
	if partition == 0 {
		return math.MinInt
	} 

	return nums[partition - 1]
}


func getMinRight(partition int, nums []int) int {
	if partition == len(nums) {
		return math.MaxInt
	}

	return nums[partition]
}

// 不符合题意的做法，直接再次排序，然后返回中位数
func findMedianSortedArraysSolution1(nums1, nums2 []int) float64 {
	return 0
}
