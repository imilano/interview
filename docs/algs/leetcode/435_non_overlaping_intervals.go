package leetcode

import "sort"

/*
Given an array of intervals intervals where intervals[i] = [starti, endi],
return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
*/

// 贪心算法。
// 首先需要对区间按照开始时间进行排序，然后怎么贪心呢？其实可以发现，当两个区间重叠时，如果我们想要得到最小的去除数量，那么我们就要尽可能的保留区间，
// 那么，对于两个重叠的区间，我们就需要去除结束时间比较大的那个区间，保留结束时间比较小的那个区间，这样才能够尽可能的留出区间来放置其他区间，从而达到删除
// 最少区间的目的。

// 具体做法是，如果前一个区间的 end 大于后一个区间的 start ，那么就发生了重叠，此时我们就需要删去一个区间， res+1，并且需要删去结束区间比较大的那个区间。此处我们并非进行
// 真实删除，而是用一个变量 last 指向上一个需要比较的区间，我们将 last 指向 end 值比较小的那个区间。如果两个区间没有重叠，那么此时 last 指向当前区间，继续进行下一次遍历。
func eraseOverlapIntervals(intervals [][]int) int {
	size := len(intervals)
	var last, res int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < size; i++ {
		if intervals[i][0] < intervals[last][1] {
			res++
			if intervals[i][1] < intervals[last][1] {
				last = i
			}
		} else {
			last = i
		}
	}

	return res
}

// 同样是贪心，只不过做法不一样
// 可以明确的是，如果连个区间重叠了，那么我们需要保留跨度比较小的那个区间，而删去跨度比较大的那个区间，所以这里根据结束时间对区间进行排序比较好。
// 排好序后，在遍历时，当前区间的开始时间大于等于前一个区间的结束时间，才需要更新前一个区间的结束时间，同时对连续的不重叠区间计数值加 1。
// 遍历结束后，用总长度减去不重叠的区间长度即可。
func eraseOverlapIntervalsSolution2(intervals [][]int) int {
    var res int
    size := len(intervals)
    if size <= 1 {
        return res
    }
    sort.Slice(intervals, func(i,j int) bool {
            return intervals[i][1] < intervals[j][1]
    })
    
    ans, right := 1, intervals[0][1]
    for _, interval := range intervals[1:] {
        if interval[0] >= right {
            ans++
            right = interval[1]
        } 
    }
    

	return size - ans
}

