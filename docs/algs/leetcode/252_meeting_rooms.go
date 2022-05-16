package leetcode

import "sort"

/*
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), determine if a person could attend all meetings.
*/

// 这也是很简单的一题，其实就是考虑是否会出现区间重叠的情况。
// 解决办法也很简单，我们可以先给区间按照开始时间排个序，然后检查结束时间是否会有重叠即可。
func canAttendMeetings(intervals [][]int) bool {
	size := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < size; i++ {
		if intervals[i-1][1] > intervals[i][0] {
			return false
		}
	}

	return true
}

// for test
func CanAttendMeetings(intervals [][]int) bool {
	return canAttendMeetings(intervals)
}
