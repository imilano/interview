package leetcode

/*
You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.
*/

// 这个问题需要注意的一点是，原有区间是不重叠的。也就是说，如果要插入新区间，那么插入的新区间与原有区间只可能在末尾时间有重叠，而不会在开始时间有重叠。
// 所以插入的时候，我们只需要将新区间插入到第一个结束时间比新区间开始时间小的区间之后即可。之后再来考虑结束时间重叠的问题。
// 分三个阶段来完成。
// 第一个阶段，在已有区间中找出结束时间比新区间的开始时间小的时区间，然后在这个位置之后将新区间插入进入。
// 第二阶段，此时插入的新区间可能和后面的区间有重叠，也可能没有重叠。
// 	- 没有重叠的话，只需要将剩余区间加到新区间之后即可。
// 	- 有重叠的话，此时就可能会需要合并区间。对于重叠的部分，新区间的开始时间等于重叠部分开始时间的最小值，新区间的结束时间等于重叠部分结束时间的最大值，
// 		扫描直到到达不重叠部分即可。将合并后的区间插入结果。
// 第三阶段，将最后的不重叠区间插入结果中。
func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	var cur int
	size := len(intervals)
	// 阶段 1，添加不重叠区间
	for cur < size && intervals[cur][1] < newInterval[0] {
		res = append(res, intervals[cur])
		cur++
	}
	// 阶段 2，插入新区间，并进行可能的合并
	// for cur < size && newInterval[1] <= intervals[cur][0] {
	for cur < size && newInterval[1] >= intervals[cur][0] {
		newInterval[0] = min(newInterval[0], intervals[cur][0])
		newInterval[1] = max(newInterval[1], intervals[cur][1])
		cur++
	}
	res = append(res, newInterval)

	// 添加剩余不重叠区间
	for cur < size {
		res = append(res, intervals[cur])
		cur++
	}
	return res
}
