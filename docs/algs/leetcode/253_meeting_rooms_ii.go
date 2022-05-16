package leetcode

import (
	"container/heap"
	"sort"
)

/*
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms required.
*/
func minMeetingRoom(intervals [][]int) int {
	return minMeetingRoomSolution2(intervals)
}

func minMeetingRoomSolution2(intervals [][]int) int {
	size := len(intervals)
	var start, end []int
	for _, interval := range intervals {
		start = append(start, interval[0])
		end = append(end, interval[1])
	}

	sort.Ints(start)
	sort.Ints(end)

	var res, endPos int
	for i := 0; i < size; i++ {
		if start[i] < end[endPos] {
			res++
		} else {
			endPos++
		}
	}

	return res
}

// 这个解法很巧妙，其实可以这么想：横坐标是一条线，那么现在可以把所有的会议标注在二维坐标上，而我么所要求的的就是一条 x = a 的分割线，
// 这条分割线能经过的会议数量最多。
// 一个 interval 可以看成一间会议室的开始时间和关闭时间，而我们所需要找的就是在某一时间段内最多的同时处于开始状态的会议室数量。
func minMeetingRoomSolution1(intervals [][]int) int {
	m := make(map[int]int)
	for _, interval := range intervals {
		m[interval[0]] += 1
		m[interval[1]] -= 1
	}

	var keys []int
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	var rooms, res int
	for _, key := range keys {
		rooms += m[key]
		res = max(res, rooms)
	}
	return res
}


// 使用最小堆来做。这个题目的关键其实还是在于，如何重复利用会议室。之所以需要根据开始时间进行排序，是为了跟已经安排的会议室的结束时间作比较，从而复用这个会议室。
// 那么如何复用呢？由于我们需要的是最早结束的房间，所以我们需要一个数据结构来帮助我们判断哪个会议最早结束，这自然就容易想到最小堆。
// 首先需要将区间按照起始时间从小到大排序。遍历每个区间，每次将该会议的结束时间压入堆中，如果堆非空并且堆顶值小于等于当前遍历到的区间的开始时间，则说明已经有一场会议结束，会议室可用， 
// 无需再安排新房间，可以复用旧房间，则将堆顶元素弹出。最后堆的大小就是我们需要的最多会议室数量。
func minMeetingRoomSolution3(intervals [][]int) int {
	var minHeap MinHeapArr
	heap.Init(&minHeap)
	sort.Slice(intervals, func (i, j int) bool  {
		return intervals[i][0] < intervals[j][0]
	})

	for _, interval := range intervals {
		if minHeap.Len() != 0 && interval[0] >= minHeap.Top().(int) {
			minHeap.Pop()
		}  

		minHeap.Push(interval[1])
	}

	return minHeap.Len()
}
// for test
func MinMeetingRoom(intervals [][]int) int {
	return minMeetingRoom(intervals)
}
