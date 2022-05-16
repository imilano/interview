package leetcode

import "sort"

/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/

func merge(intervals [][]int) [][]int {
	return mergeSolution2(intervals)
}

// 方法 2，还是先对区间按照开始时间进行排序，然后将第一个区间存入结果中，然后从第二个开始遍历区间集，如果结果中最后一个区间和当前遍历区间无重叠，
// 则直接将当前区间存入结果中，如果有重叠 ，则将结果中最后一个区间的 end 值更新为结果中最后区间的 end 值和当前 end 值之间的较大者，然后继续遍历区间。
func mergeSolution2(intervals [][]int) [][]int {
	size := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	if size > 0 {
		res = append(res, intervals[0])
	}
	rIndex := 0
	for i := 1; i < size; i++ {
		if res[rIndex][1] < intervals[i][0] {
			res = append(res, intervals[i])
			rIndex++
		} else {
			res[rIndex][1] = max(res[rIndex][1], intervals[i][1])
		}
	}

	return res
}

// 方法 1，使用跟 57 题区间插入差不多的方法
// 首先，按照开始时间对区间进行排序，然后进行遍历合并交叉区间。
func mergeSolution1(intervals [][]int) [][]int {
	size := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var cur int
	var res [][]int
	for cur < size {
		newInterval := intervals[cur]
		i := cur
		for i < size && intervals[i][0] <= newInterval[1] {
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
			i++
		}
		res = append(res, newInterval)
		if i > cur {
			cur = i
		} else {
			cur++
		}
	}

	return res
}
